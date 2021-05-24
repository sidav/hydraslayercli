package main

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
	game := initGame(selectDifficulty())
	game.run()
	cw.closeConsole()
	defer func() {
		if r := recover(); r != nil {
			cw.closeConsole()
			panic(r)
		}
	}()
}
