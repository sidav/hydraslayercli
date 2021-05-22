package main

type stage struct {
	enemies  []stageEnemyData
	treasure int
}

type stageEnemyData struct {
	minHeads, maxHeads int
}

var StageInfo = []stage{
	0: {
		enemies: []stageEnemyData{
			{minHeads: 1, maxHeads: 3},
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
			{minHeads: 3, maxHeads: 5},
		},
		treasure: 3,
	},
	3: {
		enemies: []stageEnemyData{
			{minHeads: 1, maxHeads: 3},
			{minHeads: 3, maxHeads: 5},

		},
		treasure: 4,
	},
	4: {
		enemies: []stageEnemyData{
			{minHeads: 1, maxHeads: 4},
			{minHeads: 2, maxHeads: 4},
			{minHeads: 3, maxHeads: 4},
		},
		treasure: 5,
	},
	5: {
		enemies: []stageEnemyData{
			{minHeads: 7, maxHeads: 10},
		},
	},
}
