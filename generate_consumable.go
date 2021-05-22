package main

var consumablesData = []*consumableItemInfo {
	{
		consumableType: ITEM_HEAL,
		name:           "",
		frequency:      2,
	},
	{
		consumableType: ITEM_ENCHANTER,
		name:           "",
		frequency:      2,
	},
	{
		consumableType: ITEM_DESTROY_HYDRA,
		name:           "",
		frequency:      1,
	},
	{
		consumableType: ITEM_CONFUSE_HYDRA,
		name:           "",
		frequency:      1,
	},
	{
		consumableType: ITEM_MASS_CONFUSION,
		name:           "",
		frequency:      1,
	},
	{
		consumableType: ITEM_INCREASE_HP,
		name:           "",
		frequency:      2,
	},
	{
		consumableType: ITEM_STRENGTH,
		name:           "",
		frequency:      2,
	},
	{
		consumableType: ITEM_CHANGE_ELEMENT,
		name:           "",
		frequency:      1,
	},
	{
		consumableType: ITEM_UNELEMENT_ENEMIES,
		name:           "",
		frequency:      1,
	},
	{
		consumableType: ITEM_DECAPITATION,
		name:           "",
		frequency:      1,
	},
	{
		consumableType: ITEM_IMPROVE_COOLDOWN,
		name:           "",
		frequency:      1,
	},
	{
		consumableType: ITEM_AMMO,
		name:           "",
		frequency:      2,
	},
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
