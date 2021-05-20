package main

import (
	"fmt"
	"math"
)

func (g *game) getPossibleAttackStringDescription(w *weapon, e *enemy) string {
	hDmg := g.calculateDamageOnHeads(w, e)
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

func (g *game) performPlayerHit(w *weapon, e *enemy) {
	e.heads -= g.calculateDamageOnHeads(w, e)
	if e.heads > 0 {
		e.heads += g.calculateHeadsRegrowAfterHitBy(e, w)
	}
	g.turnMade = true
}

func (g *game) calculateDamageOnHeads(weapon *weapon, enemy *enemy) int {
	// TODO: consider elements
	switch weapon.weaponType {
	case WTYPE_SUBSTRACTOR:
		return weapon.damage
	case WTYPE_DIVISOR:
		if enemy.heads % weapon.damage != 0 {
			return 0
		}
		return enemy.heads - enemy.heads / weapon.damage
	}
	return 0
}

func (g *game) calculateHeadsRegrowAfterHitBy(enemy *enemy, weapon *weapon) int {
	return 1
}

func (g *game) calculateDamageByHeads(headsNum int) int {
	// TODO: consider elements
	damage := int(math.Log2(float64(headsNum)))
	if damage == 0 {
		damage = 1
	}
	return damage
}
