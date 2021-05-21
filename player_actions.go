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
	if st == INDEX_ENEMY && len(g.currentEnemies) > targetIndex {
		targetEnemy = g.currentEnemies[targetIndex]
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
	switch item.itemType {
	case ITEM_HEAL:
		g.currLog = fmt.Sprintf("You sniff %s and feel good.", item.getName())
		g.player.hp = g.player.maxhp
	default:
		g.currLog = fmt.Sprintf("No valid target for %s.", item.getName())
		return
	}
	g.player.spendItem(item)
	g.turnMade = true
}

func (g *game) useItemOnEnemy(item *item, enemy *enemy) {
	switch item.itemType {
	case ITEM_HEAL:
		g.currLog = fmt.Sprintf("Use %s on enemy? Srsly?", item.getName())
		return
	case ITEM_DESTROY_HYDRA:
		g.currLog = fmt.Sprintf("The magic obliterates poor %s!", enemy.getName())
		enemy.heads = 0
	default:
		g.currLog = fmt.Sprintf("No valid target for %s.", item.getName())
		return
	}
	g.player.spendItem(item)
	g.turnMade = true
}

func (g *game) useItemOnItem(item, targetItem *item) {
	switch item.itemType {
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
	default:
		g.currLog = fmt.Sprintf("No valid target for %s.", item.getName())
		return
	}
	g.player.spendItem(item)
	g.turnMade = true
}
