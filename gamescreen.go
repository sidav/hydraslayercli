package main

import "fmt"

type gameScreen struct {}

func (gs *gameScreen) renderScreen(g *game) {
	println("Turn ", g.currentTurn)
	println("")
	println("You see here enemies:")
	for _, e := range g.currentEnemies {
		println(" ", e.getName())
	}
	println(fmt.Sprintf("You have %d/%d hp and:", g.player.hp, g.player.maxhp))
	for i, w := range g.player.weapons {
		println(fmt.Sprintf(" %d: %s", i, w.getName()))
	}
}
