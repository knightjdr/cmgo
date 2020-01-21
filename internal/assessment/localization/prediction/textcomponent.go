package prediction

import "github.com/knightjdr/cmgo/internal/pkg/read/csv"

import "strconv"

func calculateTextComponent(options parameters, inputfiles fileContent) *preyTextScore {
	textAnnotations := readTextAnnotations(options.compartmentsText)

	return calculateTextScoreComponent(textAnnotations, inputfiles)
}

func readTextAnnotations(filename string) map[string]map[string]float64 {
	textAnnotations := make(map[string]map[string]float64)

	reader := csv.Read(filename, true)

	for {
		eof, line := csv.Readline(reader)
		if eof {
			break
		}

		gene := line[1]
		goID := line[2]
		score, _ := strconv.ParseFloat(line[5], 64)
		if _, ok := textAnnotations[gene]; !ok {
			textAnnotations[gene] = make(map[string]float64)
		}
		textAnnotations[gene][goID] = score
	}

	return textAnnotations

}
