package prediction

type preyTextScore map[string]*textScoreComponents

type textScoreComponents struct {
	GOID  string
	score float64
}

func calculateTextScoreComponent(textAnnotations map[string]map[string]float64, inputfile fileContent) *preyTextScore {
	scores := &preyTextScore{}

	cellmapPredictions := getCellmapPredictions(inputfile.predictions, inputfile.predictionSummary)
	for prey, cellmapGoIDs := range cellmapPredictions {
		goID, score := findBestTextAnnotation(cellmapGoIDs, textAnnotations[prey])
		(*scores)[prey] = &textScoreComponents{
			GOID:  goID,
			score: score / 4,
		}
	}

	return scores
}

func findBestTextAnnotation(cellmapIDs []string, annotations map[string]float64) (string, float64) {
	if annotations == nil {
		return "", 0
	}

	bestID := ""
	max := float64(0)
	for _, cellmapID := range cellmapIDs {
		if _, ok := annotations[cellmapID]; ok && annotations[cellmapID] > max {
			bestID = cellmapID
			max = annotations[cellmapID]
		}
	}

	return bestID, max
}
