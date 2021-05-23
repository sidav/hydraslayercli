package main

import "fmt"

const (
	ITEM_EFFECT_HEALER = iota
	ITEM_EFFECT_REGENERATOR
	ITEM_EFFECT_ACTIVE_ELEMENT_SHIFTING
	ITEM_EFFECT_PASSIVE_ELEMENT_SHIFTING
	TOTAL_PASSIVE_PASSIVE_EFFECTTYPES_NUMBER // for generators
)

type effectData struct {
	canBeOnWeapon        bool
	canBeOnRing          bool
	isActivatable        bool
	name, info           string
	defaultActivatesEach int
}

var effectsStaticData = map[uint8]*effectData{
	ITEM_EFFECT_HEALER: {
		canBeOnWeapon: true,
		canBeOnRing:   true,
		isActivatable: true,
		name:          "healing",
		info:          "Can be activated to heal you.",
	},
	ITEM_EFFECT_REGENERATOR: {
		canBeOnRing:          true,
		canBeOnWeapon:        true,
		defaultActivatesEach: 4,
		info:                 "Heals 1 hp once in a few turns. Also fully heals you between rooms.",
	},
	ITEM_EFFECT_ACTIVE_ELEMENT_SHIFTING: {
		canBeOnWeapon: true,
		isActivatable: true,
		name:          "shifting elements",
		info:          "Can be activated to change element.",
	},
	ITEM_EFFECT_PASSIVE_ELEMENT_SHIFTING: {
		canBeOnWeapon:        true,
		name:                 "instability",
		defaultActivatesEach: 1,
		info:                 "Changes its element each turn randomly.",
	},
}

type effect struct {
	effectCode    uint8
	canBeUsed     bool
	activatesEach int
}

func (e *effect) getInfo() string {
	return effectsStaticData[e.effectCode].info
}

func getRandomEffect(forWeapon, forRing bool) *effect {
	rndEffCode := uint8(rnd.Rand(TOTAL_PASSIVE_PASSIVE_EFFECTTYPES_NUMBER))
	var data *effectData
	for {
		data = effectsStaticData[rndEffCode]
		if (!forWeapon || data.canBeOnWeapon) && (!forRing || data.canBeOnRing) {
			break
		}
		rndEffCode = uint8(rnd.Rand(TOTAL_PASSIVE_PASSIVE_EFFECTTYPES_NUMBER))
	}
	return &effect{
		effectCode:    rndEffCode,
		canBeUsed:     false,
		activatesEach: data.defaultActivatesEach,
	}
}

func (e *effect) isActivatable() bool {
	return effectsStaticData[e.effectCode].isActivatable
}

func (i *item) applyActiveEffect(g *game) {
	if i.effect == nil {
		return
	}
	switch i.effect.effectCode {
	case ITEM_EFFECT_HEALER:
		g.player.hp = g.player.maxhp
	case ITEM_EFFECT_ACTIVE_ELEMENT_SHIFTING:
		i.element = getRandomElement(true, false)
	}
	i.effect.canBeUsed = false
}

func (i *item) applyPassiveEffect(g *game) {
	if i.effect == nil {
		return
	}
	activatable := i.effect.activatesEach == 0 || (g.currentTurn%i.effect.activatesEach == 0)
	if !activatable {
		return
	}
	switch i.effect.effectCode {
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
	case ITEM_EFFECT_PASSIVE_ELEMENT_SHIFTING:
		i.element = getRandomElement(true, false)
	}
}

func (pe *effect) getName() string {
	switch pe.effectCode {
	case ITEM_EFFECT_REGENERATOR:
		return fmt.Sprintf("%d-turn regen", pe.activatesEach)
	}
	data, found := effectsStaticData[pe.effectCode]
	if !found || data.name == "" {
		panic("No effect name...")
	}
	return data.name
}
