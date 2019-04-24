package subset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterByRank(t *testing.T) {
	basis := [][]float64{
		{0.5, 0.1, 0.25, 0.3},
		{0.25, 0.1, 0.5, 0.2},
		{0.4, 0.1, 0.3, 0.7},
		{0.9, 0.1, 0.75, 0.7},
		{0.1, 0.05, 0.1, 0.01},
	}
	rank1Indices := []int{0, 2}
	rank2Indices := []int{1}
	rows := []string{"a", "b", "c", "d"}
	wantedBasis := [][]float64{
		{0.5, 0.1, 0.25, 0.3},
		{0.25, 0.1, 0.5, 0.2},
		{0.9, 0.1, 0.75, 0.7},
	}
	wantedRows := []string{"a", "b", "d"}
	basisResult, rowResult := filterByRank(basis, rows, rank1Indices, rank2Indices, 0.15)
	assert.Equal(t, wantedBasis, basisResult, "Should filter basis matrix")
	assert.Equal(t, wantedRows, rowResult, "Should return rows matching basis matrix")
}
