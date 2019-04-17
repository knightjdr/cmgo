package heatmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnrichmentMatrix(t *testing.T) {
	enrichment := map[string]map[string]float64{
		"1": map[string]float64{
			"PDZ":    8.906,
			"FERM_C": 13.120,
			"FERM_N": 12.391,
		},
		"2": map[string]float64{
			"KRAB":    9.905,
			"zf-C2H2": 5.333,
			"PDZ":     8.254,
		},
	}

	wantedColumns := []string{"1", "2"}
	wantedMatrix := [][]float64{
		{13.120, 0},
		{12.391, 0},
		{0, 9.905},
		{8.906, 8.254},
		{0, 5.333},
	}
	wantedRows := []string{"FERM_C", "FERM_N", "KRAB", "PDZ", "zf-C2H2"}
	matrix, columns, rows := enrichmentMatrix(enrichment)
	assert.Equal(t, wantedColumns, columns, "Should return columns sorted in order")
	assert.Equal(t, wantedMatrix, matrix, "Should convert 2D map to matrix")
	assert.Equal(t, wantedRows, rows, "Should return rows sorted in order")
}
