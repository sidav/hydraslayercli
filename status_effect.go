package main

const (
	STATUS_CONFUSED uint8 = iota
)

type statusEffect struct {
	statusType     uint8
	turnsRemaining int
}

func (e *enemy) applyStatusEffects() {
	for i := len(e.statuses)-1; i >= 0; i-- {
		e.statuses[i].turnsRemaining--
		if e.statuses[i].turnsRemaining == 0 {
			e.statuses = append(e.statuses[:i], e.statuses[i+1:]...)
		}
	}
}

func (e *enemy) hasStatusEffectOfType(t uint8) bool {
	for i := len(e.statuses)-1; i >= 0; i-- {
		if e.statuses[i].statusType == t {
			return true
		}
	}
	return false
}

