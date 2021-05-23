package main

func (g *game) generateCurrentStage() {
	g.enemies = []*enemy{}
	g.treasure = []*item{}
	for _, ed := range StageInfo[g.currentStageNumber].enemies {
		g.addRandomHydra(ed)
	}
	for i := 0; i < StageInfo[g.currentStageNumber].treasure; i++ {
		g.addRandomTreasure(g.currentStageNumber, 0)
	}
}

func (g *game) addRandomHydra(data stageEnemyData) {
	element := getRandomElement(data.allowComplexElement, data.allowSpecialElement)
	if data.forceComplexElement {
		element = getRandomNonBasicElement()
	}
	if data.forceSpecialElement {
		element = getRandomSpecialElement()
	}
	hydra := &enemy{
		name:    "hydra",
		heads:   rnd.RandInRange(data.minHeads, data.maxHeads),
		element: element,
	}
	g.enemies = append(g.enemies, hydra)
}

func (g *game) addRandomTreasure(depth, difficulty int) {
	g.treasure = append(g.treasure, g.generateTreasure(depth))
}

func (g *game) generateTreasure(depth int) *item {
	perc := rnd.RandomPercent()
	isWeapon := perc < 25
	if isWeapon {
		minDamage := depth - depth/2
		maxDamage := depth + depth/2 + 2
		if minDamage == 0 {
			minDamage = 1
		}
		newWeapon := &weapon{
			weaponType: WTYPE_SUBSTRACTOR,
			damage:     rnd.RandInRange(minDamage, maxDamage),
			canShoot: rnd.OneChanceFrom(5),
		}
		weaponTypePercent := rnd.RandomPercent()
		if weaponTypePercent < 25 {
			newWeapon.weaponType = WTYPE_DIVISOR
			newWeapon.damage = 2
		} else {

		}

		return &item{
			element:      getRandomElement(true, false),
			asConsumable: nil,
			weaponInfo:   newWeapon,
		}
	}
	isSpecialItem := perc < 33
	if isSpecialItem {
		return &item{
			specialName: "Ring",
			effect: &effect{
				effectType:    getRandomPassiveEffect(),
				activatesEach: 4,
			},
			weaponInfo: nil,
		}
	}

	item := &item{
		asConsumable: getWeightedRandomConsumableItemType(),
		weaponInfo:   nil,
	}
	if item.asConsumable.consumableType == ITEM_AMMO {
		item.count = rnd.RandInRange(1, 3)
	}
	return item
}
