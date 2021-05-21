package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type gameScreen struct {}

func (gs *gameScreen) renderScreen(g *game) {
	fmt.Println("\033[2J") // linux only!
	println("Turn ", g.currentTurn)
	println("")
	println("You see here enemies:")
	for i, e := range g.currentEnemies {
		selectStr := "   "
		if i == g.currSelectedEnemy {
			selectStr = "-> "
		}
		println(fmt.Sprintf("%s%d: %s", selectStr, i+1, e.getName()))
	}
	println(fmt.Sprintf("You have %d/%d hp and:", g.player.hp, g.player.maxhp))
	for i, w := range g.player.items {
		selectStr := "  "
		if i == g.currSelectedItem {
			selectStr = "->"
		}
		println(fmt.Sprintf("%s %c: %s", selectStr, 'A'+i, w.getName()))
	}
	if len(g.currentEnemies) > 0 {
		for len(g.currentEnemies) <= g.currSelectedEnemy {
			g.currSelectedEnemy--
		}
	}
	println(g.currLog)
	print("Your action?\n> ")
}

func (gs *gameScreen) readPlayerInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.Trim(input, " \n")
	return input
}
