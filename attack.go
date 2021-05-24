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
	hRegrw := g.calculateHeadsRegrowAfterHitBy(e, w)
	resHeads := e.heads-hDmg+hRegrw
	as := fmt.Sprintf("If you attack %s with %s", e.getName(), w.getName())
	if hDmg == 0 {
		as += ", it will lose no heads"
	} else {
		as += fmt.Sprintf(", it will lose %d heads", hDmg)
	}
	if e.heads - hDmg == 0 {
		as += " and die."
	} else {
		as += fmt.Sprintf(" and regrow %d, resulting in %d heads. It will bite you for %d damage.", hRegrw, resHeads, g.calculateDamageByHeads(resHeads))
	}
	return as
}

func (g *game) getShortPossibleAttackStringDescription(w *item, e *enemy) string {
	if !g.showShortCombatDescription {
		return ""
	}
	hDmg := g.calculateDamageOnHeads(w.weaponInfo, e)
	hRegrw := g.calculateHeadsRegrowAfterHitBy(e, w)
	resHeads := e.heads-hDmg+hRegrw
	enemyDmgStr := fmt.Sprintf("; bite %d", g.calculateDamageByHeads(resHeads))
	if hDmg < e.heads {
		return fmt.Sprintf(" (%d-%d+%d=%d%s)", e.heads, hDmg, hRegrw, resHeads, enemyDmgStr)
	} else {
		return fmt.Sprintf(" (-> kill)")
	}
}


func (g *game) performPlayerHit(w *item, e *enemy) {
	damage := g.calculateDamageOnHeads(w.weaponInfo, e)
	g.currLog = fmt.Sprintf("You hit %s with %s, cutting %d heads. ",
		e.getName(),
		w.getName(), damage)
	regrow := g.calculateHeadsRegrowAfterHitBy(e, w)
	e.heads -= damage
	if e.heads > 0 {
		g.currLog += fmt.Sprintf("It grows %d heads!", regrow)
		e.heads += regrow
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
	regrow := g.calculateHeadsRegrowAfterHitBy(e, w)
	e.heads -= damage
	if e.heads > 0 {
		g.currLog += fmt.Sprintf("It grows %d heads!", regrow)
		e.heads += regrow
	} else {
		g.currLog += fmt.Sprintf("It drops dead!")
	}
	g.player.spendAmmo()
	g.turnMade = true
	g.enemiesSkipTurn = true
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
		if enemy.heads % weapon.damage != 0 {
			return 0
		}
		return enemy.heads - enemy.heads / weapon.damage
	}
	return 0
}

func (g *game) calculateHeadsRegrowAfterHitBy(enemy *enemy, weapon *item) int {
	regrow := getHeadRegrowForElement(enemy.element, weapon.element)
	if regrow == -2 {
		return enemy.heads - g.calculateDamageOnHeads(weapon.weaponInfo, enemy)
	}
	return regrow
}

func (g *game) calculateDamageByHeads(headsNum int) int {
	// TODO: consider elements
	damage := int(math.Log2(float64(headsNum)))
	if damage == 0 {
		damage = 1
	}
	return damage
}
