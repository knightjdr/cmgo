package prediction

import (
	"github.com/knightjdr/cmgo/internal/pkg/read/geneontology"
	"github.com/knightjdr/cmgo/internal/pkg/read/localization"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Calculate study score component", func() {
	It("should return score and consistent localizations", func() {
		goHierarchy := &geneontology.GOhierarchy{
			"CC": map[string]*geneontology.GOterm{
				"GO:111111": &geneontology.GOterm{},
				"GO:222222": &geneontology.GOterm{},
				"GO:333333": &geneontology.GOterm{},
				"GO:444444": &geneontology.GOterm{
					Children: []string{"GO:111111"},
				},
				"GO:555555": &geneontology.GOterm{
					Children: []string{"GO:333333"},
				},
			},
		}
		predictions := map[string]int{
			"prey1": 1,
			"prey2": 2,
			"prey3": 1,
			"prey4": 1,
		}
		predictionSummary := localization.Summary{
			1: localization.CompartmentSummary{
				GOid: []string{"GO:111111", "GO:222222"},
			},
			2: localization.CompartmentSummary{
				GOid: []string{"GO:333333"},
			},
		}
		inputfiles := fileContent{
			goHierarchy:       goHierarchy,
			predictions:       predictions,
			predictionSummary: predictionSummary,
		}
		studyPredictions := map[string]map[string][]string{
			"fractionation": map[string][]string{
				"prey1": []string{"GO:111111", "GO:222222"},
				"prey2": []string{"GO:555555"},
				"prey4": []string{"GO:555555"},
			},
			"hpa": map[string][]string{
				"prey1": []string{"GO:444444"},
				"prey2": []string{"GO:222222"},
				"prey3": []string{"GO:111111"},
				"prey4": []string{"GO:333333"},
			},
		}

		expected := preyStudyScore{
			"prey1": &studyScoreComponents{
				fractionation: []string{"GO:111111", "GO:222222"},
				hpa:           []string{"GO:444444"},
				score:         1,
			},
			"prey2": &studyScoreComponents{
				fractionation: []string{"GO:555555"},
				hpa:           []string{},
				score:         0.5,
			},
			"prey3": &studyScoreComponents{
				fractionation: []string{},
				hpa:           []string{"GO:111111"},
				score:         0.5,
			},
			"prey4": &studyScoreComponents{
				fractionation: []string{},
				hpa:           []string{},
				score:         0,
			},
		}

		actual := calculateStudyComponentScore(studyPredictions, inputfiles)
		for prey, scoreComponent := range *actual {
			Expect(scoreComponent.fractionation).To(Equal(expected[prey].fractionation), "should have consistent fractionation predictions")
			Expect(scoreComponent.hpa).To(Equal(expected[prey].hpa), "should have consistent hpa predictions")
			Expect(scoreComponent.score).To(BeNumerically("~", expected[prey].score, 0.00001))
		}
	})
})

var _ = Describe("Get cellmap predictions", func() {
	It("should map compartment predictions to GO IDs", func() {
		predictions := map[string]int{
			"prey1": 1,
			"prey2": 2,
			"prey3": 1,
			"prey4": 1,
		}
		predictionSummary := localization.Summary{
			1: localization.CompartmentSummary{
				GOid: []string{"GO:111111", "GO:222222"},
			},
			2: localization.CompartmentSummary{
				GOid: []string{"GO:333333"},
			},
		}

		expected := map[string][]string{
			"prey1": []string{"GO:111111", "GO:222222"},
			"prey2": []string{"GO:333333"},
			"prey3": []string{"GO:111111", "GO:222222"},
			"prey4": []string{"GO:111111", "GO:222222"},
		}

		Expect(getCellmapPredictions(predictions, predictionSummary)).To(Equal(expected))
	})
})

var _ = Describe("Consistent IDs", func() {
	It("should identifiy shared consistent ID", func() {
		goHierarchy := &geneontology.GOhierarchy{
			"CC": map[string]*geneontology.GOterm{
				"GO:111111": &geneontology.GOterm{},
				"GO:222222": &geneontology.GOterm{},
				"GO:333333": &geneontology.GOterm{},
				"GO:444444": &geneontology.GOterm{
					Children: []string{"GO:111111"},
				},
				"GO:555555": &geneontology.GOterm{
					Children: []string{"GO:333333"},
				},
			},
		}

		prey := "prey1"
		cellmapIDs := []string{"GO:111111", "GO:222222"}
		studyPredictions := map[string][]string{
			"prey1": []string{"GO:333333", "GO:444444"},
		}

		expected := []string{"GO:444444"}
		Expect(getConsitentIDs(prey, cellmapIDs, studyPredictions, goHierarchy)).To(Equal(expected))
	})

	It("should identifiy shared consistent IDs (multiple)", func() {
		goHierarchy := &geneontology.GOhierarchy{
			"CC": map[string]*geneontology.GOterm{
				"GO:111111": &geneontology.GOterm{},
				"GO:222222": &geneontology.GOterm{},
				"GO:333333": &geneontology.GOterm{},
				"GO:444444": &geneontology.GOterm{
					Children: []string{"GO:111111"},
				},
				"GO:555555": &geneontology.GOterm{
					Children: []string{"GO:333333"},
				},
			},
		}

		prey := "prey1"
		cellmapIDs := []string{"GO:111111", "GO:333333"}
		studyPredictions := map[string][]string{
			"prey1": []string{"GO:444444", "GO:555555"},
		}

		expected := []string{"GO:444444", "GO:555555"}
		Expect(getConsitentIDs(prey, cellmapIDs, studyPredictions, goHierarchy)).To(Equal(expected))
	})

	It("should return empty slice for missing prey prediction in other study", func() {
		goHierarchy := &geneontology.GOhierarchy{
			"CC": map[string]*geneontology.GOterm{
				"GO:111111": &geneontology.GOterm{},
				"GO:222222": &geneontology.GOterm{},
				"GO:333333": &geneontology.GOterm{},
				"GO:444444": &geneontology.GOterm{
					Children: []string{"GO:111111"},
				},
				"GO:555555": &geneontology.GOterm{
					Children: []string{"GO:333333"},
				},
			},
		}

		prey := "prey1"
		cellmapIDs := []string{"GO:111111", "GO:333333"}
		studyPredictions := map[string][]string{}

		expected := []string{}
		Expect(getConsitentIDs(prey, cellmapIDs, studyPredictions, goHierarchy)).To(Equal(expected))
	})
})
