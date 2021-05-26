package main

import "fmt"

func (g *game) performUseAction(usedIndex int, usedType INDEXTYPE, targetIndex int, targetType INDEXTYPE) {
	var usedItem *item
	var usedFromGround bool
	if usedType != INDEX_LETTER {
		if usedIndex < len(g.treasure) && len(g.enemies) == 0 {
			usedItem = g.treasure[usedIndex]
			usedFromGround = true
		} else {
			g.setLogMessage("Wat?")
			return
		}
	} else {
		if len(g.player.items) > usedIndex {
			usedItem = g.player.items[usedIndex]
		}
	}
	if targetIndex == -1 && usedItem != nil {
		g.justUseItem(usedItem, usedFromGround)
		return
	}

	var targetEnemy *enemy
	if targetType == INDEX_NUMBER && len(g.enemies) > targetIndex {
		targetEnemy = g.enemies[targetIndex]
	}
	if usedItem != nil && targetEnemy != nil {
		g.useItemOnEnemy(usedItem, targetEnemy)
		return
	}

	var targetItem *item
	if targetType == INDEX_LETTER && len(g.player.items) > targetIndex {
		targetItem = g.player.items[targetIndex]
	}
	if targetType == INDEX_NUMBER && len(g.treasure) > targetIndex {
		targetItem = g.treasure[targetIndex]
	}
	if usedItem != nil && targetItem != nil {
		g.useItemOnItem(usedItem, targetItem, usedFromGround, 999)
	}
}

func (g *game) justUseItem(item *item, usedFromGround bool) {
	if item.brand != nil {
		if !item.brand.canBeUsed {
			g.appendToLogMessage("%s can't be used yet.", item.getName())
		}
		item.applyActiveEffect(g)
		return
	}
	if item.asConsumable == nil {
		g.setLogMessage("And what? How did you want to use %s?", item.getName())
		return
	}
	switch item.asConsumable.consumableType {
	case ITEM_HEAL:
		g.setLogMessage("You sniff %s and feel good.", item.getName())
		g.player.hp = g.player.maxhp
	case ITEM_INCREASE_HP:
		g.setLogMessage("%s makes you feel amazing!", item.getName())
		g.player.maxhp += 2
		g.player.hp = g.player.maxhp
	case ITEM_DECAPITATION:
		for _, e := range g.enemies {
			e.heads -= e.heads/2
		}
		g.setLogMessage("Wave of unholy power cuts all enemies in half!")
	case ITEM_UNELEMENT_ENEMIES:
		for _, e := range g.enemies {
			e.element = elementsData[0]
		}
		g.setLogMessage("All enemies lose their magic!")
	case ITEM_STRENGTH:
		g.setLogMessage("%s makes you feel stronger!", item.getName())
		g.player.maxItems += 1
	case ITEM_MASS_CONFUSION:
		g.setLogMessage("Enemies freeze in confusion!")
		for _, enemy := range g.enemies {
			enemy.statuses = append(enemy.statuses, &statusEffect{
				statusType:     STATUS_CONFUSED,
				turnsRemaining: 4,
			})
		}
	case ITEM_MERGE_HYDRAS_INTO_ONE:
		g.setLogMessage(fmt.Sprintf("Magic quickly surges around enemies... and creates the new one!"))
		randomEnemy := rnd.Rand(len(g.enemies))
		g.appendToLogMessage(fmt.Sprintf("%s is now", g.enemies[randomEnemy].getName()))
		totalHeadsAdd := 0
		for _, enemy := range g.enemies {
			totalHeadsAdd += enemy.heads
			enemy.heads = 0
		}
		g.enemies[randomEnemy].heads += totalHeadsAdd
		g.appendToLogMessage(fmt.Sprintf("%s!", g.enemies[randomEnemy].getName()))

	default:
		g.setLogMessage("ERROR: ADD SIMPLE USAGE OF %s.", item.getName())
		return
	}
	if usedFromGround {
		g.removeTreasure(item)
	} else {
		g.player.spendItem(item, g)
	}
	g.turnMade = true
	g.allEnemiesSkipTurn = true
}

func (g *game) useItemOnEnemy(item *item, enemy *enemy) {
	if item.asConsumable == nil {
		g.setLogMessage("And what? How did you want to use %s?", item.getName())
		return
	}
	switch item.asConsumable.consumableType {
	case ITEM_HEAL:
		g.setLogMessage("Use %s on enemy? Srsly?", item.getName())
		return
	case ITEM_INCREASE_HP:
		g.setLogMessage("Are you nuts?")
		return
	case ITEM_DESTROY_HYDRA:
		g.setLogMessage("The magic obliterates poor %s!", enemy.getName())
		enemy.heads = 0
	case ITEM_CONFUSE_HYDRA:
		g.setLogMessage("The %s starts behaving like crazy.", enemy.getName())
		enemy.statuses = append(enemy.statuses, &statusEffect{
			statusType:     STATUS_CONFUSED,
			turnsRemaining: 6,
		})
	case ITEM_CHANGE_ELEMENT_RANDOM:
		g.setLogMessage("You use %s on %s, making it into ", item.getName(), enemy.getName())
		enemy.element = getRandomElement(true, false, false)
		g.appendToLogMessage("%s.", enemy.getName())
	case ITEM_CHANGE_ELEMENT_SPECIFIC:
		g.setLogMessage("You use %s on %s, making it into ", item.getName(), enemy.getName())
		enemy.element = item.auxiliaryElement
		g.appendToLogMessage("%s.", enemy.getName())
	default:
		g.setLogMessage("ERROR: ADD USAGE %s ON ENEMY.", item.getName())
		return
	}
	g.player.spendItem(item, g)
	g.allEnemiesSkipTurn = true
	g.turnMade = true
}

func (g *game) useItemOnItem(item, targetItem *item, usedFromGround bool, count int) {
	if item.asConsumable == nil {
		g.setLogMessage("And what? How did you want to use %s?", item.getName())
		return
	}
	switch item.asConsumable.consumableType {
	case ITEM_HEAL:
		g.setLogMessage("Use %s on %s? But how?", item.getName(), targetItem.getName())
	case ITEM_ENCHANTER:
		if targetItem.weaponInfo == nil {
			g.setLogMessage("But %s is not a weapon!", targetItem.getName())
			return
		}
		g.setLogMessage("You use %s on %s, making it into ", item.getName(), targetItem.getName())
		targetItem.weaponInfo.damage++
		g.appendToLogMessage("%s.", targetItem.getName())
	case ITEM_CHANGE_ELEMENT_RANDOM:
		g.setLogMessage("You use %s on %s, making it into ", item.getName(), targetItem.getName())
		targetItem.element = getRandomElement(true, false, true)
		g.appendToLogMessage("%s.", targetItem.getName())
	case ITEM_CHANGE_ELEMENT_SPECIFIC:
		g.setLogMessage("You use %s on %s, making it into ", item.getName(), targetItem.getName())
		targetItem.element = item.auxiliaryElement
		g.appendToLogMessage("%s.", targetItem.getName())
	case ITEM_BRANDING_RANDOM:
		g.setLogMessage("You use %s on %s, making it into ", item.getName(), targetItem.getName())
		targetItem.brand = getRandomBrand(targetItem.isWeapon(), !targetItem.isWeapon())
		g.appendToLogMessage("%s.", targetItem.getName())
	case ITEM_BRANDING_SPECIFIC:
		if targetItem.isWeapon() && !item.auxiliaryBrand.getStaticData().canBeOnWeapon {
			g.setLogMessage("Better try using this on another item!")
			return
		}
		if !targetItem.isWeapon() && !item.auxiliaryBrand.getStaticData().canBeOnRing {
			g.setLogMessage("Better try using this on a weapon!")
			return
		}
		g.setLogMessage("You use %s on %s, making it into ", item.getName(), targetItem.getName())
		targetItem.brand = item.auxiliaryBrand
		g.appendToLogMessage("%s.", targetItem.getName())
	case ITEM_AMMO:
		if !targetItem.hasEffect() || !targetItem.brand.isChargeable() {
			g.setLogMessage("But %s can't be charged!", targetItem.getName())
			return
		}
		charges := item.count
		if count < charges {
			charges = count
		}
		g.setLogMessage("You charge %s with %d additional charges.", targetItem.getName(), charges)
		targetItem.brand.charges += charges
		for i := 0; i < charges; i++ {
			g.player.spendItem(item, g)
		}
		g.turnMade = true
		return
	case ITEM_IMPROVE_MAGIC:
		if targetItem.brand == nil {
			g.setLogMessage("But %s can't be improved!", targetItem.getName())
			return
		}
		g.setLogMessage("You use %s on %s, making it into ", item.getName(), targetItem.getName())
		if targetItem.brand.activatesEach > 1 && targetItem.brand.getStaticData().defaultActivatesEach > 0 {
			targetItem.brand.activatesEach--
		}
		if targetItem.brand.getStaticData().defaultAdditionalDamage > 0 {
			targetItem.brand.additionalDamage++
		}
		g.appendToLogMessage("%s.", targetItem.getName())
	default:
		g.setLogMessage("ERROR: ADD USAGE %s ON ITEM.", item.getName())
		return
	}
	if usedFromGround {
		g.removeTreasure(item)
	} else {
		g.player.spendItem(item, g)
	}
	g.turnMade = true
}

func (g *game) pickupItemNumber(i int) {
	if i == -1 { // pick up all
		if len(g.treasure) + len(g.player.items) > g.player.maxItems {
			g.setLogMessage("You can't pick up everything!")
			return
		}
		g.setLogMessage("You pick up everything: ")
		for i := 0; i < len(g.treasure); i++ {
			if i > 0 {
				g.currLog += ", "
			}
			g.player.addItem(g.treasure[i])
			g.currLog +=  g.treasure[i].getName()
		}
		g.currLog += "."
		g.treasure = []*item{}
		return
	}
	if i < len(g.treasure) {
		if len(g.player.items) >= g.player.maxItems && !(g.player.hasAmmo() && g.treasure[i].isAmmo()) {
			g.setLogMessage("You are overburdened!")
			return
		}
		g.player.addItem(g.treasure[i])
		g.setLogMessage("You pick up the %s.", g.treasure[i].getName())
		g.removeTreasure(g.treasure[i])
	}
}

func (g *game) removeTreasure(t *item) {
	for i, tt := range g.treasure {
		if tt == t {
			g.treasure = append(g.treasure[:i], g.treasure[i+1:]...)
			return
		}
	}
}

func (g *game) dropItemNumber(i int) {
	if i < len(g.player.items) {
		g.treasure = append(g.treasure, g.player.items[i])
		g.setLogMessage("You drop the %s.", g.player.items[i].getName())
		g.player.items  = append(g.player.items[:i], g.player.items[i+1:]...)
	}
}
