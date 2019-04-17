package svg

import (
	"testing"

	"github.com/knightjdr/cmgo/strfunc"
	"github.com/stretchr/testify/assert"
)

func TestHeader(t *testing.T) {
	// Mock writeString.
	svg := make([]string, 0)
	writeString := func(str string) {
		svg = append(svg, str)
	}

	// TEST
	dims := heatmapDimensions{
		svgHeight: 250,
		svgWidth:  150,
	}
	expected := "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"" +
		" xml:space=\"preserve\" width=\"150\" height=\"250\" viewBox=\"0 0 150 250\">\n"
	header(dims, writeString)
	assert.Equal(t, expected, strfunc.Concat(svg), "Should return SVG header")
}
