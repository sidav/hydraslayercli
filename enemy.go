package main

import (
	"fmt"
	"strings"
)

type enemy struct {
	name          string
	heads         int
	element       *element
	statuses      []*statusEffect
	skipsThisTurn bool
}

func (e *enemy) getName() string {
	name := fmt.Sprintf("%d-headed %s", e.heads, e.name)
	if e.element.name != "" {
		name = e.element.name + " " + name
	}
	return colorizeString(e.element.colorString, strings.Title(name))
}

func (e *enemy) getNameWithStatus() string {
	statusLine := ""
	for i, se := range e.statuses {
		if i > 0 {
			statusLine += ", "
		}
		statusLine += se.getName()
	}
	if statusLine != "" {
		statusLine = " (" + statusLine + ")"
	}
	return e.getName() + statusLine
}

func (e *enemy) getInfo() string {
	return fmt.Sprintf("%s. %s", e.getName(), e.element.description)
}

var confusedThings = []string{
	"sniffs ground",
	"talks gibberish",
	"meows",
	"purrs",
	"thinks of chicken salad",
	"yowls",
	"sneezes",
	"looks at you in awe",
	"imagines attacking you for 8 damage",
	"thinks that it wants to have more heads",
	"twerks",
	"thinks of laying eggs",
}

func (e *enemy) getConfusedActionDescription() string {
	return fmt.Sprintf(confusedThings[auxrnd.Rand(len(confusedThings))]) // , randomHead1, randomHead2)
}

func (g *game) actForEnemies() {
	for i := len(g.enemies) - 1; i >= 0; i-- {
		if g.enemies[i].heads == 0 {
			g.enemies = append(g.enemies[:i], g.enemies[i+1:]...)
		} else {
			if g.enemies[i].element.elementCode == ELEMENT_REGROW_AURA {
				for _, e := range g.enemies {
					healed := false
					if e.heads > 0 && e.element.elementCode != ELEMENT_REGROW_AURA {
						e.heads++
						g.enemies[i].skipsThisTurn = true
						healed = true
					}
					if healed {
						g.appendToLogMessage("%s imbues other hydras with healing power!", g.enemies[i].getName())
					}
				}
			}
			g.enemies[i].applyStatusEffects(g)
			if g.enemies[i].skipsThisTurn {
				g.enemies[i].skipsThisTurn = false
			} else {
				if g.allEnemiesSkipTurn {
					continue
				}
				damage := g.calculateDamageByHeads(g.enemies[i].heads)
				g.appendToLogMessage("%s bites you for %d damage. ", g.enemies[i].getName(), damage)
				g.player.hp -= damage
				if g.enemies[i].element.elementCode == ELEMENT_VAMPIRIC {
					g.appendToLogMessage("%s grows itself %d heads from your blood!! ", g.enemies[i].getName(), damage)
					g.enemies[i].heads += damage
				}
			}
		}
	}
	if g.allEnemiesSkipTurn {
		g.allEnemiesSkipTurn = false
	}
}
