package main

import "fmt"

const (
	ITEM_EFFECT_HEALER = iota
	ITEM_EFFECT_REGENERATOR
	ITEM_EFFECT_ACTIVE_ELEMENT_SHIFTING
	ITEM_EFFECT_PASSIVE_ELEMENT_SHIFTING
	ITEM_EFFECT_PASSIVE_DISTORTION
	ITEM_EFFECT_ONHIT_ADDDAMAGE
	ITEM_EFFECT_SHOOT
	TOTAL_ITEM_EFFECTTYPES_NUMBER // for generators
)

type brandData struct {
	canBeOnWeapon           bool
	canBeOnRing             bool
	isActivatable           bool
	name, info              string
	defaultActivatesEach    int
	isChargeable            bool
	defaultCharges          int
	defaultAdditionalDamage int
}

var brandsStaticData = map[uint8]*brandData{
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
	ITEM_EFFECT_PASSIVE_DISTORTION: {
		canBeOnWeapon:        true,
		name:                 "distortion",
		defaultActivatesEach: 1,
		info:                 "Changes its damage each turn randomly.",
	},
	ITEM_EFFECT_ONHIT_ADDDAMAGE: {
		canBeOnWeapon:           true,
		defaultAdditionalDamage: 1,
		name:                    "",
		info:                    "Deals additional damage after initial hit with no regrow.",
	},
	ITEM_EFFECT_SHOOT: {
		canBeOnWeapon:           true,
		defaultAdditionalDamage: 1,
		defaultCharges:          1,
		isChargeable:            true,
		name:                    "",
		info:                    "Can shoot (use fire or shoot command).",
	},
}

type effect struct {
	effectCode       uint8
	canBeUsed        bool
	activatesEach    int
	additionalDamage int
	charges          int
}

func (e *effect) getHelpText() string {
	return brandsStaticData[e.effectCode].info
}

func (e *effect) getStaticData() *brandData {
	return brandsStaticData[e.effectCode]
}

func (e *effect) isChargeable() bool {
	return e.getStaticData().isChargeable
}

func getRandomEffect(forWeapon, forRing bool) *effect {
	rndEffCode := uint8(rnd.Rand(TOTAL_ITEM_EFFECTTYPES_NUMBER))
	var data *brandData
	for {
		data = brandsStaticData[rndEffCode]
		if (!forWeapon || data.canBeOnWeapon) && (!forRing || data.canBeOnRing) {
			break
		}
		rndEffCode = uint8(rnd.Rand(TOTAL_ITEM_EFFECTTYPES_NUMBER))
	}
	return &effect{
		effectCode:       rndEffCode,
		canBeUsed:        false,
		activatesEach:    data.defaultActivatesEach,
		charges:          data.defaultCharges,
		additionalDamage: data.defaultAdditionalDamage,
	}
}

func (e *effect) isActivatable() bool {
	return brandsStaticData[e.effectCode].isActivatable
}

func (i *item) applyActiveEffect(g *game) {
	if i.effect == nil {
		return
	}
	switch i.effect.effectCode {
	case ITEM_EFFECT_HEALER:
		g.player.hp += g.player.maxhp / 2
		if g.player.hp > g.player.maxhp {
			g.player.hp = g.player.maxhp
		}
	case ITEM_EFFECT_ACTIVE_ELEMENT_SHIFTING:
		i.element = getRandomElement(true, false, true)
	}
	i.effect.canBeUsed = false
}

func (i *item) applyPassiveEffect(g *game) {
	if i.effect == nil {
		return
	}
	if g.currentTurn == 0 {
		i.effect.canBeUsed = true
	}
	isActiveCurrentTurn := i.effect.activatesEach == 0 || (g.currentTurn%i.effect.activatesEach == 0)
	if !isActiveCurrentTurn {
		return
	}
	switch i.effect.effectCode {
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
		i.element = getRandomElement(true, false, true)
	case ITEM_EFFECT_PASSIVE_DISTORTION:
		i.weaponInfo.damage = auxrnd.RandInRange(1, g.currentStageNumber+2)
	}
}

func (i *item) applyOnHitEffect(g *game, target *enemy) {
	if i.effect == nil {
		return
	}
	switch i.effect.effectCode {
	case ITEM_EFFECT_ONHIT_ADDDAMAGE:
		if target.heads > i.effect.additionalDamage {
			target.heads -= i.effect.additionalDamage
			g.appendToLogMessage("%s discharges its magic on %s for %d damage!", i.getName(), target.getName(), i.effect.additionalDamage)
		}
	}
}

func (pe *effect) getName() string {
	switch pe.effectCode {
	case ITEM_EFFECT_REGENERATOR:
		return fmt.Sprintf("%d-turn regen", pe.activatesEach)
	case ITEM_EFFECT_ONHIT_ADDDAMAGE:
		return fmt.Sprintf("+%d damage", pe.additionalDamage)
	case ITEM_EFFECT_SHOOT:
		return fmt.Sprintf("discharge(%d)", pe.charges)
	}
	data, found := brandsStaticData[pe.effectCode]
	if !found || data.name == "" {
		panic("No effect name...")
	}
	return data.name
}
