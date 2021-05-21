package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type gameScreen struct{}

func (gs *gameScreen) renderScreen(g *game) {
	fmt.Println("\033[2J") // linux only!
	println(fmt.Sprintf("Stage %d: Turn %d", g.currentStageNumber, g.currentTurn))
	println("")

	if len(g.enemies) > 0 {
		println("You see here enemies:")
		for i, e := range g.enemies {
			selectStr := "   "
			if i == g.currSelectedEnemy {
				selectStr = "-> "
			}
			println(fmt.Sprintf("%s%d: %s", selectStr, i+1, e.getName()))
		}
	} else if len(g.treasure) > 0 {
		println("You see here treasure:")
		for i, t := range g.treasure {
			println(fmt.Sprintf("%d: %s", i+1, t.getName()))
		}
	} else {
		println("There is nothing to do here. Use \"move\" command to move to the next stage!")
	}
	println(fmt.Sprintf("You have %d/%d hp and %d/%d items:",
		g.player.hp, g.player.maxhp, len(g.player.items), g.player.maxItems))
	for i, w := range g.player.items {
		selectStr := "  "
		if i == g.currSelectedItem {
			selectStr = "->"
		}
		println(fmt.Sprintf("%s %c: %s", selectStr, 'A'+i, w.getName()))
	}
	if len(g.enemies) > 0 {
		for len(g.enemies) <= g.currSelectedEnemy {
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
