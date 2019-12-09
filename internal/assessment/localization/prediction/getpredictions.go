package prediction

import (
	"github.com/knightjdr/cmgo/internal/pkg/read/nmf"
	"github.com/knightjdr/cmgo/internal/pkg/read/safe"
)

func getPredictions(options parameters) map[string]int {
	return readPredictions(options.predictionType, options.predictions)
}

func readPredictions(predictionType, filename string) map[string]int {
	predictedCompartment := make(map[string]int, 0)

	if predictionType == "safe" {
		safePredictions := safe.Localization(filename)
		for prey, prediction := range safePredictions {
			predictedCompartment[prey] = prediction.Compartment
		}
	} else {
		nmfPredictions := nmf.Localization(filename)
		for prey, prediction := range nmfPredictions {
			predictedCompartment[prey] = prediction.Compartment
		}
	}

	return predictedCompartment
}
