package main

import (
	"strings"
)

type INDEXTYPE uint8

const (
	INDEX_ENEMY = iota
	INDEX_ITEM
	INDEX_WRONG
)

func charToIndexWithType(char byte) (int, INDEXTYPE) {
	if char >= 'a' && char <= 'z' {
		return int(255 - ('a' - char - 1)), INDEX_ITEM
	}
	if char >= '0' && char <= '9' {
		return int(char - '1'), INDEX_ENEMY
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
		g.currLog = "Huh?"
		return
	}
	if splitted[0] == "exit" {
		g.exit = true
		return
	}
	if splitted[0] == "hit" {
		if len(g.currentEnemies) > 0 {
			g.performPlayerHit(g.player.items[g.currSelectedItem].weaponInfo,
				g.currentEnemies[g.currSelectedEnemy])
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
	}

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
	}

	//select enemy
	if indtype == INDEX_ENEMY {
		// print(splitted[0][0], '1', '0', enemynum)
		if len(g.currentEnemies) > index {
			// gs.currLog = g.currentEnemies[enemynum].getName()
			g.currSelectedEnemy = index
			g.currLog = ""
		}
	}
}
