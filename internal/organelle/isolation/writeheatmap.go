package isolation

import (
	"fmt"
	"strings"

	"github.com/knightjdr/cmgo/internal/pkg/image/svg/heatmap"
	"github.com/knightjdr/cmgo/internal/pkg/read/localization"
)

func writeHeatmap(scores *isolationScores, nmfSummary localization.Summary, options parameters) {
	compartmentNames := getCompartmentNames(nmfSummary)
	fileHandle := strings.Split(options.svgFile, ".svg")[0]

	absoluteMatrix, relativeMatrix := createMatrices(scores, len(compartmentNames))

	absoluteSettings := heatmap.Settings{
		Filename:     fmt.Sprintf("%s-absolute.svg", fileHandle),
		FillColor:    "blueBlack",
		AbundanceCap: options.abundanceCap,
	}
	heatmap.Draw(absoluteMatrix, compartmentNames, compartmentNames, absoluteSettings)

	relativeSettings := heatmap.Settings{
		Filename:     fmt.Sprintf("%s-relative.svg", fileHandle),
		FillColor:    "blueBlack",
		AbundanceCap: 1,
	}
	heatmap.Draw(relativeMatrix, compartmentNames, compartmentNames, relativeSettings)
}

func getCompartmentNames(nmfSummary localization.Summary) []string {
	noCompartments := len(nmfSummary)

	compartments := make([]string, noCompartments)
	for i := 0; i < noCompartments; i++ {
		compartments[i] = strings.Join(nmfSummary[i+1].DisplayTerms, ", ")
	}

	return compartments
}

func createMatrices(scores *isolationScores, noCompartments int) ([][]float64, [][]float64) {
	absoluteMatrix := make([][]float64, noCompartments)
	relativeMatrix := make([][]float64, noCompartments)

	for compartment, score := range *scores {
		rowIndex := compartment - 1
		absoluteMatrix[rowIndex] = make([]float64, noCompartments)
		relativeMatrix[rowIndex] = make([]float64, noCompartments)

		for i := 1; i <= noCompartments; i++ {
			columnIndex := i - 1
			absoluteMatrix[rowIndex][columnIndex] = float64(score.sharedCompartments[i])
			relativeMatrix[rowIndex][columnIndex] = float64(score.sharedCompartments[i]) / float64(score.edgesOutside)
		}
	}

	return absoluteMatrix, relativeMatrix
}
