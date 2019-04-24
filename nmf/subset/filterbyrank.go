package subset

import (
	"github.com/knightjdr/cmgo/stats"
)

func filterByRank(basis [][]float64, rows []string, rank1Indices, rank2Indices []int, minNMFScore float64) ([][]float64, []string) {
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
		maxRank := stats.MaxFloatSlice(append(rank1Values, rank2Values...))

		// Keep the row if it has a rank value matching the row maximum.
		if maxRank == maxRow && maxRank >= minNMFScore {
			filtered = append(filtered, basisRow)
			filteredRows = append(filteredRows, rows[i])
		}
	}

	return filtered, filteredRows
}
