package svg

import (
	"testing"

	"github.com/knightjdr/cmgo/internal/pkg/read/tsne"
	"github.com/stretchr/testify/assert"
)

func TestScaleCoordinates(t *testing.T) {
	coordinates := map[string]tsne.Coordinate{
		"AAAS": tsne.Coordinate{X: -50, Y: 25},
		"AAK1": tsne.Coordinate{X: -25, Y: -50},
		"AAR2": tsne.Coordinate{X: 150, Y: 75},
	}
	wanted := map[string]tsne.Coordinate{
		"AAAS": tsne.Coordinate{X: 100, Y: 400},
		"AAK1": tsne.Coordinate{X: 200, Y: 100},
		"AAR2": tsne.Coordinate{X: 900, Y: 600},
	}
	resultCoordinates, resultHeight := scaleCoordinates(coordinates, 1000, 24, 1)
	assert.Equal(t, wanted, resultCoordinates, "Should scale coordinates")
	assert.Equal(t, float64(700), resultHeight, "Should return new plot height")
}
