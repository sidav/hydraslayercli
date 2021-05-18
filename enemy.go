package main

import (
	"fmt"
	"strings"
)

type enemy struct {
	name string
	heads int
	element element
}

func (e *enemy) getName() string {
	return strings.Title(fmt.Sprintf("%s %s (%d heads)", getElementName(e.element), e.name, e.heads))
}
