package svg

import (
	"testing"

	"github.com/knightjdr/cmgo/strfunc"
	"github.com/stretchr/testify/assert"
)

func TestRowNames(t *testing.T) {
	// Mock writeString.
	svg := make([]string, 0)
	writeString := func(str string) {
		svg = append(svg, str)
	}

	dims := heatmapDimensions{
		cellSize:   20,
		fontSize:   12,
		leftMargin: 50,
		topMargin:  50,
	}
	rows := []string{"prey1", "prey2", "prey3"}

	// TEST
	expected := "\t<g transform=\"translate(0, 50)\">\n" +
		"\t\t<text y=\"15\" x=\"48\" font-size=\"12\" text-anchor=\"end\">prey1</text>\n" +
		"\t\t<text y=\"35\" x=\"48\" font-size=\"12\" text-anchor=\"end\">prey2</text>\n" +
		"\t\t<text y=\"55\" x=\"48\" font-size=\"12\" text-anchor=\"end\">prey3</text>\n" +
		"\t</g>\n"
	rowNames(dims, rows, writeString)
	assert.Equal(t, expected, strfunc.Concat(svg), "Should return SVG row name element")
}
