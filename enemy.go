package main

import (
	"fmt"
	"strings"
)

type enemy struct {
	name     string
	heads    int
	element  uint8
	statuses []*statusEffect
}

func (e *enemy) getName() string {
	name := fmt.Sprintf("%d-headed %s",e.heads,  e.name)
	if getElementName(e.element) != "" {
		name = getElementName(e.element) + " " + name
	}
	return colorizeString(getElementColorStr(e.element), strings.Title(name))
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

var confusedThings = []string{
	"sniffs ground",
	"talks gibberish",
	"meows",
	"purrs",
	"thinks of chicken salad",
	"yowls",
	"sneezes",
	"looks at you in awe",
}

func (e *enemy) getConfusedActionDescription() string {
	//randomHead1 := auxrnd.Rand(e.heads)
	//randomHead2 := -1
	//if e.heads > 1 {
	//	randomHead2 = rnd.Rand(e.heads)
	//}
	return fmt.Sprintf(confusedThings[auxrnd.Rand(len(confusedThings))]) // , randomHead1, randomHead2)
}
