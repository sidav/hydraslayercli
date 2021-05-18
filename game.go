package main

type game struct {
	currentTurn int
	currentEnemies []*enemy
	player *player
}

func initGame() *game {
	return &game{
		currentTurn:    0,
		currentEnemies: []*enemy{
			{
				name:    "hydra",
				heads:   5,
				element: ELEMENT_NONE,
			},
		},
		player: &player{
			hp:      10,
			maxhp:   10,
			weapons: []*weapon{
				{
					weaponType: WTYPE_SUBSTRACTOR,
					element:    ELEMENT_NONE,
					damage:     1,
				},
			},
		},
	}
}

func (g *game) run() {
	gs := gameScreen{}
	gs.renderScreen(g)
}
