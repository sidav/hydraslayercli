package main

const (
	ELEMENT_NONE uint8 = iota
	ELEMENT_FIRE
	ELEMENT_ICE
	ELEMENT_STONE
	ELEMENT_STORM
	ELEMENT_MAGMA
	ELEMENT_ENERGY

	ELEMENT_REGROW_AURA

	ELEMENT_VAMPIRIC
	ELEMENT_GROWING
	ELEMENTS_TOTAL // for random
)

type element struct {
	elementCode                  uint8
	name                         string
	defaultRegrowForMissingValue int
	isNonBasic                   bool
	isSupporting                 bool // useful only when the hydra is not alone
	isBoss                       bool
	isForEnemiesOnly             bool
	colorString                  string
	description                  string
}

var elementsData = []*element{
	{
		elementCode: ELEMENT_NONE,
		name:        "",
		colorString: White,
		isNonBasic:  false,
		description: "It regenerates no heads from elemental damage. ",
	},
	{
		elementCode: ELEMENT_FIRE,
		name:        "Blazing",
		colorString: Red,
		isNonBasic:  false,
		description: "It fears ice.",
	},
	{
		elementCode: ELEMENT_ICE,
		name:        "Frozen",
		colorString: Blue,
		isNonBasic:  false,
		description: "It fears fire.",
	},
	{
		elementCode: ELEMENT_STONE,
		name:        "Stone",
		colorString: Gray,
		description: "It fears storms.",
	},
	{
		elementCode: ELEMENT_STORM,
		name:        "Storm",
		colorString: Yellow,
		description: "It fears stone.",
	},
	{
		elementCode: ELEMENT_MAGMA,
		name:        "Magmatic",
		colorString: Red,
		isNonBasic:  true,
		description: "It fears no fire and stone. ",
	},
	{
		elementCode: ELEMENT_ENERGY,
		name:        "Energy",
		colorString: Cyan,
		isNonBasic:  true,
		description: "It regenerates no heads after non-elemental damage. ",
	},

	{
		elementCode:      ELEMENT_REGROW_AURA,
		name:             "Healer",
		colorString:      Green,
		isNonBasic:       true,
		isSupporting:     true,
		isForEnemiesOnly: true,
		description:      "It allows other hydras to grow additional head each turn. ",
	},

	{
		elementCode: ELEMENT_VAMPIRIC,
		name:        "Vampiric",
		colorString: Red,
		isBoss:      true,
		description: "It regenerates no heads after damage, but grows heads after damaging someone.",
	},
	{
		elementCode:                  ELEMENT_GROWING,
		name:                         "Fast-healing",
		colorString:                  Green,
		isBoss:                       true,
		isForEnemiesOnly:             true,
		description:                  "It regenerates 3 heads each time.",
		defaultRegrowForMissingValue: 3,
	},
}

func (e *element) getElementColorStr(element uint8) string {
	if e.colorString != "" {
		return e.colorString
	} else {
		return "MISSING ELEMENT COLOR"
	}
}

func getRandomElement(allowNonBasic, allowSpecial, isForItem bool) *element {
	var element *element
	conditionsSatisfied := false
	for !conditionsSatisfied {
		element = elementsData[rnd.Rand(len(elementsData))]
		conditionsSatisfied = (allowNonBasic || !element.isNonBasic) && (allowSpecial || !element.isBoss) &&
			(!isForItem || !element.isForEnemiesOnly)
	}
	return element
}

func getRandomNonBasicElement() *element {
	var element *element
	conditionsSatisfied := false
	counter := 0
	for !conditionsSatisfied {
		element = elementsData[rnd.Rand(len(elementsData))]
		conditionsSatisfied = element.isNonBasic
		counter++
		if counter == 100 {
			panic("randomSpecialElement")
		}
	}
	return element
}

func getRandomSpecialElement() *element {
	var element *element
	conditionsSatisfied := false
	counter := 0
	for !conditionsSatisfied {
		element = elementsData[rnd.Rand(len(elementsData))]
		conditionsSatisfied = element.isBoss
		counter++
		if counter == 100 {
			panic("randomSpecialElement")
		}
	}
	return element
}

const (
	REGROW_SIMPLE    = ""
	REGROW_DUPLICATE = "duplicate"
)

func getHeadRegrowForElement(headsElement, weaponElement *element) int {
	regrow, found := headRegrowsForElementsTable[headsElement.elementCode][weaponElement.elementCode]
	if !found {
		// print("ELEMENT NOT FOUND IN TABLE")
		return headsElement.defaultRegrowForMissingValue
	}
	return regrow
}

var headRegrowsForElementsTable = map[uint8]map[uint8]int{
	// HEADS_ELEM: {WEAPON_ELEM: REGROW}
	// -2 regrow means duplicate remaining heads
	ELEMENT_NONE: {ELEMENT_NONE: 1, ELEMENT_FIRE: 0, ELEMENT_ICE: 0, ELEMENT_STONE: 0, ELEMENT_STORM: 0},

	ELEMENT_FIRE:  {ELEMENT_NONE: 1, ELEMENT_FIRE: -2, ELEMENT_STONE: 1, ELEMENT_STORM: 1, ELEMENT_MAGMA: 2},
	ELEMENT_ICE:   {ELEMENT_NONE: 1, ELEMENT_ICE: -2, ELEMENT_STONE: 1, ELEMENT_STORM: 1},
	ELEMENT_STONE: {ELEMENT_NONE: 1, ELEMENT_FIRE: 1, ELEMENT_ICE: 1, ELEMENT_STONE: -2, ELEMENT_MAGMA: 2},
	ELEMENT_STORM: {ELEMENT_NONE: 1, ELEMENT_FIRE: 1, ELEMENT_ICE: 1, ELEMENT_STORM: -2, ELEMENT_MAGMA: 1},

	ELEMENT_MAGMA:  {ELEMENT_MAGMA: -2, ELEMENT_FIRE: 2, ELEMENT_STONE: 2},
	ELEMENT_ENERGY: {ELEMENT_ENERGY: -2, ELEMENT_FIRE: 2, ELEMENT_ICE: 2, ELEMENT_STONE: 2, ELEMENT_STORM: 2},

	ELEMENT_REGROW_AURA: {},

	ELEMENT_VAMPIRIC: {},
	ELEMENT_GROWING:  {},
}
