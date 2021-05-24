package main

type stage struct {
	enemiesVariants [][]*stageEnemyData
	treasure        int
}

type stageEnemyData struct {
	minHeads, maxHeads                       int
	allowComplexElement, forceComplexElement bool
	allowSpecialElement, forceSpecialElement bool
}

var DifficultyStageData = map[string]*[]stage{
	"easy": {
		0: {
			enemiesVariants: [][]*stageEnemyData{
				// variant 1
				{
					{minHeads: 3, maxHeads: 3},
				},
				// variant 2
				{
					{minHeads: 1, maxHeads: 2},
					{minHeads: 1, maxHeads: 2},
				},
			},
			treasure: 3,
		},
		1: {
			enemiesVariants: [][]*stageEnemyData{
				{
					{minHeads: 1, maxHeads: 3},
					{minHeads: 1, maxHeads: 3},
				},
				{
					{minHeads: 2, maxHeads: 4},
					{minHeads: 1, maxHeads: 2},
				},
			},
			treasure: 3,
		},
		2: {
			enemiesVariants: [][]*stageEnemyData{
				{
					{minHeads: 3, maxHeads: 5, allowComplexElement: true},
				},
			},
			treasure: 3,
		},
		3: {
			enemiesVariants: [][]*stageEnemyData{
				{
					{minHeads: 1, maxHeads: 3, allowComplexElement: true},
					{minHeads: 3, maxHeads: 5, allowComplexElement: true},
				},
			},
			treasure: 4,
		},
		4: {
			enemiesVariants: [][]*stageEnemyData{
				{
					{minHeads: 1, maxHeads: 5, allowComplexElement: true},
					{minHeads: 2, maxHeads: 5, allowComplexElement: true},
					{minHeads: 3, maxHeads: 5, forceComplexElement: true},
				},
			},
			treasure: 5,
		},
		5: {
			enemiesVariants: [][]*stageEnemyData{
				{
					{minHeads: 9, maxHeads: 16, forceComplexElement: true},
				},
				{
					{minHeads: 12, maxHeads: 16},
				},
			},
			treasure: 5,
		},
		6: {
			enemiesVariants: [][]*stageEnemyData{
				{
					{minHeads: 2, maxHeads: 5, allowComplexElement: true},
					{minHeads: 3, maxHeads: 6, forceComplexElement: true},
					{minHeads: 4, maxHeads: 7, forceComplexElement: true},
				},
			},
			treasure: 5,
		},
		7: {
			enemiesVariants: [][]*stageEnemyData{
				{
					{minHeads: 12, maxHeads: 30, forceSpecialElement: true},
				},
				{
					{minHeads: 8, maxHeads: 12, forceSpecialElement: true},
					{minHeads: 8, maxHeads: 12, forceSpecialElement: true},
				},
			},
		},
	},

	"medium": {
		0: {
			enemiesVariants: [][]*stageEnemyData{
				// variant 1
				{
					{minHeads: 3, maxHeads: 3},
				},
				// variant 2
				{
					{minHeads: 2, maxHeads: 3},
					{minHeads: 2, maxHeads: 3},
				},
			},
			treasure: 3,
		},
		1: {
			enemiesVariants: [][]*stageEnemyData{
				{
					{minHeads: 1, maxHeads: 3},
					{minHeads: 1, maxHeads: 3},
				},
				{
					{minHeads: 2, maxHeads: 4},
					{minHeads: 1, maxHeads: 2},
				},
			},
			treasure: 3,
		},
		2: {
			enemiesVariants: [][]*stageEnemyData{
				{
					{minHeads: 3, maxHeads: 5, allowComplexElement: true},
				},
			},
			treasure: 3,
		},
		3: {
			enemiesVariants: [][]*stageEnemyData{
				{
					{minHeads: 1, maxHeads: 3, allowComplexElement: true},
					{minHeads: 3, maxHeads: 5, allowComplexElement: true},
				},
			},
			treasure: 4,
		},
		4: {
			enemiesVariants: [][]*stageEnemyData{
				{
					{minHeads: 1, maxHeads: 5, allowComplexElement: true},
					{minHeads: 2, maxHeads: 5, allowComplexElement: true},
					{minHeads: 3, maxHeads: 5, forceComplexElement: true},
				},
			},
			treasure: 5,
		},
		5: {
			enemiesVariants: [][]*stageEnemyData{
				{
					{minHeads: 9, maxHeads: 16, forceComplexElement: true},
				},
				{
					{minHeads: 12, maxHeads: 16},
				},
			},
			treasure: 5,
		},
		6: {
			enemiesVariants: [][]*stageEnemyData{
				{
					{minHeads: 2, maxHeads: 5, allowComplexElement: true},
					{minHeads: 3, maxHeads: 6, forceComplexElement: true},
					{minHeads: 4, maxHeads: 7, forceComplexElement: true},
				},
			},
			treasure: 5,
		},
		7: {
			enemiesVariants: [][]*stageEnemyData{
				{
					{minHeads: 12, maxHeads: 30, forceSpecialElement: true},
				},
				{
					{minHeads: 8, maxHeads: 12, forceSpecialElement: true},
					{minHeads: 8, maxHeads: 12, forceSpecialElement: true},
				},
			},
		},
	},

	"hard": {
		{
			enemiesVariants: [][]*stageEnemyData{
				// variant 1
				{
					{minHeads: 3, maxHeads: 5, allowComplexElement: true},
				},
				// variant 2
				{
					{minHeads: 1, maxHeads: 5, allowComplexElement: true},
					{minHeads: 1, maxHeads: 5, allowComplexElement: true},
				},
			},
			treasure: 2,
		},
		{
			enemiesVariants: [][]*stageEnemyData{
				{
					{minHeads: 2, maxHeads: 5, allowComplexElement: true},
					{minHeads: 2, maxHeads: 5, allowComplexElement: true},
				},
				{
					{minHeads: 2, maxHeads: 7, allowComplexElement: true},
					{minHeads: 1, maxHeads: 5, allowComplexElement: true},
				},
			},
			treasure: 3,
		},
		{
			enemiesVariants: [][]*stageEnemyData{
				{
					{minHeads: 3, maxHeads: 5, allowComplexElement: true},
					{minHeads: 3, maxHeads: 5, allowComplexElement: true},
				},
			},
			treasure: 3,
		},
		{
			enemiesVariants: [][]*stageEnemyData{
				{
					{minHeads: 1, maxHeads: 3, allowComplexElement: true},
					{minHeads: 3, maxHeads: 5, allowComplexElement: true},
				},
			},
			treasure: 4,
		},
		{
			enemiesVariants: [][]*stageEnemyData{
				{
					{minHeads: 1, maxHeads: 5, allowComplexElement: true},
					{minHeads: 2, maxHeads: 5, allowComplexElement: true},
					{minHeads: 3, maxHeads: 5, forceComplexElement: true},
				},
			},
			treasure: 4,
		},
		{
			enemiesVariants: [][]*stageEnemyData{
				{
					{minHeads: 9, maxHeads: 16, forceComplexElement: true},
				},
				{
					{minHeads: 12, maxHeads: 16},
				},
			},
			treasure: 4,
		},
		{
			enemiesVariants: [][]*stageEnemyData{
				{
					{minHeads: 2, maxHeads: 5, allowComplexElement: true},
					{minHeads: 3, maxHeads: 6, forceComplexElement: true},
					{minHeads: 4, maxHeads: 7, forceComplexElement: true},
				},
			},
			treasure: 4,
		},
		{
			enemiesVariants: [][]*stageEnemyData{
				{
					{minHeads: 3, maxHeads: 6, allowComplexElement: true},
					{minHeads: 4, maxHeads: 7, forceComplexElement: true},
					{minHeads: 5, maxHeads: 8, forceComplexElement: true},
				},
			},
			treasure: 4,
		},
		{
			enemiesVariants: [][]*stageEnemyData{
				{
					{minHeads: 8, maxHeads: 16, allowComplexElement: true},
					{minHeads: 8, maxHeads: 16, forceComplexElement: true},
				},
			},
			treasure: 4,
		},
		{
			enemiesVariants: [][]*stageEnemyData{
				{
					{minHeads: 16, maxHeads: 32, forceSpecialElement: true},
				},
				{
					{minHeads: 12, maxHeads: 20, forceSpecialElement: true},
					{minHeads: 12, maxHeads: 20, forceSpecialElement: true},
				},
			},
		},
	},
}
