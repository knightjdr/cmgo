package gradient

import (
	"github.com/knightjdr/cmgo/read/localization"
	"github.com/knightjdr/cmgo/slice"
)

func adjacentSimilarity(baits []string, expected localization.ExpectedLocalizations) []int {
	similarity := make([]int, len(baits))

	totalBaits := len(baits)
	for i := range baits {
		similarCount := 0
		if i != 0 && slice.HasIntersect(expected[baits[i]].Terms, expected[baits[i-1]].Terms) {
			similarCount++
		}
		if i != totalBaits-1 && slice.HasIntersect(expected[baits[i]].Terms, expected[baits[i+1]].Terms) {
			similarCount++
		}
		similarity[i] = similarCount
	}

	return similarity
}
