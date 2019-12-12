package prediction

import (
	"strings"

	"github.com/knightjdr/cmgo/internal/pkg/read/csv"
)

func calculateStudyComponent(options parameters, inputfiles fileContent) *preyStudyScore {
	studyPredictions := map[string]map[string][]string{
		"hpa":           readFractionationPredictions(options.hpa, 2),
		"fractionation": readFractionationPredictions(options.fractionation, 5),
	}

	return calculateStudyComponentScore(studyPredictions, inputfiles)
}

func readFractionationPredictions(filename string, predictionIndex int) map[string][]string {
	predictions := make(map[string][]string, 0)

	reader := csv.Read(filename, false)
	for {
		eof, line := csv.Readline(reader)
		if eof {
			break
		}

		if line[predictionIndex] != "" {
			predictions[line[0]] = strings.Split(line[predictionIndex], ";")
		}
	}

	return predictions
}
