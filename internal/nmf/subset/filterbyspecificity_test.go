package subset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterBySpecificity(t *testing.T) {
	basis := [][]float64{
		{0.5, 0.1, 0.25, 0.05},
		{0.25, 0.1, 0.5, 0.2},
		{0.4, 0.1, 0.3, 0.7},
		{0.9, 0.1, 0.75, 0.03},
		{0, 0.05, 0.1, 0},
	}
	rank1Indices := []int{0, 2}
	rank2Indices := []int{1}
	rows := []string{"a", "b", "c", "d", "e"}
	wantedBasis := [][]float64{
		{0.5, 0.1, 0.25, 0.05},
		{0.9, 0.1, 0.75, 0.03},
		{0, 0.05, 0.1, 0},
	}
	wantedRows := []string{"a", "d", "e"}
	basisResult, rowResult := filterBySpecificity(basis, rows, rank1Indices, rank2Indices, 2)
	assert.Equal(t, wantedBasis, basisResult, "Should filter basis matrix")
	assert.Equal(t, wantedRows, rowResult, "Should return rows matching basis matrix")
}
