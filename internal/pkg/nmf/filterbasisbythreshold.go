// Package nmf implements common functions for NMF analysis.
package nmf

import (
	"github.com/knightjdr/cmgo/pkg/slice"
)

// FilterBasisByTreshold filters a basis matrix to remove preys not passing minimum rank value
func FilterBasisByTreshold(matrix [][]float64, rowNames []string, minRankValue float64) ([][]float64, []string) {
	filterFunc := generateFilterFunc(minRankValue)

	filteredMatrix := make([][]float64, 0)
	filteredRowNames := make([]string, 0)
	for i, row := range matrix {
		if slice.SomeFloat(row, filterFunc) {
			filteredMatrix = append(filteredMatrix, row)
			filteredRowNames = append(filteredRowNames, rowNames[i])
		}
	}
	return filteredMatrix, filteredRowNames
}

func generateFilterFunc(threshold float64) func(value float64) bool {
	return func(value float64) bool {
		if value >= threshold {
			return true
		}
		return false
	}
}
