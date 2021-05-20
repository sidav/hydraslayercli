package main

type element uint8

const (
	ELEMENT_NONE element = iota
	ELEMENT_FIRE
	ELEMENT_ICE
	ELEMENT_STONE
	ELEMENT_STORM
	ELEMENTS_TOTAL // for random
)

func getElementName(element element) string {
	switch element {
	case ELEMENT_NONE:
		return ""
	case ELEMENT_FIRE:
		return "blazing"
	case ELEMENT_ICE:
		return "ice"
	case ELEMENT_STONE:
		return "stone"
	case ELEMENT_STORM:
		return "storm"
	}
	return "MISSING ELEMENT NAME"
}

func getElementColorStr(element element) string {
	switch element {
	case ELEMENT_NONE:
		return White
	case ELEMENT_FIRE:
		return Red
	case ELEMENT_ICE:
		return Blue
	case ELEMENT_STONE:
		return Gray
	case ELEMENT_STORM:
		return Cyan
	}
	return "MISSING ELEMENT NAME"
}

func getRandomElement() element {
	return element(uint8(rnd.Rand(int(ELEMENTS_TOTAL))))
}

var headRegrowsForElement = map[element]map[element]int{
	// HEADS_ELEM: {WEAPON_ELEM: REGROW}
	// -2 regrow means duplicate remaining heads
	ELEMENT_NONE: {ELEMENT_NONE: 1, ELEMENT_FIRE: 1, ELEMENT_ICE: 1, ELEMENT_STONE: 1, ELEMENT_STORM: 1},
	ELEMENT_FIRE: {ELEMENT_NONE: 1, ELEMENT_FIRE: -2, ELEMENT_ICE: 0, ELEMENT_STONE: 1, ELEMENT_STORM: 1},
	ELEMENT_ICE: {ELEMENT_NONE: 1, ELEMENT_FIRE: 0, ELEMENT_ICE: -2, ELEMENT_STONE: 1, ELEMENT_STORM: 1},
	ELEMENT_STONE: {ELEMENT_NONE: 1, ELEMENT_FIRE: 1, ELEMENT_ICE: 1, ELEMENT_STONE: -2, ELEMENT_STORM: 0},
	ELEMENT_STORM: {ELEMENT_NONE: 1, ELEMENT_FIRE: 1, ELEMENT_ICE: 1, ELEMENT_STONE: 0, ELEMENT_STORM: -2},
}
