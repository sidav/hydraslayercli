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
	g.gameScreen.renderScreen(g)
	input := g.gameScreen.readPlayerInput()
	g.parsePlayerInput(input)
}

func (g *game) parsePlayerInput(input string) {
	splitted := strings.Split(input, " ")
	// println(fmt.Sprintf("%v, %d)", splitted, len(splitted)))

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
		g.abortGame = true
		return
	}
	if splitted[0] == "hit" {
		if len(g.enemies) > 0 {
			g.performPlayerHit(g.player.items[g.currSelectedItem],
				g.enemies[g.currSelectedEnemy])
		}
		return
	}

	if splitted[0] == "cheat" {
		for i := uint8(1); i < TOTAL_ITEM_TYPES_NUMBER; i++ {
			g.player.items = append(g.player.items, &item{itemConsumableType: i})
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
		if len(splitted) == 2 {
			if splitted[1] == "all" {
				g.pickupItemNumber(-1)
			}
			takeIndex, indtype := charToIndexWithType(splitted[1][0])
			if indtype == INDEX_ENEMY_OR_TREASURE {
				g.pickupItemNumber(takeIndex)
			}
		}
		return
	}
	if splitted[0] == "drop" {
		if len(splitted) == 2 {
			dropIndex, indtype := charToIndexWithType(splitted[1][0])
			if indtype == INDEX_ITEM {
				g.dropItemNumber(dropIndex)
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
