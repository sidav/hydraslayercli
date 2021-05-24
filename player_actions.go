package main

import "fmt"

func (g *game) performUseAction(usedIndex int, usedType INDEXTYPE, targetIndex int, targetType INDEXTYPE) {
	var usedItem *item
	var usedFromGround bool
	if usedType != INDEX_ITEM {
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
	if targetType == INDEX_ENEMY_OR_TREASURE && len(g.enemies) > targetIndex {
		targetEnemy = g.enemies[targetIndex]
	}
	if usedItem != nil && targetEnemy != nil {
		g.useItemOnEnemy(usedItem, targetEnemy)
		return
	}

	var targetItem *item
	if targetType == INDEX_ITEM && len(g.player.items) > targetIndex {
		targetItem = g.player.items[targetIndex]
	}
	if targetType == INDEX_ENEMY_OR_TREASURE && len(g.treasure) > targetIndex {
		targetItem = g.treasure[targetIndex]
	}
	if usedItem != nil && targetItem != nil {
		g.useItemOnItem(usedItem, targetItem, usedFromGround)
	}
}

func (g *game) justUseItem(item *item, usedFromGround bool) {
	if item.effect != nil {
		if !item.effect.canBeUsed {
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
		g.currLog = fmt.Sprintf("You sniff %s and feel good.", item.getName())
		g.player.hp = g.player.maxhp
	case ITEM_INCREASE_HP:
		g.currLog = fmt.Sprintf("%s makes you feel amazing!", item.getName())
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
		g.currLog = fmt.Sprintf("%s makes you feel stronger!", item.getName())
		g.player.maxItems += 1
	case ITEM_MASS_CONFUSION:
		g.currLog = fmt.Sprintf("Enemies freeze in confusion!")
		for _, enemy := range g.enemies {
			enemy.statuses = append(enemy.statuses, &statusEffect{
				statusType:     STATUS_CONFUSED,
				turnsRemaining: 4,
			})
		}
	default:
		g.currLog = fmt.Sprintf("ERROR: ADD SIMPLE USAGE OF %s.", item.getName())
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
		g.currLog = fmt.Sprintf("Use %s on enemy? Srsly?", item.getName())
		return
	case ITEM_INCREASE_HP:
		g.currLog = fmt.Sprintf("Are you nuts?")
		return
	case ITEM_DESTROY_HYDRA:
		g.currLog = fmt.Sprintf("The magic obliterates poor %s!", enemy.getName())
		enemy.heads = 0
	case ITEM_CONFUSE_HYDRA:
		g.currLog = fmt.Sprintf("The %s starts behaving like crazy.", enemy.getName())
		enemy.statuses = append(enemy.statuses, &statusEffect{
			statusType:     STATUS_CONFUSED,
			turnsRemaining: 6,
		})
	case ITEM_CHANGE_ELEMENT:
		g.currLog = fmt.Sprintf("You use %s on %s, making it into ", item.getName(), enemy.getName())
		enemy.element = getRandomElement(true, false)
		g.currLog += fmt.Sprintf("%s.", enemy.getName())
	default:
		g.currLog = fmt.Sprintf("ERROR: ADD USAGE %s ON ENEMY.", item.getName())
		return
	}
	g.player.spendItem(item, g)
	g.allEnemiesSkipTurn = true
	g.turnMade = true
}

func (g *game) useItemOnItem(item, targetItem *item, usedFromGround bool) {
	if item.asConsumable == nil {
		g.setLogMessage("And what? How did you want to use %s?", item.getName())
		return
	}
	switch item.asConsumable.consumableType {
	case ITEM_HEAL:
		g.currLog = fmt.Sprintf("Use %s on %s? But how?", item.getName(), targetItem.getName())
	case ITEM_ENCHANTER:
		if targetItem.weaponInfo == nil {
			g.currLog = fmt.Sprintf("But %s is not a weapon!", targetItem.getName())
			return
		}
		g.currLog = fmt.Sprintf("You use %s on %s, making it into ", item.getName(), targetItem.getName())
		targetItem.weaponInfo.damage++
		g.currLog += fmt.Sprintf("%s.", targetItem.getName())
	case ITEM_CHANGE_ELEMENT:
		g.currLog = fmt.Sprintf("You use %s on %s, making it into ", item.getName(), targetItem.getName())
		targetItem.element = getRandomElement(true, false)
		g.currLog += fmt.Sprintf("%s.", targetItem.getName())
	case ITEM_GAIN_EFFECT:
		g.currLog = fmt.Sprintf("You use %s on %s, making it into ", item.getName(), targetItem.getName())
		targetItem.effect = getRandomEffect(targetItem.isWeapon(), !targetItem.isWeapon())
		g.currLog += fmt.Sprintf("%s.", targetItem.getName())
	case ITEM_IMPROVE_MAGIC:
		if targetItem.effect == nil {
			g.setLogMessage("But %s can't be improved!", targetItem.getName())
			return
		}
		g.currLog = fmt.Sprintf("You use %s on %s, making it into ", item.getName(), targetItem.getName())
		if targetItem.effect.activatesEach > 1 && targetItem.effect.getStaticData().defaultActivatesEach > 0 {
			targetItem.effect.activatesEach--
		}
		if targetItem.effect.getStaticData().defaultAdditionalDamage > 0 {
			targetItem.effect.additionalDamage++
		}
		g.currLog += fmt.Sprintf("%s.", targetItem.getName())
	default:
		g.currLog = fmt.Sprintf("ERROR: ADD USAGE %s ON ITEM.", item.getName())
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
			g.currLog = fmt.Sprintf("You can't pick up everything!")
			return
		}
		g.currLog = fmt.Sprintf("You pick up everything: ")
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
			g.currLog = fmt.Sprintf("You are overburdened!")
			return
		}
		g.player.addItem(g.treasure[i])
		g.currLog = fmt.Sprintf("You pick up the %s.", g.treasure[i].getName())
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
		g.currLog = fmt.Sprintf("You drop the %s.", g.player.items[i].getName())
		g.player.items  = append(g.player.items[:i], g.player.items[i+1:]...)
	}
}
