package main

func (g *game) generateCurrentStage() {
	for i := 0; i < StageInfo[g.currentStageNumber].enemies; i++ {
		g.addRandomHydra(g.currentStageNumber, 0)
	}
	for i := 0; i < StageInfo[g.currentStageNumber].treasure; i++ {
		g.addRandomTreasure(g.currentStageNumber, 0)
	}
}

func (g *game) addRandomHydra(depth, difficulty int) {
	g.enemies = append(g.enemies, g.generateHydra(depth, 1))
}

func (g *game) generateHydra(depth, difficulty int) *enemy {
	minHeads := depth+1
	maxHeads := minHeads + depth/2 + 2
	return &enemy{
		name:    "hydra",
		heads:   rnd.RandInRange(minHeads, maxHeads),
		element: getRandomElement(),
	}
}

func (g *game) addRandomTreasure(depth, difficulty int) {
	g.treasure = append(g.treasure, g.generateTreasure(depth))
}

func (g *game) generateTreasure(depth int) *item {
	isWeapon := rnd.OneChanceFrom(2)
	if isWeapon {
		minDamage := depth-depth/2
		maxDamage := depth+depth/2+2
		if minDamage == 0 {
			minDamage = 1
		}
		newWeapon := &weapon{
			weaponType: WTYPE_SUBSTRACTOR,
			element:    0,
			damage:     rnd.RandInRange(minDamage, maxDamage),
		}
		weaponTypePercent := rnd.RandomPercent()
		if weaponTypePercent < 25 {
			newWeapon.weaponType = WTYPE_DIVISOR
			newWeapon.damage = 2
		} else {

		}

		return &item{
			itemType:   0,
			weaponInfo: newWeapon,
		}
	}

	return &item{
		itemType:   uint8(rnd.RandInRange(1, TOTAL_ITEM_TYPES_NUMBER-1)),
		weaponInfo: nil,
	}
}
