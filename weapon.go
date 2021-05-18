package main

import "fmt"

type weapon struct {
	weaponType weaponType
	element    element

	// for substractors
	damage int
}

func (w *weapon) getName() string {
	switch w.weaponType {
	case WTYPE_SUBSTRACTOR:
		return fmt.Sprintf("-%d Substractor", w.damage)
	default:
		return "SomeWeapon"
	}
}
