package main

import "fmt"

type game struct {
	currentTurn    int
	currentEnemies []*enemy
	player         *player
	gameScreen     gameScreen

	// player-related
	exit                                bool
	turnMade                            bool
	currLog                             string
	currSelectedItem, currSelectedEnemy int
}

func initGame() *game {
	g := &game{
		currentTurn:    0,
		currentEnemies: []*enemy{},
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
				{
					weaponInfo: &weapon{
						weaponType: WTYPE_DIVISOR,
						element:    getRandomElement(),
						damage:     2,
					},
				},
				{
					itemType: ITEM_HEAL,
				},
				{
					itemType: ITEM_ENCHANTER,
				},
				{
					itemType: ITEM_DESTROY_HYDRA,
				},
			},
		},
	}
	g.currentEnemies = append(g.currentEnemies, g.generateHydra(1, 1))
	g.currentEnemies = append(g.currentEnemies, g.generateHydra(1, 1))
	return g
}

func (g *game) run() {
	for !g.exit {
		if g.currLog == "" {
			g.currLog = g.getPossibleAttackStringDescription(g.player.items[g.currSelectedItem].weaponInfo,
				g.currentEnemies[g.currSelectedEnemy])
		}
		g.playerTurn()
		if g.turnMade {
			g.actForEnemies()
			g.currentTurn++
			g.turnMade = false
		}
	}
}

func (g *game) actForEnemies() {
	for i := len(g.currentEnemies) - 1; i >= 0; i-- {
		if g.currentEnemies[i].heads == 0 {
			g.currentEnemies = append(g.currentEnemies[:i], g.currentEnemies[i+1:]...)
		} else {
			damage := g.calculateDamageByHeads(g.currentEnemies[i].heads)
			g.currLog += fmt.Sprintf(" %s bites you for %d damage. ", g.currentEnemies[i].getName(), damage)
			g.player.hp -= damage
		}
	}
}
