package main

const (
	ELEMENT_NONE uint8 = iota
	ELEMENT_FIRE
	ELEMENT_ICE
	ELEMENT_STONE
	ELEMENT_STORM
	ELEMENT_MAGMA
	ELEMENT_DISTORTION
	ELEMENTS_TOTAL // for random
)

func getElementName(element uint8) string {
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
	case ELEMENT_MAGMA:
		return "magma"
	case ELEMENT_DISTORTION:
		return "distorted"
	}
	return "MISSING ELEMENT NAME"
}

func getElementColorStr(element uint8) string {
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
	case ELEMENT_MAGMA:
		return Red
	case ELEMENT_DISTORTION:
		return Green
	}
	return "MISSING ELEMENT NAME"
}

func getRandomElement() uint8 {
	return uint8(rnd.Rand(int(ELEMENTS_TOTAL)))
}

var headRegrowsForElement = map[uint8]map[uint8]int{
	// HEADS_ELEM: {WEAPON_ELEM: REGROW}
	// -2 regrow means duplicate remaining heads
	ELEMENT_NONE:  {ELEMENT_NONE: 1, ELEMENT_FIRE: 0, ELEMENT_ICE: 0, ELEMENT_STONE: 0, ELEMENT_STORM: 0},

	ELEMENT_FIRE:  {ELEMENT_NONE: 1, ELEMENT_FIRE: -2, ELEMENT_STONE: 1, ELEMENT_STORM: 1, ELEMENT_MAGMA: 2},
	ELEMENT_ICE:   {ELEMENT_NONE: 1, ELEMENT_ICE: -2, ELEMENT_STONE: 1, ELEMENT_STORM: 1},
	ELEMENT_STONE: {ELEMENT_NONE: 1, ELEMENT_FIRE: 1, ELEMENT_ICE: 1, ELEMENT_STONE: -2, ELEMENT_MAGMA: 2},
	ELEMENT_STORM: {ELEMENT_NONE: 1, ELEMENT_FIRE: 1, ELEMENT_ICE: 1,ELEMENT_STORM: -2, ELEMENT_MAGMA: 1},

	ELEMENT_MAGMA: {ELEMENT_MAGMA: -2, ELEMENT_FIRE: 2, ELEMENT_STONE: 2},
	ELEMENT_DISTORTION: {ELEMENT_DISTORTION: -2, ELEMENT_FIRE: 2, ELEMENT_ICE: 2, ELEMENT_STONE: 2, ELEMENT_STORM: 2},
}
