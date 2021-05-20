package main

type element uint8

const (
	ELEMENT_NONE element = iota
	ELEMENT_FIRE
	ELEMENT_ICE
	ELEMENT_STONE
	ELEMENT_STORM
	ELEMENTS_TOTAL
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

func getRandomElement() element {
	return element(uint8(rnd.Rand(int(ELEMENTS_TOTAL))))
}
