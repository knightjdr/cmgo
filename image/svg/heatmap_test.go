package svg

import (
	"testing"

	"github.com/knightjdr/cmgo/fs"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestHeatmap(t *testing.T) {
	// Mock fs.
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Create test directory.
	fs.Instance.MkdirAll("test", 0755)

	matrix := [][]float64{
		{25, 5, 50.2},
		{100, 30, 7},
		{5, 2.3, 8},
	}
	columns := []string{"bait1", "bait2", "bait3"}
	rows := []string{"prey1", "prey2", "prey3"}
	parameters := HeatmapSettings{
		FillColor:    "blueBlack",
		AbundanceCap: 50,
		MinAbundance: 0,
		InvertColor:  false,
	}

	wanted := "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" xml:space=\"preserve\" width=\"117\" height=\"117\" viewBox=\"0 0 117 117\">\n" +
		"\t<g transform=\"translate(57)\">\n" +
		"\t\t<text y=\"55\" x=\"6\" font-size=\"12\" text-anchor=\"end\" transform=\"rotate(90, 6, 55)\">bait1</text>\n" +
		"\t\t<text y=\"55\" x=\"26\" font-size=\"12\" text-anchor=\"end\" transform=\"rotate(90, 26, 55)\">bait2</text>\n" +
		"\t\t<text y=\"55\" x=\"46\" font-size=\"12\" text-anchor=\"end\" transform=\"rotate(90, 46, 55)\">bait3</text>\n" +
		"\t</g>\n" +
		"\t<g transform=\"translate(0, 57)\">\n" +
		"\t\t<text y=\"15\" x=\"55\" font-size=\"12\" text-anchor=\"end\">prey1</text>\n" +
		"\t\t<text y=\"35\" x=\"55\" font-size=\"12\" text-anchor=\"end\">prey2</text>\n" +
		"\t\t<text y=\"55\" x=\"55\" font-size=\"12\" text-anchor=\"end\">prey3</text>\n" +
		"\t</g>\n" +
		"\t<g id=\"minimap\" transform=\"translate(57, 57)\">\n" +
		"\t\t<rect fill=\"#0040ff\" y=\"0\" x=\"0\" width=\"20\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#ccd9ff\" y=\"0\" x=\"20\" width=\"20\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#000000\" y=\"0\" x=\"40\" width=\"20\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#000000\" y=\"20\" x=\"0\" width=\"20\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#0033cc\" y=\"20\" x=\"20\" width=\"20\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#b8c9ff\" y=\"20\" x=\"40\" width=\"20\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#ccd9ff\" y=\"40\" x=\"0\" width=\"20\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#e6ecff\" y=\"40\" x=\"20\" width=\"20\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#adc2ff\" y=\"40\" x=\"40\" width=\"20\" height=\"20\" />\n" +
		"\t</g>\n" +
		"</svg>\n"
	svg := Heatmap(matrix, columns, rows, parameters)
	assert.Equal(t, wanted, svg, "Should return SVG string")

	// TEST: full heatmap image writing to file
	parameters.Filename = "test/heatmap.svg"
	wanted = "<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" xml:space=\"preserve\" width=\"117\" height=\"117\" viewBox=\"0 0 117 117\">\n" +
		"\t<g transform=\"translate(57)\">\n" +
		"\t\t<text y=\"55\" x=\"6\" font-size=\"12\" text-anchor=\"end\" transform=\"rotate(90, 6, 55)\">bait1</text>\n" +
		"\t\t<text y=\"55\" x=\"26\" font-size=\"12\" text-anchor=\"end\" transform=\"rotate(90, 26, 55)\">bait2</text>\n" +
		"\t\t<text y=\"55\" x=\"46\" font-size=\"12\" text-anchor=\"end\" transform=\"rotate(90, 46, 55)\">bait3</text>\n" +
		"\t</g>\n" +
		"\t<g transform=\"translate(0, 57)\">\n" +
		"\t\t<text y=\"15\" x=\"55\" font-size=\"12\" text-anchor=\"end\">prey1</text>\n" +
		"\t\t<text y=\"35\" x=\"55\" font-size=\"12\" text-anchor=\"end\">prey2</text>\n" +
		"\t\t<text y=\"55\" x=\"55\" font-size=\"12\" text-anchor=\"end\">prey3</text>\n" +
		"\t</g>\n" +
		"\t<g id=\"minimap\" transform=\"translate(57, 57)\">\n" +
		"\t\t<rect fill=\"#0040ff\" y=\"0\" x=\"0\" width=\"20\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#ccd9ff\" y=\"0\" x=\"20\" width=\"20\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#000000\" y=\"0\" x=\"40\" width=\"20\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#000000\" y=\"20\" x=\"0\" width=\"20\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#0033cc\" y=\"20\" x=\"20\" width=\"20\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#b8c9ff\" y=\"20\" x=\"40\" width=\"20\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#ccd9ff\" y=\"40\" x=\"0\" width=\"20\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#e6ecff\" y=\"40\" x=\"20\" width=\"20\" height=\"20\" />\n" +
		"\t\t<rect fill=\"#adc2ff\" y=\"40\" x=\"40\" width=\"20\" height=\"20\" />\n" +
		"\t</g>\n" +
		"</svg>\n"
	svg = Heatmap(matrix, columns, rows, parameters)
	bytes, _ := afero.ReadFile(fs.Instance, "test/heatmap.svg")
	assert.Equal(t, "", svg, "Should return null when writing to file")
	assert.Equal(t, wanted, string(bytes), "Should write SVG to file")
}
