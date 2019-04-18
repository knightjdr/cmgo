package heatmap

import (
	"testing"

	"github.com/knightjdr/cmgo/strfunc"
	"github.com/stretchr/testify/assert"
)

func TestHeatmapRows(t *testing.T) {
	// Mock writeString.
	svg := make([]string, 0)
	writeString := func(str string) {
		svg = append(svg, str)
	}

	abundance := [][]float64{
		{25, 5, 50.2},
		{100, 30, 7},
		{5, 2.3, 8},
	}
	dims := heatmapDimensions{
		cellSize:   20,
		leftMargin: 50,
		topMargin:  50,
	}
	parameters := Settings{
		FillColor:    "blueBlack",
		AbundanceCap: 50,
		InvertColor:  false,
	}

	// TEST: create svg.
	heatmapRows(abundance, dims, parameters, writeString)
	expected := "\t<g id=\"minimap\" transform=\"translate(50, 50)\">\n" +
		"\t\t<rect fill=\"#0040ff\" y=\"0\" x=\"0\" width=\"20\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#ccd9ff\" y=\"0\" x=\"20\" width=\"20\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#000000\" y=\"0\" x=\"40\" width=\"20\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#000000\" y=\"20\" x=\"0\" width=\"20\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#0033cc\" y=\"20\" x=\"20\" width=\"20\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#b8c9ff\" y=\"20\" x=\"40\" width=\"20\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#ccd9ff\" y=\"40\" x=\"0\" width=\"20\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#e6ecff\" y=\"40\" x=\"20\" width=\"20\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#adc2ff\" y=\"40\" x=\"40\" width=\"20\" height=\"20\" />\n" +
		"\t</g>\n"
	assert.Equal(t, expected, strfunc.Concat(svg), "Should return SVG cell element")
}
