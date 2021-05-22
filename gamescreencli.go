package main

import (
	"fmt"
)

type consoleWrapper interface {
	init()
	print(string)
	println(string)
	clear()
	flush()
	read() string
	closeConsole()
}

type gameScreen struct{
	cw consoleWrapper
}

func(gs *gameScreen) init() {
	gs.cw = &consoleWrapperStdout{}
	gs.cw = &cwtcell{}
	gs.cw.init()
}

func (gs *gameScreen) renderScreen(g *game) {
	gs.cw.clear()
	stageString := fmt.Sprintf("Stage %d/%d", g.currentStageNumber+1, len(StageInfo))
	if g.currentStageNumber == len(StageInfo) - 1 {
		stageString = fmt.Sprintf("FINAL STAGE", g.currentStageNumber)
	}
	gs.cw.println(fmt.Sprintf("%s: Turn %d", stageString, g.currentTurn))
	gs.cw.println("")

	if len(g.enemies) > 0 {
		gs.cw.println("You see here enemies:")
		for i, e := range g.enemies {
			selectStr := "   "
			attDescrStr := ""
			attDescrStr = g.getShortPossibleAttackStringDescription(
				g.player.items[g.currSelectedItem],
				e)
			if i == g.currSelectedEnemy {
				selectStr = "-> "
			}
			gs.cw.println(fmt.Sprintf("%s%d: %s%s", selectStr, i+1, e.getNameWithStatus(), attDescrStr))
		}
	} else if len(g.treasure) > 0 {
		gs.cw.println("You see here treasure:")
		for i, t := range g.treasure {
			gs.cw.println(fmt.Sprintf("%d: %s", i+1, t.getName()))
		}
	} else {
		gs.cw.println("There is nothing to do here. Use \"move\" command to move to the next stage!")
	}

	hpColor := Green
	if g.player.hp < g.player.maxhp/3 {
		hpColor = Red
	}
	gs.cw.println(fmt.Sprintf("You have %s hp and %d/%d items:",
		colorizeString(hpColor, fmt.Sprintf("%d/%d", g.player.hp, g.player.maxhp)),
		len(g.player.items), g.player.maxItems))

	for i, w := range g.player.items {
		selectStr := "  "
		if i == g.currSelectedItem {
			selectStr = "->"
		}
		gs.cw.println(fmt.Sprintf("%s %c: %s", selectStr, 'A'+i, w.getName()))
	}
	if len(g.enemies) > 0 {
		for len(g.enemies) <= g.currSelectedEnemy {
			g.currSelectedEnemy--
		}
	}
	gs.cw.println("")
	gs.cw.println(g.currLog)
	gs.cw.print("Your action?\n> ")
	gs.cw.flush()
}

func (gs *gameScreen) readPlayerInput() string {
	return gs.cw.read()
}
