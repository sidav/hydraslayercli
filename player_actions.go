package main

import "fmt"

func (g *game) performUseAction(usedIndex int, ft INDEXTYPE, targetIndex int, st INDEXTYPE) {
	if ft != INDEX_ITEM {
		g.currLog = "Wat?"
		return
	}
	var usedItem *item
	if len(g.player.items) > usedIndex {
		usedItem = g.player.items[usedIndex]
	}
	if targetIndex == -1 {
		g.justUseItem(usedItem)
		return
	}

	var targetEnemy *enemy
	if st == INDEX_ENEMY_OR_TREASURE && len(g.enemies) > targetIndex {
		targetEnemy = g.enemies[targetIndex]
	}
	if usedItem != nil && targetEnemy != nil {
		g.useItemOnEnemy(usedItem, targetEnemy)
		return
	}

	var targetItem *item
	if st == INDEX_ITEM && len(g.player.items) > targetIndex {
		targetItem = g.player.items[targetIndex]
	}
	if usedItem != nil && targetItem != nil {
		g.useItemOnItem(usedItem, targetItem)
	}
}

func (g *game) justUseItem(item *item) {
	switch item.itemConsumableType {
	case ITEM_HEAL:
		g.currLog = fmt.Sprintf("You sniff %s and feel good.", item.getName())
		g.player.hp = g.player.maxhp
	case ITEM_INCREASE_HP:
		g.currLog = fmt.Sprintf("%s makes you feel amazing!", item.getName())
		g.player.maxhp += 1
		g.player.hp = g.player.maxhp
	case ITEM_STRENGTH:
		g.currLog = fmt.Sprintf("%s makes you feel stronger!", item.getName())
		g.player.maxItems += 1
	default:
		g.currLog = fmt.Sprintf("ERROR: ADD SIMPLE USAGE OF %s.", item.getName())
		return
	}
	g.player.spendItem(item)
	g.turnMade = true
}

func (g *game) useItemOnEnemy(item *item, enemy *enemy) {
	switch item.itemConsumableType {
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
			turnsRemaining: 3,
		})
	case ITEM_CHANGE_ELEMENT:
		g.currLog = fmt.Sprintf("You use %s on %s, making it into ", item.getName(), enemy.getName())
		enemy.element = getRandomElement()
		g.currLog += fmt.Sprintf("%s.", enemy.getName())
	default:
		g.currLog = fmt.Sprintf("ERROR: ADD USAGE %s ON ENEMY.", item.getName())
		return
	}
	g.player.spendItem(item)
	g.turnMade = true
}

func (g *game) useItemOnItem(item, targetItem *item) {
	switch item.itemConsumableType {
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
		targetItem.element = getRandomElement()
		g.currLog += fmt.Sprintf("%s.", targetItem.getName())
	default:
		g.currLog = fmt.Sprintf("ERROR: ADD USAGE %s ON ITEM.", item.getName())
		return
	}
	g.player.spendItem(item)
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
			g.player.items = append(g.player.items, g.treasure[i])
			g.currLog +=  g.treasure[i].getName()
		}
		g.currLog += "."
		g.treasure = []*item{}
		return
	}
	if i < len(g.treasure) {
		if len(g.player.items) >= g.player.maxItems {
			g.currLog = fmt.Sprintf("You are overburdened!")
			return
		}
		g.player.items = append(g.player.items, g.treasure[i])
		g.currLog = fmt.Sprintf("You pick up the %s.", g.treasure[i].getName())
		g.treasure = append(g.treasure[:i], g.treasure[i+1:]...)
	}
}

func (g *game) dropItemNumber(i int) {
	if i < len(g.player.items) {
		g.treasure = append(g.treasure, g.player.items[i])
		g.currLog = fmt.Sprintf("You drop the %s.", g.player.items[i].getName())
		g.player.items  = append(g.player.items[:i], g.player.items[i+1:]...)
	}
}
