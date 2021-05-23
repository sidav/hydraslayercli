package main

import (
	"fmt"
	"strings"
)

type item struct {
	element      *element
	asConsumable *consumableItemInfo
	specialName  string // for randarts and non-consumables
	effect       *effect
	weaponInfo   *weapon
	count        int
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

func (i *item) hasEffect() bool {
	return i.effect != nil
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
	if i.hasEffect() {
		name += " of " + i.effect.getName()
		if !i.effect.canBeUsed && i.effect.isActivatable() {
			name += colorizeString(Gray, " (inactive)")
		}
	}
	if name == "" {
		panic("No item name!")
	}
	if i.element != nil {
		return colorizeString(i.element.colorString, strings.Title(name))
	}
	return strings.Title(name)
}
