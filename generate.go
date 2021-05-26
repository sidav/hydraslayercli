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
			}
		default:
			additionalDamage := 0
			if rnd.OneChanceFrom(4) {
				additionalDamage++
			}
			newWeapon = &weapon{
				weaponType: wData.wtype,
				damage:     wData.minDamageForGeneration + additionalDamage,
			}
		}
		hasEffect := rnd.OneChanceFrom(4)
		var eff *brand
		if hasEffect {
			eff = getRandomBrand(true, false)
		}
		return &item{
			element:      getRandomElement(true, false, true),
			asConsumable: nil,
			weaponInfo:   newWeapon,
			brand:        eff,
		}
	}

	isSpecialItem := perc < 33
	if isSpecialItem {
		return &item{
			specialName: "Ring",
			brand:       getRandomBrand(false, true),
			weaponInfo:  nil,
		}
	}

	return getWeightedRandomConsumableItem()
}
