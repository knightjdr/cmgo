package prediction

import (
	"sort"

	"github.com/knightjdr/cmgo/internal/pkg/read/geneontology"
	"github.com/knightjdr/cmgo/internal/pkg/read/localization"
	"github.com/knightjdr/cmgo/pkg/math"
	"github.com/knightjdr/cmgo/pkg/slice"
)

type preyStudyScore map[string]*studyScoreComponents

type studyScoreComponents struct {
	fractionation []string
	hpa           []string
	score         float64
}

func calculateStudyComponentScore(studyPredictions map[string]map[string][]string, inputfile fileContent) *preyStudyScore {
	scores := &preyStudyScore{}

	cellmapPredictions := getCellmapPredictions(inputfile.predictions, inputfile.predictionSummary)
	for prey, cellmapGoIDs := range cellmapPredictions {
		(*scores)[prey] = &studyScoreComponents{
			fractionation: getConsitentIDs(prey, cellmapGoIDs, studyPredictions["fractionation"], inputfile.goHierarchy),
			hpa:           getConsitentIDs(prey, cellmapGoIDs, studyPredictions["hpa"], inputfile.goHierarchy),
		}
		(*scores)[prey].score = calculateFinalStudyScore((*scores)[prey])
	}

	return scores
}

func getCellmapPredictions(predictions map[string]int, predictionSummary localization.Summary) map[string][]string {
	predictionsAsGoID := make(map[string][]string, 0)

	for gene, compartment := range predictions {
		predictionsAsGoID[gene] = predictionSummary[compartment].GOid
	}

	return predictionsAsGoID
}

func getConsitentIDs(prey string, cellmapIDs []string, studyPredictions map[string][]string, goHierarchy *geneontology.GOhierarchy) []string {
	consistentIDs := make([]string, 0)

	if _, ok := studyPredictions[prey]; ok {
		for _, cellmapID := range cellmapIDs {
			for _, studyID := range studyPredictions[prey] {
				if cellmapID != "" && goHierarchy.AreConsistent("CC", cellmapID, studyID) {
					consistentIDs = append(consistentIDs, studyID)
				}
			}
		}
	}

	consistentIDs = slice.UniqueStrings(consistentIDs)
	sort.Strings(consistentIDs)

	return consistentIDs
}

func calculateFinalStudyScore(scoreComponents *studyScoreComponents) float64 {
	score := float64(0)
	if len(scoreComponents.fractionation) > 0 {
		score++
	}
	if len(scoreComponents.hpa) > 0 {
		score++
	}

	return math.Round(score/float64(2), 0.00001)
}
