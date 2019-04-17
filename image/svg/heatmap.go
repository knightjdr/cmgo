// Package svg creates svg files for various image types.
package svg

import (
	"github.com/knightjdr/cmgo/image/file"
	"github.com/knightjdr/cmgo/strfunc"
	"github.com/spf13/afero"
)

// HeatmapSettings configuration for drawing heatmap
type HeatmapSettings struct {
	Filename     string
	FillColor    string
	AbundanceCap float64
	InvertColor  bool
	MinAbundance float64
}

func write(svg *[]string, file afero.File) func(str string) {
	if file != nil {
		return func(str string) {
			file.WriteString(str)
		}
	}
	return func(str string) {
		*svg = append(*svg, str)
	}
}

// Heatmap creates a heatmap from an input matrix of abundance.
func Heatmap(matrix [][]float64, columns, rows []string, parameters HeatmapSettings) string {
	svg := make([]string, 0)
	dims := dimensions(matrix, columns, rows)

	// Open file for writing if requested
	file, _ := file.Create(parameters.Filename)
	if file != nil {
		defer file.Close()
	}

	writeString := write(&svg, file)

	header(dims, writeString)
	columnNames(dims, columns, writeString)
	rowNames(dims, rows, writeString)
	heatmapRows(matrix, dims, parameters, writeString)

	// Add end element wrapper for svg.
	writeString("</svg>\n")
	return strfunc.Concat(svg)
}
