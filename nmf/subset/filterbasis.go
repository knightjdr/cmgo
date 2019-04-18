package subset

import (
	"github.com/knightjdr/cmgo/stats"
)

func withinLimit(max float64, rankValues []float64, threshold float64) bool {
	checkLimit := false
	for _, value := range rankValues {
		if (value / max) < threshold {
			checkLimit = true
			break
		}
	}
	return checkLimit
}

func filterBasis(basis [][]float64, rows []string, rank1Indices, rank2Indices []int, threshold float64) ([][]float64, []string) {
	filtered := make([][]float64, 0)
	filteredRows := make([]string, 0)

	for i, basisRow := range basis {
		// Maximum in row.
		maxRow := stats.MaxFloatSlice(basisRow)

		// Get desired rank values and get the maximum of all.
		rank1Values := make([]float64, len(rank1Indices))
		rank2Values := make([]float64, len(rank2Indices))
		for j, column := range rank1Indices {
			rank1Values[j] = basisRow[column]
		}
		for j, column := range rank2Indices {
			rank2Values[j] = basisRow[column]
		}
		max := stats.MaxFloatSlice(append(rank1Values, rank2Values...))

		// If a rank value is the maximum in the row, check if ranks in other compartment
		// are within threshold
		if max == maxRow {
			max1 := stats.MaxFloatSlice(rank1Values)
			max2 := stats.MaxFloatSlice(rank2Values)
			checkLimit := false
			if max1 > max2 {
				checkLimit = withinLimit(max1, rank2Values, threshold)
			} else {
				checkLimit = withinLimit(max2, rank1Values, threshold)
			}
			if checkLimit {
				filtered = append(filtered, basisRow)
				filteredRows = append(filteredRows, rows[i])
			}
		}
	}

	return filtered, filteredRows
}
