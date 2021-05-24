package main

import "os"

var rnd FibRandom
var auxrnd FibRandom
var abortGame bool
var screen gameScreen
var cw consoleWrapper

func main() {
	cw = &cwtcell{}
	cw.init()
	rnd.InitDefault()
	auxrnd.InitDefault()

	var initialDiff string
	if len(os.Args) > 1 {
		initialDiff = os.Args[1]
	} else {
		initialDiff = selectDifficulty()
	}

	game := initGame(initialDiff)
	game.run()
	cw.closeConsole()
	defer func() {
		if r := recover(); r != nil {
			cw.closeConsole()
			panic(r)
		}
	}()
}
