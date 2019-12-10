package prediction

import (
	"sort"

	"github.com/knightjdr/cmgo/pkg/math"
	"github.com/knightjdr/cmgo/pkg/slice"
)

type preyBaitScore map[string]*baitScoreComponents

type baitScoreComponents struct {
	baits []string
	score float64
}

func calculateBaitScoreComponent(predictions map[string]int, baitCompartments baitInformation, baitInteractors map[string][]string) *preyBaitScore {
	preyScores := &preyBaitScore{}

	for prey, prediction := range predictions {
		(*preyScores)[prey] = &baitScoreComponents{
			baits: make([]string, 0),
		}
		for bait, interactors := range baitInteractors {
			if slice.ContainsString(prey, interactors) && slice.ContainsInt(prediction, baitCompartments.localizations[bait]) {
				(*preyScores)[prey].baits = append((*preyScores)[prey].baits, bait)
				(*preyScores)[prey].score++
			}
		}
		sort.Strings((*preyScores)[prey].baits)
		(*preyScores)[prey].score = computeFinalBaitComponentScore((*preyScores)[prey].score, baitCompartments.compartmentCounts[prediction])
	}

	return preyScores
}

func computeFinalBaitComponentScore(baitCount float64, compartmentCount int) float64 {
	if compartmentCount > 0 {
		return math.Round(baitCount/float64(compartmentCount), 0.00001)
	}
	return 0
}
