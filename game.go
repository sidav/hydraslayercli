package main

import "fmt"

type game struct {
	currentStageNumber int
	dungeon            *dungeon

	currentTurn int
	enemies     []*enemy
	treasure    []*item
	player      *player

	// conditions
	allEnemiesSkipTurn bool

	// player-related
	coordsX, coordsY                    int
	turnMade, stageFinished             bool
	currLog                             string
	currSelectedItem, currSelectedEnemy int

	// settings-related
	showShortCombatDescription bool
}

func (g *game) getCurrentStage() *stage {
	return g.dungeon.getStageNumber(g.currentStageNumber)
}

func initGame(difficulty string) *game {
	g := &game{
		currentTurn:                1,
		currentStageNumber:         0,
		showShortCombatDescription: true,
		enemies:                    []*enemy{},
		player: &player{
			hp:       10,
			maxhp:    10,
			maxItems: 5,
			items: []*item{
				{
					element: getRandomElement(false, false, true),
					weaponInfo: &weapon{
						weaponType: WTYPE_SUBSTRACTOR,
						damage:     2,
					},
				},
				{
					element: getRandomElement(false, false, true),
					weaponInfo: &weapon{
						weaponType: WTYPE_SUBSTRACTOR,
						damage:     1,
					},
				},
			},
		},
	}

	g.dungeon = dungeons[difficulty]
	g.coordsX, g.coordsY = g.dungeon.generate()
	return g
}

func (g *game) run() {
	screen.init()
	g.player.addItem(g.generateTreasure(0))
	g.player.addItem(g.generateTreasure(0))
	g.generateCurrentStage()
	for !abortGame {
		if g.turnMade || g.currentTurn == 0 {
			g.turnMade = false
			for _, i := range g.player.items {
				i.applyPassiveEffect(g)
			}
		}
		if g.currentTurn == 0 {
			g.currentTurn = 1
		}
		if g.currSelectedItem >= len(g.player.items) {
			g.currSelectedItem = 0
		}
		if g.currLog == "" && len(g.enemies) > 0 {
			g.currLog = g.getPossibleAttackStringDescription(g.player.items[g.currSelectedItem],
				g.enemies[g.currSelectedEnemy])
		}
		g.playerTurn()
		if g.turnMade {
			g.actForEnemies()
			g.currentTurn++
			if g.player.hp <= g.player.maxhp/3 {
				g.appendToLogMessage(colorizeString(Red, "\nBe careful: you're almost dead!"))
			}
		}
		if g.player.hp <= 0 {
			g.appendToLogMessage(colorizeString(Red, "\nYou died... Press ENTER to exit.\n"))
			screen.renderScreen(g)
			screen.readPlayerInput()
			return
		}
		if g.stageFinished {
			if g.currentStageNumber == g.dungeon.getTotalStages()-1 {
				g.appendToLogMessage(colorizeString(Yellow, "\nYou won! Press ENTER to exit.\n"))
				screen.renderScreen(g)
				screen.readPlayerInput()
				return
			}
			if (g.currentStageNumber+1)%3 == 0 {
				g.selectReward()
			}
			g.currentStageNumber++
			g.currentTurn = 0
			g.generateCurrentStage()
			g.stageFinished = false
			g.currSelectedEnemy = 0
			g.setLogMessage("Welcome to stage %d! \n%s", g.currentStageNumber,
				g.getPossibleAttackStringDescription(g.player.items[g.currSelectedItem],
					g.enemies[g.currSelectedEnemy]))
		}
	}
}

func (g *game) getTotalExpectedEnemyDamage() int {
	totaldmg := 0
	for _, e := range g.enemies {
		if e.skipsThisTurn {
			continue
		}
		totaldmg += g.calculateDamageByHeads(e.heads)
	}
	return totaldmg
}

func (g *game) selectReward() {
	choice := showSelectScreen(fmt.Sprintf("Congratulations, you've cleared stage %d!\n"+
		"Select your reward:", g.currentStageNumber+1),
		[]string{"Max health + 2", "Strength +1", "Acquire an item"})
	switch choice {
	case 0:
		g.player.maxhp += 2
	case 1:
		g.player.maxItems += 1
	case 2:
		var itemsToGive []*item
		var namesSlice []string
		for i := 0; i < 5; i++ {
			itemsToGive = append(itemsToGive, g.generateTreasure(g.currentStageNumber))
			namesSlice = append(namesSlice, itemsToGive[i].getName())
		}
		itemChoice := showSelectScreen("Select an item to acquire:", namesSlice)
		g.player.addItem(itemsToGive[itemChoice])
	}
}

func (g *game) setLogMessage(msg string, args ...interface{}) {
	g.currLog = fmt.Sprintf(msg, args...)
}

func (g *game) appendToLogMessage(msg string, args ...interface{}) {
	g.currLog += fmt.Sprintf(" "+msg, args...)
}
