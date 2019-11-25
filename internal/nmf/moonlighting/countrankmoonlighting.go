package moonlighting

func countRankMoonlighting(moonlightingScores moonScores, numberOfRanks int, minRankValue float64) [][]int {
	rankMoonlighting := initializeAndAllocateRankMoonlighting(numberOfRanks)

	for _, prey := range moonlightingScores {
		if prey.SecondaryRank >= 0 && prey.SecondaryScore >= minRankValue {
			rankMoonlighting[prey.PrimaryRank][prey.SecondaryRank]++
		}
	}

	return rankMoonlighting
}

func initializeAndAllocateRankMoonlighting(numberOfRanks int) [][]int {
	rankMoonlighting := make([][]int, numberOfRanks)
	for i := 0; i < numberOfRanks; i++ {
		rankMoonlighting[i] = make([]int, numberOfRanks)
	}

	return rankMoonlighting
}
