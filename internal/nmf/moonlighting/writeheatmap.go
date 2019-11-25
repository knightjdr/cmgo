package moonlighting

import (
	"fmt"
	"math"
	"path"
	"strconv"

	"github.com/knightjdr/cmgo/internal/pkg/image/svg/heatmap"
	"github.com/knightjdr/cmgo/internal/pkg/image/svg/legend"
	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/spf13/afero"
)

func writeHeatmap(matrix [][]int, outfile string) {
	abundanceCap := float64(findMaxValue(matrix))
	floatMatrix := convertMatrixToFloat(matrix)
	labels := createRankLabels(len(matrix))

	settings := heatmap.Settings{
		Filename:     outfile,
		FillColor:    "blueBlack",
		AbundanceCap: abundanceCap,
		MinAbundance: 0,
	}
	heatmap.Draw(floatMatrix, labels, labels, settings)
	writeLegend(abundanceCap, outfile)
}

func findMaxValue(matrix [][]int) int {
	max := math.MinInt64
	for _, row := range matrix {
		for _, value := range row {
			if value > max {
				max = value
			}
		}
	}
	return max
}

func convertMatrixToFloat(matrix [][]int) [][]float64 {
	floatMatrix := make([][]float64, len(matrix))
	for i, row := range matrix {
		floatMatrix[i] = make([]float64, len(row))
		for j, value := range row {
			floatMatrix[i][j] = float64(value)
		}
	}
	return floatMatrix
}

func createRankLabels(numberOfRanks int) []string {
	labels := make([]string, numberOfRanks)

	for i := 0; i < numberOfRanks; i++ {
		labels[i] = strconv.Itoa(i + 1)
	}

	return labels
}

func writeLegend(abundanceCap float64, outfileHeatmap string) {
	legendSVG := legend.Gradient("blueBlack", "moonlighting", 101, 0, abundanceCap, false)
	dir := path.Dir(outfileHeatmap)
	outfile := fmt.Sprintf("%s/legend.svg", dir)
	afero.WriteFile(fs.Instance, outfile, []byte(legendSVG), 0644)
}
