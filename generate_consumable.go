package main

var _consumableWeights = []int {
	ITEM_HEAL: 2,
	ITEM_ENCHANTER: 2,
	ITEM_DESTROY_HYDRA: 1,
	ITEM_CONFUSE_HYDRA: 1,
	ITEM_INCREASE_HP: 2,
	ITEM_STRENGTH: 2,
	ITEM_CHANGE_ELEMENT: 1,
}

func getWeightedRandomConsumableItemType() uint8 {
	totalWeights := 0
	for i := range _consumableWeights {
		totalWeights += _consumableWeights[i]
	}

	rand := rnd.Rand(totalWeights)
	for i := range _consumableWeights {
		if rand < _consumableWeights[i] {
			return uint8(i)
		}
		rand -= _consumableWeights[i]
	}
	panic("GWRCIT")
}
