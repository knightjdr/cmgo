package subset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithinThreshold(t *testing.T) {
	max := 0.5
	threshold := 0.5

	// TEST1: slice has a value within threshold of max
	rankValues := []float64{0.25, 0.1}
	assert.True(t, withinThreshold(max, rankValues, threshold), "Should return true when slice has a value within threshold of max")

	// TEST1: slice does not have a value within threshold of max
	rankValues = []float64{0.2, 0.1}
	assert.False(t, withinThreshold(max, rankValues, threshold), "Should return false when slice does not have a value within threshold of max")

}

func TestFilterByThreshold(t *testing.T) {
	basis := [][]float64{
		{0.5, 0.25, 0.3, 0.3},
		{0.25, 0.75, 0.5, 0.2},
		{0.8, 0.1, 0.3, 0.7},
		{0.9, 0.1, 0.75, 0.7},
	}
	rank1Indices := []int{0, 2}
	rank2Indices := []int{1}
	rows := []string{"a", "b", "c", "d"}
	wantedBasis := [][]float64{
		{0.5, 0.25, 0.3, 0.3},
		{0.25, 0.75, 0.5, 0.2},
	}
	wantedRows := []string{"a", "b"}
	basisResult, rowResult := filterByThreshold(basis, rows, rank1Indices, rank2Indices, 0.5)
	assert.Equal(t, wantedBasis, basisResult, "Should filter basis matrix")
	assert.Equal(t, wantedRows, rowResult, "Should return rows matching basis matrix")
}
