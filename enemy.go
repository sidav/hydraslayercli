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
