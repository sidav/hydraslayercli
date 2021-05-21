package main

const (
	ITEM_NOTYPE = iota
	ITEM_HEAL
	ITEM_ENCHANTER
	ITEM_DESTROY_HYDRA
	ITEM_CONFUSE_HYDRA
	ITEM_INCREASE_HP
	ITEM_STRENGTH
	ITEM_CHANGE_ELEMENT
	TOTAL_ITEM_TYPES_NUMBER // for generators
)

type item struct {
	itemConsumableType uint8
	specialName        string // for randarts and non-consumables 
	passiveEffect      *passiveEffect
	weaponInfo         *weapon
}

func (i *item) isWeapon() bool {
	return i.weaponInfo != nil
}

func (i *item) hasPassiveEffect() bool {
	return i.passiveEffect != nil
}

func (i *item) getName() string {
	name := ""
	if i.specialName != "" {
		name = i.specialName
	}
	if i.isWeapon() {
		name = i.weaponInfo.getName()
	}
	switch i.itemConsumableType {
	case ITEM_HEAL:
		name = "Healing powder"
	case ITEM_ENCHANTER:
		name = "Scroll of enchant weapon +1"
	case ITEM_DESTROY_HYDRA:
		name = "Scroll of destroy hydra"
	case ITEM_CONFUSE_HYDRA:
		name = "Scroll of confuse hydra"
	case ITEM_INCREASE_HP:
		name = "Potion of vitality"
	case ITEM_STRENGTH:
		name = "Potion of strength"
	case ITEM_CHANGE_ELEMENT:
		name = "Scroll of change element"
	}
	if i.hasPassiveEffect() {
		name += " of " + i.passiveEffect.getName()
	}
	if name == "" {
		panic("No item name!")
	}
	return name
}
