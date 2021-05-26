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
	ITEM_IMPROVE_MAGIC
	ITEM_GAIN_EFFECT
	ITEM_AMMO
	ITEM_MERGE_HYDRAS_INTO_ONE
	TOTAL_ITEM_TYPES_NUMBER // for generators
)

var consumablesData = []*consumableItemInfo{
	{
		consumableType: ITEM_HEAL,
		name:           "Healing powder",
		frequency:      2,
		info:           "Can be used to recover HP.",
	},
	{
		consumableType: ITEM_ENCHANTER,
		name:           "Glyph of enchant weapon",
		frequency:      3,
		info:           "Can be used to increase weapon damage.",
	},
	{
		consumableType: ITEM_DESTROY_HYDRA,
		name:           "Glyph of destroy hydra",
		frequency:      1,
		info:           "Can be used to destroy hydra.",
	},
	{
		consumableType: ITEM_CONFUSE_HYDRA,
		name:           "Glyph of confuse hydra",
		frequency:      1,
		info:           "Can be used to confuse hydra.",
	},
	{
		consumableType: ITEM_MASS_CONFUSION,
		name:           "Scroll of confused party",
		info:           "Can be used to confuse all enemies.",
		frequency:      1,
	},
	{
		consumableType: ITEM_INCREASE_HP,
		name:           "Potion of vitality",
		frequency:      3,
	},
	{
		consumableType: ITEM_STRENGTH,
		name:           "Potion of strength",
		info:           "Permanently increases your inventory size.",
		frequency:      3,
	},
	{
		consumableType: ITEM_CHANGE_ELEMENT,
		name:           "Glyph of change element",
		info:           "Changes hydra's or item's element.",
		frequency:      1,
	},
	{
		consumableType: ITEM_UNELEMENT_ENEMIES,
		name:           "Scroll of nullification",
		info:           "Permanently makes all present enemies non-elemental.",
		frequency:      1,
	},
	{
		consumableType: ITEM_DECAPITATION,
		name:           "Scroll of mass bisection",
		info:           "Divides all heads of presen hydras by 2.",
		frequency:      1,
	},
	{
		consumableType: ITEM_IMPROVE_MAGIC,
		name:           "Glyph of improve brand",
		info:           "Can be used to improve branded (magic) items",
		frequency:      1,
	},
	{
		consumableType: ITEM_GAIN_EFFECT,
		name:           "Glyph of imbuing",
		info:           "Used to imbue a random effect onto item.",
		frequency:      1,
	},
	{
		consumableType: ITEM_MERGE_HYDRAS_INTO_ONE,
		name:           "Scroll of merge hydras",
		info:           "Used to make a single hydra from many.",
		frequency:      1,
	},
	{
		consumableType: ITEM_AMMO,
		name:           "Charge crystals",
		info:           "Used with chargeable items.",
		frequency:      2,
	},
}

type consumableItemInfo struct {
	consumableType uint8
	name, info     string
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
