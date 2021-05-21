package main

import "fmt"

const (
	PASSIVE_EFFECT_HEALER = iota
	PASSIVE_EFFECT_REGENERATOR
	TOTAL_PASSIVE_PASSIVE_EFFECTTYPES_NUMBER // for generators
)

type passiveEffect struct {
	effectType    uint8
	activatesEach int
}

func getRandomPassiveEffect() uint8 {
	return uint8(rnd.Rand(TOTAL_PASSIVE_PASSIVE_EFFECTTYPES_NUMBER))
}

func (i *item) applyPassiveEffect(g *game) {
	if i.passiveEffect == nil {
		return
	}
	activatable := i.passiveEffect.activatesEach == 0 || g.currentTurn % i.passiveEffect.activatesEach == 0
	if !activatable {
		return
	}
	switch i.passiveEffect.effectType {
	case PASSIVE_EFFECT_HEALER:
		if g.currentTurn == 1 {
			g.player.hp = g.player.maxhp
		}
	case PASSIVE_EFFECT_REGENERATOR:
		if g.player.hp < g.player.maxhp {
			g.player.hp++
		}
	}
}

func (pe *passiveEffect) getName() string {
	switch pe.effectType {
	case PASSIVE_EFFECT_HEALER:
		return "healing"
	case PASSIVE_EFFECT_REGENERATOR:
		return fmt.Sprintf("%d-turn regen", pe.activatesEach)
	}
	panic("No passive effect name...")
}
