package main

const (
	ITEM_NOTYPE = iota
	ITEM_ENCHANTER
	ITEM_HEAL
	ITEM_DESTROY_HYDRA
	ITEM_CONFUSE_HYDRA
	TOTAL_ITEM_TYPES_NUMBER // for generators
)

type item struct {
	itemType   uint8
	weaponInfo *weapon
}

func (i *item) isWeapon() bool {
	return i.weaponInfo != nil
}

func (i *item) getName() string {
	if i.isWeapon() {
		return i.weaponInfo.getName()
	}
	switch i.itemType {
	case ITEM_HEAL: return "Healing powder"
	case ITEM_ENCHANTER: return "Scroll of enchant weapon +1"
	case ITEM_DESTROY_HYDRA: return "Scroll of destroy hydra"
	}
	panic("No item name!")
}
