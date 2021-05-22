package main

var rnd FibRandom
var auxrnd FibRandom
var abortGame bool

func main() {
	rnd.InitDefault()
	auxrnd.InitDefault()
	game := initGame()
	game.run()
	game.gameScreen.cw.closeConsole()
	defer func() {
		if r := recover(); r != nil {
			game.gameScreen.cw.closeConsole()
			panic(r)
		}
	}()
}
