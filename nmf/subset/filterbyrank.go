package subset

import (
	"github.com/knightjdr/cmgo/stats"
)

func filterByRank(basis [][]float64, rows []string, rank1Indices, rank2Indices []int, minNMFScore float64) ([][]float64, []string) {
	filtered := make([][]float64, 0)
	filteredRows := make([]string, 0)

	// Merge ranks indices.
	rankIndices := append(rank1Indices, rank2Indices...)

	for i, basisRow := range basis {
		// Maximum in row.
		maxRow := stats.MaxFloat(basisRow)

		// Get desired rank values and get the maximum of all.
		rankValues := make([]float64, len(rankIndices))
		for j, column := range rankIndices {
			rankValues[j] = basisRow[column]
		}
		maxRank := stats.MaxFloat(rankValues)

		// Keep the row if it has a rank value matching the row maximum.
		if maxRank == maxRow && maxRank >= minNMFScore {
			filtered = append(filtered, basisRow)
			filteredRows = append(filteredRows, rows[i])
		}
	}

	return filtered, filteredRows
}
