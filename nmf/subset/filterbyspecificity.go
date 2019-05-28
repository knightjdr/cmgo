package subset

import (
	"github.com/knightjdr/cmgo/stats"
)

func filterBySpecificity(basis [][]float64, rows []string, rank1Indices, rank2Indices []int, specificity float64) ([][]float64, []string) {
	filtered := make([][]float64, 0)
	filteredRows := make([]string, 0)

	// Merge ranks indices.
	rankIndices := append(rank1Indices, rank2Indices...)

	// Create dictionary of desired ranks.
	rankMap := make(map[int]bool, len(rankIndices))
	for _, column := range append(rankIndices) {
		rankMap[column] = true
	}

	for i, basisRow := range basis {
		// Find desired rank values and get the minimum of all excluding zeros.
		rankValues := make([]float64, 0)
		for _, column := range rankIndices {
			if basisRow[column] != 0 {
				rankValues = append(rankValues, basisRow[column])
			}
		}
		minRank := stats.MinFloat(append(rankValues))

		// Find other rank values and get the maximum of all.
		otherValues := make([]float64, 0)
		for j, value := range basisRow {
			if _, ok := rankMap[j]; !ok {
				otherValues = append(otherValues, value)
			}
		}
		maxOther := stats.MaxFloat(otherValues)

		// Keep the row if the fold-change of the rank value minimum relative to
		// the rank value maximum is at least the "specificty" argument.
		if maxOther == 0 || minRank/maxOther >= specificity {
			filtered = append(filtered, basisRow)
			filteredRows = append(filteredRows, rows[i])
		}
	}

	return filtered, filteredRows
}
