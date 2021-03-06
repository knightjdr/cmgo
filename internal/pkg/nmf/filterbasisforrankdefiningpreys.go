// Package nmf implements common functions for NMF analysis.
package nmf

import (
	"sort"

	"github.com/knightjdr/cmgo/pkg/mapfunc"
	customMath "github.com/knightjdr/cmgo/pkg/math"
	customSort "github.com/knightjdr/cmgo/pkg/sort"
)

// FilterBasisForRankDefiningPreys filters a basis matrix to return genes used for defining localization.
func FilterBasisForRankDefiningPreys(matrix [][]float64, maxGenesPerRank int, minRankValue, withinRankMax float64) [][]int {
	// For each rank store all genes that satisfy filtering criteria.
	filteredRanks := make(map[int]map[int]float64)
	for i := 0; i < len(matrix[0]); i++ {
		filteredRanks[i] = make(map[int]float64)
	}
	for i, row := range matrix {
		max := customMath.MaxSliceFloat(row)
		for j, value := range row {
			if value >= minRankValue && value/max >= withinRankMax {
				filteredRanks[j][i] = value
			}
		}
	}

	// Only keep at most top maxGenesPerRank genes for each rank.
	topRankRows := make([][]int, len(filteredRanks))
	for i, genes := range filteredRanks {
		if len(genes) <= maxGenesPerRank {
			topRankRows[i] = make([]int, len(genes))
			topRankRows[i] = mapfunc.KeysIntFloat(genes)
		} else {
			topRankRows[i] = make([]int, maxGenesPerRank)
			sorted := customSort.ByMapValueIntFloat64(genes, "descending")
			for j := 0; j < maxGenesPerRank; j++ {
				topRankRows[i][j] = sorted[j].Key
			}
		}
		sort.Ints(topRankRows[i])
	}

	return topRankRows
}
