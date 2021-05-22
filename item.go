package main

import (
	"fmt"
	"strings"
)

const (
	ITEM_HEAL = iota
	ITEM_ENCHANTER
	ITEM_DESTROY_HYDRA
	ITEM_CONFUSE_HYDRA
	ITEM_MASS_CONFUSION
	ITEM_INCREASE_HP
	ITEM_STRENGTH
	ITEM_CHANGE_ELEMENT
	ITEM_UNELEMENT_ENEMIES
	ITEM_DECAPITATION
	ITEM_IMPROVE_COOLDOWN
	ITEM_AMMO
	TOTAL_ITEM_TYPES_NUMBER // for generators
)

type consumableItemInfo struct {
	consumableType uint8
	name           string
	frequency      int
}

type item struct {
	element       uint8
	asConsumable  *consumableItemInfo
	specialName   string // for randarts and non-consumables
	passiveEffect *passiveEffect
	weaponInfo    *weapon
	count         int
}

func (i *item) isWeapon() bool {
	return i.weaponInfo != nil
}

func (i *item) isAmmo() bool {
	return i.asConsumable != nil && i.asConsumable.consumableType == ITEM_AMMO
}

func (i *item) isConsumable() bool {
	return i.asConsumable != nil
}

func (i *item) hasPassiveEffect() bool {
	return i.passiveEffect != nil
}

func (i *item) getName() string {
	name := getElementName(i.element)
	if len(name) > 0 {
		name += " "
	}
	if i.specialName != "" {
		name += i.specialName
	}
	if i.isWeapon() {
		name += i.weaponInfo.getName()
	}
	if i.asConsumable != nil {
		switch i.asConsumable.consumableType {
		case ITEM_HEAL:
			name += "Healing powder"
		case ITEM_ENCHANTER:
			name += "Scroll of enchant weapon +1"
		case ITEM_DESTROY_HYDRA:
			name += "Scroll of destroy hydra"
		case ITEM_CONFUSE_HYDRA:
			name += "Scroll of confuse hydra"
		case ITEM_INCREASE_HP:
			name += "Potion of vitality"
		case ITEM_STRENGTH:
			name += "Potion of strength"
		case ITEM_CHANGE_ELEMENT:
			name += "Scroll of change element"
		case ITEM_UNELEMENT_ENEMIES:
			name += "Scroll of nullify"
		case ITEM_DECAPITATION:
			name += "Scroll of mass decapitation"
		case ITEM_IMPROVE_COOLDOWN:
			name += "Scroll of faster cooldown"
		case ITEM_MASS_CONFUSION:
			name += "Scroll of mass confusion"
		case ITEM_AMMO:
			name += fmt.Sprintf("%d arbalest bolts", i.count)
		}
	}
	if i.hasPassiveEffect() {
		name += " of " + i.passiveEffect.getName()
	}
	if name == "" {
		panic("No item name!")
	}
	return colorizeString(getElementColorStr(i.element), strings.Title(name))
}
