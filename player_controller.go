package main

import (
	"fmt"
	"strings"
)

type INDEXTYPE uint8

const (
	INDEX_ENEMY_OR_TREASURE = iota
	INDEX_ITEM
	INDEX_WRONG
)

func charToIndexWithType(char byte) (int, INDEXTYPE) {
	if char >= 'a' && char <= 'z' {
		return int(255 - ('a' - char - 1)), INDEX_ITEM
	}
	if char >= '0' && char <= '9' {
		return int(char - '1'), INDEX_ENEMY_OR_TREASURE
	}
	return 0, INDEX_WRONG
}

func (g *game) playerTurn() {
	screen.renderScreen(g)
	input := screen.readPlayerInput()
	g.parsePlayerInput(input)
}

func (g *game) parsePlayerInput(input string) {
	splitted := strings.Split(input, " ")
	// remove all trailing spaces
	for i := len(splitted) - 1; i > 0; i-- {
		if splitted[i] == "" {
			splitted = append(splitted[:i], splitted[i+1:]...)
		}
	}

	if splitted[0] == "" {
		g.currLog = ""
		return
	}
	if splitted[0] == "wait" {
		g.setLogMessage("You do nothing.")
		g.turnMade = true
		return
	}
	if splitted[0] == "exit" {
		abortGame = true
		return
	}
	if splitted[0] == "show" {
		g.showShortCombatDescription = !g.showShortCombatDescription
		return
	}
	if splitted[0] == "hit" {
		if len(g.enemies) > 0 {
			g.performPlayerHit(g.player.items[g.currSelectedItem],
				g.enemies[g.currSelectedEnemy])
		}
		return
	}

	if splitted[0] == "info" || splitted[0] == "help" {
		if len(splitted) == 1 {
			g.setLogMessage("HELP!")
		} else {
			ind, itype := charToIndexWithType(splitted[1][0])
			if itype == INDEX_ITEM {
				if ind < len(g.player.items) {
					g.setLogMessage(g.player.items[ind].getInfo())
					return
				}
			}
			if itype == INDEX_ENEMY_OR_TREASURE {
				if len(g.enemies) > 0 && ind < len(g.enemies) {
					g.setLogMessage(g.enemies[ind].getInfo())
					return
				}
				if len(g.treasure) > 0 && ind < len(g.treasure) {
					g.setLogMessage(g.treasure[ind].getInfo())
					return
				}
			}
		}
		return
	}

	if splitted[0] == "shoot" || splitted[0] == "fire" {
		if len(g.enemies) > 0 {
			g.performPlayerShoot(g.player.items[g.currSelectedItem],
				g.enemies[g.currSelectedEnemy])
		}
		return
	}

	if splitted[0] == "cheati" {
		for i := uint8(1); i < TOTAL_ITEM_TYPES_NUMBER; i++ {
			g.player.maxItems++
			newItem := &item{asConsumable: consumablesData[i]}
			if newItem.isAmmo() {
				newItem.count++
			}
			g.player.items = append(g.player.items, newItem)
		}
		return
	}
	if splitted[0] == "cheatt" {
		for i := 0; i < 10; i++ {
			g.player.addItem(g.generateTreasure(auxrnd.RandInRange(1, 15)))
		}
		return
	}

	if splitted[0] == "move" {
		if len(g.enemies) == 0 {
			g.stageFinished = true
		} else {
			g.currLog = "You can't move just yet!"
		}
		return
	}

	if splitted[0] == "take" {
		if len(splitted) == 1 || splitted[1] == "all" {
			g.pickupItemNumber(-1)
			return
		}
		for i := 1; i < len(splitted); i++ {
			takeIndex, indtype := charToIndexWithType(splitted[i][0])
			if indtype == INDEX_ENEMY_OR_TREASURE {
				g.pickupItemNumber(takeIndex - i + 1)
			}
		}
		return
	}
	if splitted[0] == "drop" {
		for i := 1; i < len(splitted); i++ {
			dropIndex, indtype := charToIndexWithType(splitted[i][0])
			if indtype == INDEX_ITEM {
				g.dropItemNumber(dropIndex - i + 1)
			}
		}
		return
	}

	if splitted[0] == "use" {
		if len(splitted) == 2 {
			useFI, useFType := charToIndexWithType(splitted[1][0])
			g.performUseAction(useFI, useFType, -1, INDEX_WRONG)
		}
		if len(splitted) >= 3 {
			useFI, useFType := charToIndexWithType(splitted[1][0])
			useSI, useSType := charToIndexWithType(splitted[2][0])
			g.performUseAction(useFI, useFType, useSI, useSType)
		}
		return
	}

	if len(splitted[0]) == 1 {
		index, indtype := charToIndexWithType(splitted[0][0])
		// select item
		if indtype == INDEX_ITEM {
			if len(g.player.items) > index {
				// gs.currLog = g.player.weapons[itemnum].getName()
				if g.player.items[index].isWeapon() {
					g.currSelectedItem = index
					g.currLog = ""
				} else {
					// Item description!..
				}
			}
			return
		}

		//select enemy
		if indtype == INDEX_ENEMY_OR_TREASURE {
			// print(splitted[0][0], '1', '0', enemynum)
			if len(g.enemies) > index {
				// gs.currLog = g.enemies[enemynum].getName()
				g.currSelectedEnemy = index
				g.currLog = ""
			}
			return
		}
	}
	// nothing parsed
	g.currLog = fmt.Sprintf("Command \"%s\" not recognized.", splitted[0])
}
