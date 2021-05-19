package main

type game struct {
	currentTurn    int
	currentEnemies []*enemy
	player         *player

	exit     bool
	turnMade bool
}

func initGame() *game {
	return &game{
		currentTurn: 0,
		currentEnemies: []*enemy{
			{
				name:    "hydra",
				heads:   5,
				element: ELEMENT_NONE,
			},
		},
		player: &player{
			hp:    10,
			maxhp: 10,
			weapons: []*weapon{
				{
					weaponType: WTYPE_SUBSTRACTOR,
					element:    ELEMENT_NONE,
					damage:     1,
				},
				{
					weaponType: WTYPE_SUBSTRACTOR,
					element:    ELEMENT_NONE,
					damage:     2,
				},
			},
		},
	}
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
		}
	}
}
