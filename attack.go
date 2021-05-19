package main

import (
	"fmt"
	"math"
)

func (g *game) getPossibleAttackStringDescription(w *weapon, e *enemy) string {
	hDmg := g.calculateDamageOnHeads(w, e)
	hRegrw := g.calculateHeadsRegrowAfterHitBy(e, w)
	resHeads := e.heads-hDmg+hRegrw
	as := fmt.Sprintf(
		"If you attack %s with %s, it will lose %d heads and regrow %d, resulting in %d heads. It will bite you for %d damage.",
			e.getName(), w.getName(), hDmg, hRegrw, resHeads, g.calculateDamageByHeads(resHeads),
		)
	return as
}

func (g *game) performPlayerHit(w *weapon, e *enemy) {
	e.heads -= g.calculateDamageOnHeads(w, e)
	e.heads += g.calculateHeadsRegrowAfterHitBy(e, w)
	g.turnMade = true
}

func (g *game) calculateDamageOnHeads(weapon *weapon, enemy *enemy) int {
	// TODO: consider elements
	switch weapon.weaponType {
	case WTYPE_SUBSTRACTOR:
		return weapon.damage
	}
	return 0
}

func (g *game) calculateHeadsRegrowAfterHitBy(enemy *enemy, weapon *weapon) int {
	return 1
}

func (g *game) calculateDamageByHeads(headsNum int) int {
	// TODO: consider elements
	return int(math.Log2(float64(headsNum)))
}
