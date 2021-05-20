package main

type game struct {
	currentTurn    int
	currentEnemies []*enemy
	player         *player

	exit     bool
	turnMade bool
}

func initGame() *game {
	g := &game{
		currentTurn: 0,
		currentEnemies: []*enemy{},
		player: &player{
			hp:    10,
			maxhp: 10,
			weapons: []*weapon{
				{
					weaponType: WTYPE_SUBSTRACTOR,
					element:    getRandomElement(),
					damage:     1,
				},
				{
					weaponType: WTYPE_SUBSTRACTOR,
					element:    getRandomElement(),
					damage:     2,
				},
				{
					weaponType: WTYPE_DIVISOR,
					element:    getRandomElement(),
					damage:     2,
				},
			},
		},
	}
	g.currentEnemies = append(g.currentEnemies, g.generateHydra(1, 1))
	return g
}

func (g *game) run() {
	gs := gameScreen{}
	for !g.exit {
		gs.renderScreen(g)
		if g.turnMade {
			for i := range g.currentEnemies {
				if g.currentEnemies[i].heads == 0 {
					g.currentEnemies = append(g.currentEnemies[:i], g.currentEnemies[i+1:]...)
				} else {
					g.player.hp -= g.calculateDamageByHeads(g.currentEnemies[i].heads)
				}
			}
			g.currentTurn++
			g.turnMade = false
		}
	}
}
