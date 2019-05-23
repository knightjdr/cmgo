package svg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScaleCoordinates(t *testing.T) {
	coordinates := map[string]coordinate{
		"AAAS": coordinate{X: -50, Y: 25},
		"AAK1": coordinate{X: -25, Y: -50},
		"AAR2": coordinate{X: 150, Y: 75},
	}
	wanted := map[string]coordinate{
		"AAAS": coordinate{X: 0, Y: 375},
		"AAK1": coordinate{X: 125, Y: 0},
		"AAR2": coordinate{X: 1000, Y: 625},
	}
	resultCoordinates, resultHeight := scaleCoordinates(coordinates, 1000)
	assert.Equal(t, wanted, resultCoordinates, "Should scale coordinates")
	assert.Equal(t, float64(625), resultHeight, "Should return new plot height")
}
