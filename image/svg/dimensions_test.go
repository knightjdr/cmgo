package svg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDimensions(t *testing.T) {
	matrix := [][]float64{
		{25, 5, 50.2},
		{100, 30, 7},
		{5, 2.3, 8},
	}
	columns := []string{"bait1", "bait2", "bait3"}
	rows := []string{"prey1", "prey2", "prey3"}

	// TEST: dimensions for full image with column and row names
	dims := dimensions(matrix, columns, rows)
	expected := heatmapDimensions{
		cellSize:   20,
		fontSize:   12,
		leftMargin: 57,
		plotHeight: 60,
		plotWidth:  60,
		ratio:      1,
		svgHeight:  117,
		svgWidth:   117,
		topMargin:  57,
	}
	assert.Equal(t, expected, dims, "Should return correct heat map dimensions")
}
