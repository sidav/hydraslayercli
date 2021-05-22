package main

var rnd FibRandom
var auxrnd FibRandom
var abortGame bool
var screen gameScreen

func main() {
	rnd.InitDefault()
	auxrnd.InitDefault()
	game := initGame()
	game.run()
	screen.cw.closeConsole()
	defer func() {
		if r := recover(); r != nil {
			screen.cw.closeConsole()
			panic(r)
		}
	}()
}
