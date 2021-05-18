package main

type element uint8

const (
	ELEMENT_NONE element = iota
	ELEMENT_FIRE
	ELEMENT_ICE
)

func getElementName(element element) string {
	switch element {
	case ELEMENT_NONE:
		return ""
	case ELEMENT_FIRE:
		return "blazing"
	case ELEMENT_ICE:
		return "ice"
	}
	return "MISSING ELEMENT NAME"
}
