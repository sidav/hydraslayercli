package main

const (
	STATUS_CONFUSED uint8 = iota
)

type statusEffect struct {
	statusType     uint8
	turnsRemaining int
}

func (se *statusEffect) getName() string {
	switch se.statusType {
	case STATUS_CONFUSED: return "confused"
	}
	panic("No status effect name!")
}

func (e *enemy) applyStatusEffects(g *game) {
	for i := len(e.statuses)-1; i >= 0; i-- {
		switch e.statuses[i].statusType {
		case STATUS_CONFUSED:
			g.appendToLogMessage("Confused %s %s.", e.getName(), e.getConfusedActionDescription())
		}

		e.statuses[i].turnsRemaining--
		if e.statuses[i].turnsRemaining == 0 {
			g.appendToLogMessage("%s seems no more %s.", e.getName(), e.statuses[i].getName())
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

