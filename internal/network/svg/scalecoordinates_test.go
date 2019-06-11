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
		"AAAS": coordinate{X: 100, Y: 400},
		"AAK1": coordinate{X: 200, Y: 100},
		"AAR2": coordinate{X: 900, Y: 600},
	}
	resultCoordinates, resultHeight := scaleCoordinates(coordinates, 1000, 24, 1)
	assert.Equal(t, wanted, resultCoordinates, "Should scale coordinates")
	assert.Equal(t, float64(700), resultHeight, "Should return new plot height")
}
