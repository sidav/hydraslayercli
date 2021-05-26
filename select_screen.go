package main

import (
	"fmt"
	"strings"
)

func showSelectScreen(title string, valsArr []string) int {
	for {
		cw.clear()
		cw.println(title)
		cw.println("")
		for i, val := range valsArr {
			cw.println(fmt.Sprintf(" %d: %s", i+1, strings.Title(val)))
		}
		cw.print("> ")
		cw.flush()
		input := cw.read()
		inputIndex, cType := charToIndexWithType(input[0])
		if cType == INDEX_NUMBER && inputIndex < len(valsArr) {
			return inputIndex
		}
	}
}
