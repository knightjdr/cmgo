package subset

import (
	"github.com/knightjdr/cmgo/stats"
)

func withinThreshold(max float64, rankValues []float64, threshold float64) bool {
	checkLimit := false
	for _, value := range rankValues {
		if (value / max) >= threshold {
			checkLimit = true
			break
		}
	}
	return checkLimit
}

func filterByThreshold(basis [][]float64, rows []string, rank1Indices, rank2Indices []int, threshold float64) ([][]float64, []string) {
	filtered := make([][]float64, 0)
	filteredRows := make([]string, 0)

	for i, basisRow := range basis {
		// Get desired rank values.
		rank1Values := make([]float64, len(rank1Indices))
		rank2Values := make([]float64, len(rank2Indices))
		for j, column := range rank1Indices {
			rank1Values[j] = basisRow[column]
		}
		for j, column := range rank2Indices {
			rank2Values[j] = basisRow[column]
		}

		// Check if ranks in other compartment are within threshold
		max1 := stats.MaxFloatSlice(rank1Values)
		max2 := stats.MaxFloatSlice(rank2Values)
		checkLimit := false
		if max1 > max2 {
			checkLimit = withinThreshold(max1, rank2Values, threshold)
		} else {
			checkLimit = withinThreshold(max2, rank1Values, threshold)
		}
		if checkLimit {
			filtered = append(filtered, basisRow)
			filteredRows = append(filteredRows, rows[i])
		}
	}

	return filtered, filteredRows
}
