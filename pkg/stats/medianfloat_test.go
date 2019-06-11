package stats

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMedianFloat(t *testing.T) {
	// TEST1: odd length slice
	slice := []float64{7, 6, 3, 9, 12}
	assert.Equal(t, float64(7), MedianFloat(slice), "Should return median of odd length slice")

	// TEST2: even length slice
	slice = []float64{7, 6, 9, 12}
	assert.Equal(t, float64(8), MedianFloat(slice), "Should return median of even length slice")
}
