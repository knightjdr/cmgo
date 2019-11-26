package nmf

import (
	"github.com/knightjdr/cmgo/pkg/math"
)

// GetPreysPerRank gets a list of the preys localizing to each rank.
func GetPreysPerRank(basis [][]float64, preys []string) map[int][]string {
	preysPerRank := make(map[int][]string, 0)

	for rowIndex, row := range basis {
		maxIndex := math.MaxIndexSliceFloat(row) + 1
		allocatePreysPerRankMemory(preysPerRank, maxIndex)
		preysPerRank[maxIndex] = append(preysPerRank[maxIndex], preys[rowIndex])
	}

	return preysPerRank
}

func allocatePreysPerRankMemory(preysPerRank map[int][]string, rank int) {
	if _, ok := preysPerRank[rank]; !ok {
		preysPerRank[rank] = make([]string, 0)
	}
}
