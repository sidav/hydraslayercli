package main

func (g *game) generateHydra(depth, difficulty int) *enemy {
	minHeads := depth+1
	maxHeads := minHeads + depth/2 + 2
	return &enemy{
		name:    "hydra",
		heads:   rnd.RandInRange(minHeads, maxHeads),
		element: getRandomElement(),
	}
}
