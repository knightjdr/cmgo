package stats

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMeanFloat(t *testing.T) {
	// TEST1: slice of floats
	slice := []float64{7, 6, 3, 9, 11}
	assert.Equal(t, 7.2, MeanFloat(slice), "Should return mean of a slice of floats")

	// TEST2: nil slice
	slice = []float64{}
	assert.Equal(t, float64(0), MeanFloat(slice), "Should return zero for empty slice")
}

func TestMeanInt(t *testing.T) {
	slice := []int{7, 6, 3, 9, 11}
	assert.Equal(t, 7.2, MeanInt(slice), "Should return mean of a slice of ints")

	// TEST2: nil slice
	slice = []int{}
	assert.Equal(t, float64(0), MeanInt(slice), "Should return zero for empty slice")
}
