package main

import (
	"fmt"
	"strings"
)

type item struct {
	element       *element
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
	name := ""
	if i.element != nil {
		name += i.element.name
	}
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
		if i.isAmmo() {
			name += fmt.Sprintf("%d ", i.count)
		}
		name += i.asConsumable.name
	}
	if i.hasPassiveEffect() {
		name += " of " + i.passiveEffect.getName()
	}
	if name == "" {
		panic("No item name!")
	}
	if i.element != nil {
		return colorizeString(i.element.colorString, strings.Title(name))
	}
	return strings.Title(name)
}
