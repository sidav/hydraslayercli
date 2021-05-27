package main

import "fmt"

const (
	BRAND_HEALER = iota
	BRAND_REGENERATOR
	BRAND_ACTIVE_ELEMENT_SHIFTING
	BRAND_PASSIVE_ELEMENT_SHIFTING
	BRAND_PASSIVE_DISTORTION
	BRAND_ONHIT_ADDDAMAGE
	BRAND_DISCHARGE_SHOT
	TOTAL_BRANDTYPES_NUMBER // for generators
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
	BRAND_HEALER: {
		canBeOnWeapon: true,
		canBeOnRing:   true,
		isActivatable: true,
		name:          "healing",
		info:          "Can be activated to heal you.",
	},
	BRAND_REGENERATOR: {
		canBeOnRing:          true,
		canBeOnWeapon:        true,
		defaultActivatesEach: 4,
		info:                 "Heals 1 hp once in a few turns. Also fully heals you between rooms.",
	},
	BRAND_ACTIVE_ELEMENT_SHIFTING: {
		canBeOnWeapon: true,
		isActivatable: true,
		name:          "shifting elements",
		info:          "Can be activated to change element.",
	},
	BRAND_PASSIVE_ELEMENT_SHIFTING: {
		canBeOnWeapon:        true,
		name:                 "instability",
		defaultActivatesEach: 1,
		info:                 "Changes its element each turn randomly.",
	},
	BRAND_PASSIVE_DISTORTION: {
		canBeOnWeapon:        true,
		name:                 "distortion",
		defaultActivatesEach: 1,
		info:                 "Changes its damage each turn randomly.",
	},
	BRAND_ONHIT_ADDDAMAGE: {
		canBeOnWeapon:           true,
		defaultAdditionalDamage: 1,
		name:                    "",
		info:                    "Deals additional damage after initial hit with no regrow.",
	},
	BRAND_DISCHARGE_SHOT: {
		canBeOnWeapon:           true,
		defaultAdditionalDamage: 1,
		defaultCharges:          1,
		isChargeable:            true,
		name:                    "",
		info:                    "Can shoot (use fire or shoot command).",
	},
}

type brand struct {
	effectCode       uint8
	canBeUsed        bool
	activatesEach    int
	additionalDamage int
	charges          int
}

func (e *brand) getHelpText() string {
	return brandsStaticData[e.effectCode].info
}

func (e *brand) getStaticData() *brandData {
	return brandsStaticData[e.effectCode]
}

func (e *brand) isChargeable() bool {
	return e.getStaticData().isChargeable
}

func getRandomBrand(forWeapon, forRing bool) *brand {
	rndEffCode := uint8(rnd.Rand(TOTAL_BRANDTYPES_NUMBER))
	var data *brandData
	for {
		data = brandsStaticData[rndEffCode]
		if (!forWeapon || data.canBeOnWeapon) && (!forRing || data.canBeOnRing) {
			break
		}
		rndEffCode = uint8(rnd.Rand(TOTAL_BRANDTYPES_NUMBER))
	}
	return &brand{
		effectCode:       rndEffCode,
		canBeUsed:        false,
		activatesEach:    data.defaultActivatesEach,
		charges:          data.defaultCharges,
		additionalDamage: data.defaultAdditionalDamage,
	}
}

func (e *brand) isActivatable() bool {
	return brandsStaticData[e.effectCode].isActivatable
}

func (i *item) applyActiveEffect(g *game) {
	if i.brand == nil {
		return
	}
	switch i.brand.effectCode {
	case BRAND_HEALER:
		g.player.hp += g.player.maxhp / 2
		if g.player.hp > g.player.maxhp {
			g.player.hp = g.player.maxhp
		}
	case BRAND_ACTIVE_ELEMENT_SHIFTING:
		i.element = getRandomElement(true, true, true)
	}
	i.brand.canBeUsed = false
}

func (i *item) applyPassiveEffect(g *game) {
	if i.brand == nil {
		return
	}
	if g.currentTurn == 0 {
		i.brand.canBeUsed = true
	}
	isActiveCurrentTurn := i.brand.activatesEach == 0 || (g.currentTurn%i.brand.activatesEach == 0)
	if !isActiveCurrentTurn {
		return
	}
	switch i.brand.effectCode {
	case BRAND_REGENERATOR:
		if g.currentTurn == 0 {
			g.player.hp = g.player.maxhp
			g.appendToLogMessage("You are healed.")
		}
		if g.player.hp < g.player.maxhp {
			g.player.hp++
			g.appendToLogMessage("You regenerate.")
		}
	case BRAND_PASSIVE_ELEMENT_SHIFTING:
		i.element = getRandomElement(true, true, true)
	case BRAND_PASSIVE_DISTORTION:
		i.weaponInfo.damage = auxrnd.RandInRange(1, g.currentStageNumber+2)
	}
}

func (i *item) applyOnHitEffect(g *game, target *enemy) {
	if i.brand == nil {
		return
	}
	switch i.brand.effectCode {
	case BRAND_ONHIT_ADDDAMAGE:
		if target.heads > i.brand.additionalDamage {
			target.heads -= i.brand.additionalDamage
			g.appendToLogMessage("%s discharges its magic on %s for %d damage!", i.getName(), target.getName(), i.brand.additionalDamage)
		}
	}
}

func (pe *brand) getName() string {
	switch pe.effectCode {
	case BRAND_REGENERATOR:
		return fmt.Sprintf("%d-turn regen", pe.activatesEach)
	case BRAND_ONHIT_ADDDAMAGE:
		return fmt.Sprintf("+%d damage", pe.additionalDamage)
	case BRAND_DISCHARGE_SHOT:
		return fmt.Sprintf("discharge(%d)", pe.charges)
	}
	data, found := brandsStaticData[pe.effectCode]
	if !found || data.name == "" {
		panic("No brand name...")
	}
	return data.name
}
