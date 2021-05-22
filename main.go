package main

var rnd FibRandom
var auxrnd FibRandom

func main() {
	rnd.InitDefault()
	auxrnd.InitDefault()
	initGame().run()
}
