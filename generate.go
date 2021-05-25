package main

func (g *game) generateCurrentStage() {
	g.enemies = []*enemy{}
	g.treasure = []*item{}
	stage := g.getCurrentStage()

	enemies := stage.enemies
	for _, ed := range enemies {
		g.addRandomHydra(ed)
	}
	for i := 0; i < stage.treasure; i++ {
		g.addRandomTreasure(g.currentStageNumber, 0)
	}
}

func (g *game) addRandomHydra(data *stageEnemyData) {
	element := getRandomElement(data.allowComplexElement, data.allowBossElement, false)
	if data.forceComplexElement {
		element = getRandomNonBasicElement()
	}
	if data.forceBossElement {
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
	var newWeapon *weapon
	if isWeapon {
		wData := getRandomWeaponData()
		switch wData.wtype {
		case WTYPE_SUBSTRACTOR:
			minDamage := depth - depth/2
			maxDamage := depth + depth/2 + 2
			if minDamage == 0 {
				minDamage = 1
			}
			newWeapon = &weapon{
				weaponType: WTYPE_SUBSTRACTOR,
				damage:     rnd.RandInRange(minDamage, maxDamage),
				canShoot:   rnd.OneChanceFrom(5),
			}
		default:
			additionalDamage := 0
			if rnd.OneChanceFrom(4) {
				additionalDamage++
			}
			newWeapon = &weapon{
				weaponType: wData.wtype,
				damage:     wData.minDamageForGeneration + additionalDamage,
				canShoot:   rnd.OneChanceFrom(5),
			}
		}
		hasEffect := rnd.OneChanceFrom(4)
		var eff *effect
		if hasEffect {
			eff = getRandomEffect(true, false)
		}
		return &item{
			element:      getRandomElement(true, false, true),
			asConsumable: nil,
			weaponInfo:   newWeapon,
			effect:       eff,
		}
	}

	isSpecialItem := perc < 33
	if isSpecialItem {
		return &item{
			specialName: "Ring",
			effect:      getRandomEffect(false, true),
			weaponInfo:  nil,
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
