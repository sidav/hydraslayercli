package main

import (
	"fmt"
	"strings"
)

type weapon struct {
	weaponType weaponType
	element    uint8

	damage int
}

func (w *weapon) getName() string {
	name := getElementName(w.element)
	if len(name) > 0 {
		name += " "
	}
	switch w.weaponType {
	case WTYPE_SUBSTRACTOR:
		name += fmt.Sprintf("-%d Substractor", w.damage)
	case WTYPE_DIVISOR:
		switch w.damage {
		case 2:
			name += "Bisector"
		case 3:
			name += "Trisector"
		case 10:
			name += "Decimator"
		default:
			name += fmt.Sprintf("/%d Divisor", w.damage)
		}
	default:
		name += "SomeWeapon"
	}
	return getElementColorStr(w.element) + strings.Title(name) + Reset
}
