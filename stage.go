package main

type dungeon struct {
	name           string
	stagesVariants [][]*stage
	totalStages    int
}

type stage struct {
	name     string
	enemies  []*stageEnemyData
	treasure int
}

type stageEnemyData struct {
	minHeads, maxHeads                       int
	allowComplexElement, forceComplexElement bool
	allowSpecialElement, forceSpecialElement bool
}

func (d *dungeon) getTotalStages() int {
	if d.totalStages != 0 {
		return d.totalStages
	}
	return len(d.stagesVariants)
}

func (d *dungeon) getStageNumber(num int) *stage {
	var s *stage
	if num < len(d.stagesVariants) {
		s = d.stagesVariants[num][rnd.Rand(len(d.stagesVariants[num]))]
	}
	if s != nil {
		return s
	}
	// generate random enemies data
	var sed []*stageEnemyData
	numEnemies := rnd.RandInRange(1, num)
	for i := 0; i < numEnemies; i++ {
		sed = append(sed, &stageEnemyData{
			minHeads:            num + 1,
			maxHeads:            num + 1 + rnd.Rand(num/2+1),
			allowComplexElement: num > 3,
			allowSpecialElement: num > 5,
		})
	}
	return &stage{
		name:     "Random stage",
		enemies:  sed,
		treasure: numEnemies + 1,
	}
}

var dungeons = map[string]*dungeon{
	"easy": {
		name: "easy",
		stagesVariants: [][]*stage{
			0: {
				{
					enemies: []*stageEnemyData{
						{minHeads: 3, maxHeads: 4},
					},
					treasure: 3,
				},
				{
					enemies: []*stageEnemyData{
						{minHeads: 1, maxHeads: 2},
						{minHeads: 1, maxHeads: 2},
					},
					treasure: 3,
				},
			},
			1: {
				{
					enemies: []*stageEnemyData{
						{minHeads: 1, maxHeads: 3},
						{minHeads: 1, maxHeads: 3},
					},
					treasure: 3,
				},
				{
					enemies: []*stageEnemyData{
						{minHeads: 2, maxHeads: 4},
						{minHeads: 1, maxHeads: 2},
					},
					treasure: 3,
				},
			},
			2: {
				{
					enemies: []*stageEnemyData{
						{minHeads: 3, maxHeads: 7, allowComplexElement: true},
					},
					treasure: 4,
				},
			},
			3: {
				{
					enemies: []*stageEnemyData{
						{minHeads: 3, maxHeads: 6, allowComplexElement: true},
						{minHeads: 3, maxHeads: 6, allowComplexElement: true},
					},
					treasure: 5,
				},
			},
			4: {
				{
					enemies: []*stageEnemyData{
						{minHeads: 3, maxHeads: 6, allowComplexElement: true},
						{minHeads: 2, maxHeads: 6, allowComplexElement: true},
						{minHeads: 1, maxHeads: 6, allowComplexElement: true},
					},
					treasure: 5,
				},
			},
			5: {
				{
					enemies: []*stageEnemyData{
						{minHeads: 17, maxHeads: 25, forceSpecialElement: true},
					},
				},
			},
		},
	},

	"medium": {
		name: "medium",
		stagesVariants: [][]*stage{
			0: {
				{
					enemies: []*stageEnemyData{
						{minHeads: 3, maxHeads: 4},
					},
					treasure: 3,
				},
				{
					enemies: []*stageEnemyData{
						{minHeads: 1, maxHeads: 2},
						{minHeads: 1, maxHeads: 2},
					},
					treasure: 3,
				},
			},
			1: {
				{
					enemies: []*stageEnemyData{
						{minHeads: 1, maxHeads: 3},
						{minHeads: 1, maxHeads: 3},
					},
					treasure: 3,
				},
				{
					enemies: []*stageEnemyData{
						{minHeads: 2, maxHeads: 4},
						{minHeads: 1, maxHeads: 2},
					},
					treasure: 3,
				},
			},
			2: {
				{
					enemies: []*stageEnemyData{
						{minHeads: 3, maxHeads: 7, allowComplexElement: true},
					},
					treasure: 4,
				},
			},
			3: {
				{
					enemies: []*stageEnemyData{
						{minHeads: 3, maxHeads: 6, allowComplexElement: true},
						{minHeads: 3, maxHeads: 6, allowComplexElement: true},
					},
					treasure: 4,
				},
			},
			4: {
				{
					enemies: []*stageEnemyData{
						{minHeads: 3, maxHeads: 6, allowComplexElement: true},
						{minHeads: 2, maxHeads: 6, allowComplexElement: true},
						{minHeads: 1, maxHeads: 6, allowComplexElement: true},
					},
					treasure: 4,
				},
			},
			5: nil,
			6: nil,
			7: {
				{
					enemies: []*stageEnemyData{
						{minHeads: 17, maxHeads: 25, forceSpecialElement: true},
					},
					treasure: 5,
				},
			},
			8: {
				{
					enemies: []*stageEnemyData{
						{minHeads: 17, maxHeads: 25, forceSpecialElement: true},
						{minHeads: 17, maxHeads: 25, forceSpecialElement: true},
					},
				},
			},
		},
	},

	"hard": {
		name: "hard",
		totalStages: 15,
		stagesVariants: [][]*stage{
			0: {
				{
					enemies: []*stageEnemyData{
						{minHeads: 3, maxHeads: 4},
					},
					treasure: 2,
				},
				{
					enemies: []*stageEnemyData{
						{minHeads: 1, maxHeads: 2},
						{minHeads: 1, maxHeads: 2},
					},
					treasure: 2,
				},
			},
			1: {
				{
					enemies: []*stageEnemyData{
						{minHeads: 1, maxHeads: 3},
						{minHeads: 1, maxHeads: 3},
					},
					treasure: 3,
				},
				{
					enemies: []*stageEnemyData{
						{minHeads: 2, maxHeads: 4},
						{minHeads: 1, maxHeads: 2},
					},
					treasure: 3,
				},
			},
			2: {
				{
					enemies: []*stageEnemyData{
						{minHeads: 3, maxHeads: 7, allowComplexElement: true},
					},
					treasure: 4,
				},
			},
			3: {
				{
					enemies: []*stageEnemyData{
						{minHeads: 3, maxHeads: 6, allowComplexElement: true},
						{minHeads: 3, maxHeads: 6, allowComplexElement: true},
					},
					treasure: 4,
				},
			},
			4: {
				{
					enemies: []*stageEnemyData{
						{minHeads: 3, maxHeads: 6, allowComplexElement: true},
						{minHeads: 2, maxHeads: 6, allowComplexElement: true},
						{minHeads: 1, maxHeads: 6, allowComplexElement: true},
					},
					treasure: 4,
				},
			},
			5: {
				{
					enemies: []*stageEnemyData{
						{minHeads: 17, maxHeads: 25, forceSpecialElement: true},
					},
					treasure: 5,
				},
			},
		},
	},

	"chaotic": {
		name: "chaotic",
		totalStages: 15,
		stagesVariants: [][]*stage{
		},
	},
}
