package main

import (
	"fmt"
	"strings"
)

type item struct {
	element          *element
	asConsumable     *consumableItemInfo
	specialName      string // for randarts and non-consumables
	brand            *brand
	auxiliaryBrand   *brand // needed e.g. for scrolls of specific branding.
	auxiliaryElement *element
	weaponInfo       *weapon
	count            int
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
	return i.brand != nil
}

func (i *item) getInfo() string {
	info := i.getName()
	if i.hasEffect() {
		info += ": "
		info += i.brand.getHelpText()
	}
	if i.isConsumable() {
		info += ": "
		info += i.asConsumable.info
	}
	return info
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
			name += fmt.Sprintf("x%d ", i.count)
		}
		name += i.asConsumable.name
		switch i.asConsumable.consumableType {
		case ITEM_CHANGE_ELEMENT_SPECIFIC:
			name += fmt.Sprintf("%s element",i.auxiliaryElement.name)
			name = colorizeString(i.auxiliaryElement.getElementColorStr(), name)
		case ITEM_BRANDING_SPECIFIC:
			name += fmt.Sprintf("\"%s\" brand",i.auxiliaryBrand.getName())
		}
	}
	if i.hasEffect() {
		name += " of " + i.brand.getName()
		if !i.brand.canBeUsed && i.brand.isActivatable() {
			name += colorizeString(Gray, " (inactive)")
		}
	}
	//if i.auxiliaryElement != nil {
	//	name += " " + i.auxiliaryElement.name
	//}
	//if i.auxiliaryEffect != nil {
	//	name += " " + i.auxiliaryEffect.getName()
	//}
	if name == "" {
		panic("No item name!")
	}
	if i.element != nil {
		return colorizeString(i.element.getElementColorStr(), strings.Title(name))
	}
	return strings.Title(name)
}
