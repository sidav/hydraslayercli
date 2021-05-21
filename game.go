package main

import "fmt"

type game struct {
	currentTurn        int
	currentStageNumber int
	enemies            []*enemy
	treasure           []*item
	player             *player
	gameScreen         gameScreen

	// player-related
	abortGame, turnMade, stageFinished  bool
	currLog                             string
	currSelectedItem, currSelectedEnemy int
}

func initGame() *game {
	g := &game{
		currentTurn:        1,
		currentStageNumber: 0,
		enemies:            []*enemy{},
		player: &player{
			hp:    10,
			maxhp: 10,
			items: []*item{
				{
					weaponInfo: &weapon{
						weaponType: WTYPE_SUBSTRACTOR,
						element:    getRandomElement(),
						damage:     1,
					},
				},
				{
					weaponInfo: &weapon{
						weaponType: WTYPE_SUBSTRACTOR,
						element:    getRandomElement(),
						damage:     2,
					},
				},
			},
		},
	}
	g.generateCurrentStage()
	return g
}

func (g *game) run() {
	for !g.abortGame {
		if g.currLog == "" {
			g.currLog = g.getPossibleAttackStringDescription(g.player.items[g.currSelectedItem].weaponInfo,
				g.enemies[g.currSelectedEnemy])
		}
		g.playerTurn()
		if g.turnMade {
			g.actForEnemies()
			g.currentTurn++
			g.turnMade = false
		}
		if g.stageFinished {
			g.currentStageNumber++
			g.currentTurn = 1
			g.generateCurrentStage()
			g.turnMade = false
			g.stageFinished = false
			g.currSelectedEnemy = 0
			g.currLog = fmt.Sprintf("Welcome to stage %d! \n%s", g.currentStageNumber,
				g.getPossibleAttackStringDescription(g.player.items[g.currSelectedItem].weaponInfo,
				g.enemies[g.currSelectedEnemy]))
		}
	}
}

func (g *game) actForEnemies() {
	for i := len(g.enemies) - 1; i >= 0; i-- {
		if g.enemies[i].heads == 0 {
			g.enemies = append(g.enemies[:i], g.enemies[i+1:]...)
		} else {
			damage := g.calculateDamageByHeads(g.enemies[i].heads)
			g.currLog += fmt.Sprintf(" %s bites you for %d damage. ", g.enemies[i].getName(), damage)
			g.player.hp -= damage
		}
	}
}
