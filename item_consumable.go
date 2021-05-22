package main

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

var consumablesData = []*consumableItemInfo {
	{
		consumableType: ITEM_HEAL,
		name:           "Healing powder",
		frequency:      2,
	},
	{
		consumableType: ITEM_ENCHANTER,
		name:           "Scroll of enchant weapon",
		frequency:      2,
	},
	{
		consumableType: ITEM_DESTROY_HYDRA,
		name:           "Scroll of destroy hydra",
		frequency:      1,
	},
	{
		consumableType: ITEM_CONFUSE_HYDRA,
		name:           "Scroll of confuse hydra",
		frequency:      1,
	},
	{
		consumableType: ITEM_MASS_CONFUSION,
		name:           "Scroll of confused party",
		frequency:      1,
	},
	{
		consumableType: ITEM_INCREASE_HP,
		name:           "Potion of vitality",
		frequency:      2,
	},
	{
		consumableType: ITEM_STRENGTH,
		name:           "Potion of strength",
		frequency:      2,
	},
	{
		consumableType: ITEM_CHANGE_ELEMENT,
		name:           "Scroll of change element",
		frequency:      1,
	},
	{
		consumableType: ITEM_UNELEMENT_ENEMIES,
		name:           "Scroll of nullification",
		frequency:      1,
	},
	{
		consumableType: ITEM_DECAPITATION,
		name:           "Scroll of bisection",
		frequency:      1,
	},
	{
		consumableType: ITEM_IMPROVE_COOLDOWN,
		name:           "Scroll of improve cooldown",
		frequency:      1,
	},
	{
		consumableType: ITEM_AMMO,
		name:           "Arbalest bolts",
		frequency:      2,
	},
}

type consumableItemInfo struct {
	consumableType uint8
	name           string
	frequency      int
}

func getWeightedRandomConsumableItemType() *consumableItemInfo {
	if len(consumablesData) < TOTAL_ITEM_TYPES_NUMBER {
		panic("OH MY NOT ALL CONSUMABLES SET")
	}
	totalWeights := 0
	for i := range consumablesData {
		totalWeights += consumablesData[i].frequency
	}

	rand := rnd.Rand(totalWeights)
	for i := range consumablesData {
		if rand < consumablesData[i].frequency {
			return consumablesData[i]
		}
		rand -= consumablesData[i].frequency
	}
	panic("GWRCIT")
}
