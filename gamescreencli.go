package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type gameScreen struct {
	input                               string
	currLog                             string
	currSelectedItem, currSelectedEnemy uint8
}

func (gs *gameScreen) renderScreen(g *game) {
	fmt.Println("\033[2J") // linux only!
	println("Turn ", g.currentTurn)
	println("")
	println("You see here enemies:")
	for i, e := range g.currentEnemies {
		selectStr := "   "
		if uint8(i) == gs.currSelectedEnemy {
			selectStr = "-> "
		}
		println(fmt.Sprintf("%s%d: %s", selectStr, i+1, e.getName()))
	}
	println(fmt.Sprintf("You have %d/%d hp and:", g.player.hp, g.player.maxhp))
	for i, w := range g.player.weapons {
		selectStr := "  "
		if uint8(i) == gs.currSelectedItem {
			selectStr = "->"
		}
		println(fmt.Sprintf("%s %c: %s", selectStr, 'A'+i, w.getName()))
	}
	println(gs.currLog)
	print("Your action?\n> ")
	gs.doInput(g)
}

func (gs *gameScreen) doInput(g *game) {
	gs.currLog = ""
	reader := bufio.NewReader(os.Stdin)
	gs.input, _ = reader.ReadString('\n')
	gs.input = strings.Trim(gs.input, " \n")

	splitted := strings.Split(gs.input, " ")
	// println(fmt.Sprintf("%v, %d)", splitted, len(splitted)))

	if len(splitted) == 1 {
		if splitted[0] == "" {
			gs.currLog = "Huh?"
			return
		}
		if splitted[0] == "exit" {
			g.exit = true
			return
		}
		if splitted[0] == "hit" {
			if len(g.currentEnemies) > 0 {
				g.performPlayerHit(g.player.weapons[gs.currSelectedItem],
					g.currentEnemies[gs.currSelectedEnemy])
			}
			return
		}

		// select item
		if splitted[0][0] >= 'a' && splitted[0][0] <= 'z' {
			itemnum := 255 - ('a' - splitted[0][0] - 1)
			print(splitted[0][0], " ", itemnum)
			if uint8(len(g.player.weapons)) > itemnum {
				// gs.currLog = g.player.weapons[itemnum].getName()
				gs.currSelectedItem = itemnum
			}
		}

		//select enemy
		if splitted[0][0] >= '0' && splitted[0][0] <= '9' {
			enemynum := splitted[0][0] - '1'
			// print(splitted[0][0], '1', '0', enemynum)
			if uint8(len(g.currentEnemies)) > enemynum {
				// gs.currLog = g.currentEnemies[enemynum].getName()
				gs.currSelectedEnemy = enemynum
			}
		}

		if len(g.currentEnemies) > 0 {
			gs.currLog += g.getPossibleAttackStringDescription(g.player.weapons[gs.currSelectedItem],
				g.currentEnemies[gs.currSelectedEnemy])
		}
	}
}
