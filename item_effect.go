package main

import "fmt"

const (
	ITEM_EFFECT_HEALER = iota
	ITEM_EFFECT_REGENERATOR
	TOTAL_PASSIVE_PASSIVE_EFFECTTYPES_NUMBER // for generators
)

type effect struct {
	effectType    uint8
	canBeUsed     bool
	activatesEach int
}

func getRandomPassiveEffect() uint8 {
	return uint8(rnd.Rand(TOTAL_PASSIVE_PASSIVE_EFFECTTYPES_NUMBER))
}

func (e *effect) isActivatable() bool {
	return e.effectType == ITEM_EFFECT_HEALER
}

func (i *item) applyActiveEffect(g *game) {
	if i.effect == nil {
		return
	}
	switch i.effect.effectType {
	case ITEM_EFFECT_HEALER:
		g.player.hp = g.player.maxhp
		i.effect.canBeUsed = false
	}
}

func (i *item) applyPassiveEffect(g *game) {
	if i.effect == nil {
		return
	}
	activatable := i.effect.activatesEach == 0 || (g.currentTurn%i.effect.activatesEach == 0)
	if !activatable {
		return
	}
	switch i.effect.effectType {
	case ITEM_EFFECT_HEALER:
		if g.currentTurn == 0 {
			i.effect.canBeUsed = true
		}
	case ITEM_EFFECT_REGENERATOR:
		if g.currentTurn == 0 {
			g.player.hp = g.player.maxhp
			g.appendToLogMessage("You are healed.")
		}
		if g.player.hp < g.player.maxhp {
			g.player.hp++
			g.appendToLogMessage("You regenerate.")
		}
	}
}

func (pe *effect) getName() string {
	switch pe.effectType {
	case ITEM_EFFECT_HEALER:
		return "healing"
	case ITEM_EFFECT_REGENERATOR:
		return fmt.Sprintf("%d-turn regen", pe.activatesEach)
	}
	panic("No passive effect name...")
}
