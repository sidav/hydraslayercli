package main

import "fmt"

type game struct {
	currentTurn        int
	currentStageNumber int
	enemies            []*enemy
	treasure           []*item
	player             *player

	// conditions
	enemiesSkipTurn bool

	// player-related
	turnMade, stageFinished  bool
	currLog                             string
	currSelectedItem, currSelectedEnemy int

	// settings-related
	showShortCombatDescription bool
}

func initGame() *game {
	g := &game{
		currentTurn:        1,
		currentStageNumber: 0,
		enemies:            []*enemy{},
		player: &player{
			hp:       10,
			maxhp:    10,
			maxItems: 5,
			items: []*item{
				{
					element: getRandomElement(),
					weaponInfo: &weapon{
						weaponType: WTYPE_SUBSTRACTOR,
						damage:     2,
					},
				},
				{
					element: getRandomElement(),
					weaponInfo: &weapon{
						weaponType: WTYPE_SUBSTRACTOR,
						damage:     1,
					},
				},
			},
		},
	}
	screen.init()
	g.player.addItem(g.generateTreasure(0))
	g.player.addItem(g.generateTreasure(0))
	g.generateCurrentStage()
	return g
}

func (g *game) run() {
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
			if g.player.hp < g.player.maxhp/3 {
				g.appendToLogMessage(colorizeString(Red, "Be careful: you're almost dead!"))
			}
		}
		if g.player.hp <= 0 {
			g.appendToLogMessage(" You died... Press ENTER to exit.\n")
			screen.renderScreen(g)
			screen.readPlayerInput()
			return
		}
		if g.stageFinished {
			g.currentStageNumber++
			if g.currentStageNumber == len(StageInfo) {
				g.appendToLogMessage(" You won! Press ENTER to exit.\n")
				screen.renderScreen(g)
				screen.readPlayerInput()
				return
			}
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

func (g *game) actForEnemies() {
	for i := len(g.enemies) - 1; i >= 0; i-- {
		if g.enemies[i].heads == 0 {
			g.enemies = append(g.enemies[:i], g.enemies[i+1:]...)
		} else {
			if g.enemies[i].hasStatusEffectOfType(STATUS_CONFUSED) {
			} else {
				if g.enemiesSkipTurn {
					continue
				}
				damage := g.calculateDamageByHeads(g.enemies[i].heads)
				g.appendToLogMessage("%s bites you for %d damage. ", g.enemies[i].getName(), damage)
				g.player.hp -= damage
			}
			g.enemies[i].applyStatusEffects(g)
		}
	}
	if g.enemiesSkipTurn {
		g.enemiesSkipTurn = false
	}
}

func (g *game) setLogMessage(msg string, args ...interface{}) {
	g.currLog = fmt.Sprintf(msg, args...)
}

func (g *game) appendToLogMessage(msg string, args ...interface{}) {
	g.currLog += fmt.Sprintf(" "+msg, args...)
}
