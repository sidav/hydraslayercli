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
}

func(gs *gameScreen) init() {
}

func (gs *gameScreen) renderScreen(g *game) {
	cw.clear()
	stageString := fmt.Sprintf("Stage %d/%d", g.currentStageNumber+1, g.dungeon.getTotalStages())
	if g.currentStageNumber == g.dungeon.getTotalStages() - 1 {
		stageString = fmt.Sprintf("FINAL STAGE")
	}
	cw.println(fmt.Sprintf("%s: Turn %d (%s)", stageString, g.currentTurn, g.dungeon.name))
	cw.println(fmt.Sprintf("  %s", g.getCurrentStage().name))

	if len(g.enemies) > 0 {
		cw.println("You see here enemies:")
		for i, e := range g.enemies {
			selectStr := "   "
			attDescrStr := ""
			attDescrStr = g.getShortPossibleAttackStringDescription(
				g.player.items[g.currSelectedItem],
				e)
			if i == g.currSelectedEnemy {
				selectStr = "-> "
			}
			cw.println(fmt.Sprintf("%s%d: %s%s", selectStr, i+1, e.getNameWithStatus(), attDescrStr))
		}
	} else if len(g.treasure) > 0 {
		cw.println("You see here treasure:")
		for i, t := range g.treasure {
			cw.println(fmt.Sprintf("   %d: %s", i+1, t.getName()))
		}
	} else {
		cw.println("There is nothing to do here. Use \"move\" command to move to the next stage!")
	}

	hpColor := Green
	if g.player.hp < g.player.maxhp/3 {
		hpColor = Red
	}
	cw.println(fmt.Sprintf("You have %s hp and %d/%d items:",
		colorizeString(hpColor, fmt.Sprintf("%d/%d", g.player.hp, g.player.maxhp)),
		len(g.player.items), g.player.maxItems))

	for i, w := range g.player.items {
		selectStr := "  "
		if i == g.currSelectedItem {
			selectStr = "->"
		}
		cw.println(fmt.Sprintf("%s %c: %s", selectStr, 'A'+i, w.getName()))
	}
	if len(g.enemies) > 0 {
		for len(g.enemies) <= g.currSelectedEnemy {
			g.currSelectedEnemy--
		}
	}
	cw.println("")
	cw.println(g.currLog)
	expectedDamage := g.getTotalExpectedEnemyDamage()
	expectedDamageStr := ""
	if expectedDamage > 0 {
		expectedDamageStr = fmt.Sprintf("%d damage expected (%d hp). ", expectedDamage, g.player.hp - expectedDamage)
	}
	cw.print(fmt.Sprintf("%sYour action?\n> ", expectedDamageStr))
	cw.flush()
}

func (gs *gameScreen) readPlayerInput() string {
	return cw.read()
}
