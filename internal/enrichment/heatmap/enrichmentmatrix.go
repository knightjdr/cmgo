package heatmap

import (
	"sort"

	"github.com/knightjdr/cmgo/pkg/slice"
)

func enrichmentMatrix(enrichment map[string]map[string]float64) ([][]float64, []string, []string) {
	columns := make([]string, 0)
	rows := make([]string, 0)

	// Get column and row names;
	for compartment, regions := range enrichment {
		columns = append(columns, compartment)
		for region := range regions {
			rows = append(rows, region)
		}
	}
	columns = slice.UniqueStrings(columns)
	rows = slice.UniqueStrings(rows)
	sort.Strings(columns)
	sort.Strings(rows)

	// Create matrix.
	matrix := make([][]float64, len(rows))
	for rowNumber, region := range rows {
		matrix[rowNumber] = make([]float64, len(columns))
		for colNumber, compartment := range columns {
			matrix[rowNumber][colNumber] = enrichment[compartment][region]
		}
	}

	return matrix, columns, rows
}
