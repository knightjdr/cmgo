package stats

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMeanFloat(t *testing.T) {
	slice := []float64{7, 6, 3, 9, 11}
	assert.Equal(t, float64(7.2), MeanFloat(slice), "Should return mean of a slice")
}
