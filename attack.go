package main

import (
	"fmt"
	"math"
)

func (g *game) getPossibleAttackStringDescription(w *item, e *enemy) string {
	if w == nil {
		return "ERROR: nil item in attack.go. Report this please. "
	}
	if w.weaponInfo == nil {
		return "ERROR: item is not a weapon in attack.go. Report this please. ."
	}
	if e == nil {
		return "ERROR: nil enemy in attack.go. Report this please. "
	}
	hDmg := g.calculateDamageOnHeads(w.weaponInfo, e)
	hRegrw, rType := g.calculateHeadsRegrowAfterHitBy(e, w)
	resHeads := e.heads - hDmg + hRegrw
	as := fmt.Sprintf("If you attack %s with %s", e.getName(), w.getName())
	if hDmg == 0 {
		as += ", it will lose no heads"
	} else {
		as += fmt.Sprintf(", it will lose %d heads", hDmg)
	}
	if e.heads-hDmg == 0 {
		as += " and die."
	} else {
		if hRegrw != 0 {
			switch rType {
			case REGROW_SIMPLE:	as += fmt.Sprintf(" and regrow %d, resulting in %d heads. ", hRegrw, resHeads)
			case REGROW_DUPLICATE: as += fmt.Sprintf(" and duplicate remaining,resulting in %d heads. ", resHeads)
			}
		} else {
			as += fmt.Sprintf(" with no regrow, resulting in %d heads. ", resHeads)
		}
		afterDmg := 0
		if w.hasEffect() && w.effect.effectCode == ITEM_EFFECT_ONHIT_ADDDAMAGE && w.effect.additionalDamage <= resHeads {
			afterDmg = w.effect.additionalDamage
		}
		if afterDmg > 0 {
			as += fmt.Sprintf("It will then suffer additional %d damage", afterDmg)
		}
		if resHeads - afterDmg <= 0 {
			as += " and die."
			return as
		} else if afterDmg > 0 {
			as += ". "
		}
		as += fmt.Sprintf("It will bite you for %d damage.", g.calculateDamageByHeads(resHeads-afterDmg))
	}
	return as
}

func (g *game) getShortPossibleAttackStringDescription(w *item, e *enemy) string {
	if !g.showShortCombatDescription {
		return ""
	}
	hDmg := g.calculateDamageOnHeads(w.weaponInfo, e)
	hRegrw, rType := g.calculateHeadsRegrowAfterHitBy(e, w)
	descriptionString := " "
	resHeads := e.heads - hDmg + hRegrw
	afterDmg := 0
	if w.hasEffect() && w.effect.effectCode == ITEM_EFFECT_ONHIT_ADDDAMAGE && w.effect.additionalDamage <= resHeads {
		afterDmg = w.effect.additionalDamage
	}
	if hDmg+afterDmg-hRegrw >= e.heads {
		return fmt.Sprintf(" (-> kill)")
	}
	enemyDmgStr := fmt.Sprintf("; bite %d", g.calculateDamageByHeads(resHeads))
	damStr := ""
	switch w.weaponInfo.weaponType {
	case WTYPE_DIVISOR:
		if hDmg > 0 {
			damStr = fmt.Sprintf("/%d", w.weaponInfo.damage)
		}
	default:
		damStr = fmt.Sprintf("-%d", hDmg)
	}
	regrowStr := ""
	switch rType {
	case REGROW_SIMPLE:
		if hRegrw > 0 {
			regrowStr = fmt.Sprintf("+%d", hRegrw)
		}
	case REGROW_DUPLICATE:
		descriptionString += "("
		regrowStr = ")*2"
	}
	additionalDamStr := ""
	if afterDmg > 0 {
		additionalDamStr = fmt.Sprintf("-%d", afterDmg)
		resHeads -= afterDmg
	}

	descriptionString += fmt.Sprintf("(%d%s%s%s=%d%s)",
		e.heads, damStr, regrowStr, additionalDamStr, resHeads, enemyDmgStr)
	return descriptionString
}

func (g *game) performPlayerHit(w *item, e *enemy) {
	damage := g.calculateDamageOnHeads(w.weaponInfo, e)
	g.currLog = fmt.Sprintf("You hit %s with %s, cutting %d heads. ",
		e.getName(),
		w.getName(), damage)
	regrow, regrowType := g.calculateHeadsRegrowAfterHitBy(e, w)
	e.heads -= damage
	if e.heads > 0 {
		if regrow > 0 {
			switch regrowType {
			case REGROW_SIMPLE:
				g.currLog += fmt.Sprintf("It grows %d heads!", regrow)
			case REGROW_DUPLICATE:
				g.currLog += fmt.Sprintf("It duplicates its %d heads!", e.heads)
			default:
				panic("No text for " + regrowType)
			}
			e.heads += regrow
		}
		w.applyOnHitEffect(g, e)
	} else {
		g.currLog += fmt.Sprintf("It drops dead!")
	}
	g.turnMade = true
}

func (g *game) performPlayerShoot(w *item, e *enemy) {
	if !w.weaponInfo.canShoot {
		g.setLogMessage("%s can't shoot!", w.getName())
		return
	}
	if !g.player.hasAmmo() {
		g.setLogMessage("You are out of ammunition!")
		return
	}
	damage := g.calculateDamageOnHeads(w.weaponInfo, e)
	g.setLogMessage("You shoot %s with a %s, destroying %d heads. ",
		e.getName(),
		w.getName(), damage)
	regrow, regrowType := g.calculateHeadsRegrowAfterHitBy(e, w)
	e.heads -= damage
	if e.heads > 0 {
		switch regrowType {
		case REGROW_SIMPLE:
			g.currLog += fmt.Sprintf("It grows %d heads!", regrow)
		case REGROW_DUPLICATE:
			g.currLog += fmt.Sprintf("It duplicates its %d heads!", e.heads)
		default:
			panic("No text for " + regrowType)
		}
		e.heads += regrow
	} else {
		g.currLog += fmt.Sprintf("It drops dead!")
	}
	g.player.spendAmmo()
	g.turnMade = true
	g.allEnemiesSkipTurn = true
}

func (g *game) calculateDamageOnHeads(weapon *weapon, enemy *enemy) int {
	if weapon == nil {
		return 0
	}
	// TODO: consider elements
	switch weapon.weaponType {
	case WTYPE_SUBSTRACTOR:
		if weapon.damage > enemy.heads {
			return 0
		}
		return weapon.damage
	case WTYPE_DIVISOR:
		if enemy.heads%weapon.damage != 0 {
			return 0
		}
		return enemy.heads - enemy.heads/weapon.damage
	case WTYPE_LOGARITHMER:
		base := weapon.damage
		if base < 2 {
			return 0
		}
		heads := enemy.heads
		power := 0
		for heads != 1 {
			if heads % base != 0 {
				return 0
			}
			heads = heads / base
			power += 1
		}
		return enemy.heads - power
	}
	return 0
}

func (g *game) calculateHeadsRegrowAfterHitBy(enemy *enemy, weapon *item) (int, string) {
	regrow := getHeadRegrowForElement(enemy.element, weapon.element)
	if regrow == -2 {
		return enemy.heads - g.calculateDamageOnHeads(weapon.weaponInfo, enemy), REGROW_DUPLICATE
	}
	return regrow, REGROW_SIMPLE
}

func (g *game) calculateDamageByHeads(headsNum int) int {
	// TODO: consider elements
	damage := int(math.Log2(float64(headsNum)))
	if damage == 0 {
		damage = 1
	}
	return damage
}
