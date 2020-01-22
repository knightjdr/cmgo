package prediction

import (
	"sort"

	"github.com/knightjdr/cmgo/pkg/math"
	"github.com/knightjdr/cmgo/pkg/slice"
)

type preyBaitScore map[string]*baitScoreComponents

type baitScoreComponents struct {
	interactorBaits  []string
	organelleBaits   []string
	scoreOrganelle   float64
	scoreSpecificity float64
}

func calculateBaitScoreComponent(inputFiles fileContent, baitCompartments baitInformation, baitsPerPrey map[string][]string) *preyBaitScore {
	preyScores := &preyBaitScore{}

	for prey, prediction := range inputFiles.predictions {
		(*preyScores)[prey] = &baitScoreComponents{
			interactorBaits: baitsPerPrey[prey],
			organelleBaits:  make([]string, 0),
		}
		for bait, interactors := range inputFiles.baitInteractors {
			if slice.ContainsString(prey, interactors) && slice.ContainsInt(prediction, baitCompartments.localizations[bait]) {
				(*preyScores)[prey].organelleBaits = append((*preyScores)[prey].organelleBaits, bait)
				(*preyScores)[prey].scoreOrganelle++
			}
		}
		sort.Strings((*preyScores)[prey].organelleBaits)
		(*preyScores)[prey].scoreOrganelle = computeFinalBaitComponentOrganelleScore((*preyScores)[prey].scoreOrganelle, baitCompartments.compartmentCounts[prediction])
		(*preyScores)[prey].scoreSpecificity = computeFinalBaitComponentSpecificityScore((*preyScores)[prey].organelleBaits, (*preyScores)[prey].interactorBaits)
	}

	return preyScores
}

func computeFinalBaitComponentOrganelleScore(baitCount float64, compartmentCount int) float64 {
	if compartmentCount > 0 {
		return math.Round(baitCount/float64(compartmentCount), 0.00001)
	}
	return 0
}

func computeFinalBaitComponentSpecificityScore(organelleBaits, interactorBaits []string) float64 {
	return math.Round(float64(len(organelleBaits))/float64(len(interactorBaits)), 0.00001)
}
