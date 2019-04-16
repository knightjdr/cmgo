package stats

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinFloatSlice(t *testing.T) {
	slice := []float64{0.5, 0.3, 0.31, 4.1}
	assert.Equal(t, 0.3, MinFloatSlice(slice), "Should return the minimum value in a slice")
}
