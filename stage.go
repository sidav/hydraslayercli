package main

type stage struct {
	enemies  []stageEnemyData
	treasure int
}

type stageEnemyData struct {
	minHeads, maxHeads                       int
	allowComplexElement, forceComplexElement bool
	allowSpecialElement, forceSpecialElement bool
}

var StageInfo = []stage{
	0: {
		enemies: []stageEnemyData{
			{minHeads: 3, maxHeads: 3},
		},
		treasure: 3,
	},
	1: {
		enemies: []stageEnemyData{
			{minHeads: 1, maxHeads: 3},
			{minHeads: 1, maxHeads: 3},
		},
		treasure: 3,
	},
	2: {
		enemies: []stageEnemyData{
			{minHeads: 3, maxHeads: 5, allowComplexElement: true},
		},
		treasure: 3,
	},
	3: {
		enemies: []stageEnemyData{
			{minHeads: 1, maxHeads: 3, allowComplexElement: true},
			{minHeads: 3, maxHeads: 5, allowComplexElement: true},
		},
		treasure: 4,
	},
	4: {
		enemies: []stageEnemyData{
			{minHeads: 1, maxHeads: 4, allowComplexElement: true},
			{minHeads: 2, maxHeads: 4, allowComplexElement: true},
			{minHeads: 3, maxHeads: 4, forceComplexElement: true},
		},
		treasure: 5,
	},
	5: {
		enemies: []stageEnemyData{
			{minHeads: 7, maxHeads: 10, forceComplexElement: true},
		},
		treasure: 5,
	},
	6: {
		enemies: []stageEnemyData{
			{minHeads: 2, maxHeads: 5, allowComplexElement: true},
			{minHeads: 3, maxHeads: 6, forceComplexElement: true},
			{minHeads: 4, maxHeads: 7, forceComplexElement: true},
		},
		treasure: 5,
	},
	7: {
		enemies: []stageEnemyData{
			{minHeads: 8, maxHeads: 20, forceSpecialElement: true},
		},
	},
}
