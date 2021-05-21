package main

import (
	"fmt"
	"strings"
)

type enemy struct {
	name    string
	heads   int
	element uint8
}

func (e *enemy) getName() string {
	name := fmt.Sprintf("%s (%d heads)", e.name, e.heads)
	if getElementName(e.element) != "" {
		name = getElementName(e.element) + " " + name
	}
	return colorizeString(getElementColorStr(e.element), strings.Title(name))
}
