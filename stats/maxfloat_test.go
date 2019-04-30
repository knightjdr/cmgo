package stats

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxFloatSlice(t *testing.T) {
	slice := []float64{0.5, 0.3, 0.31, 4.1}
	assert.Equal(t, 4.1, MaxFloat(slice), "Should return the maximum value in a slice")
}
