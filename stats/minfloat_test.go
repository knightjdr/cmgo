package stats

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinFloat(t *testing.T) {
	slice := []float64{0.5, 0.3, 0.31, 4.1}
	assert.Equal(t, 0.3, MinFloat(slice), "Should return the minimum value in a slice")
}
