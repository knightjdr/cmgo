package prediction

import (
	"github.com/knightjdr/cmgo/internal/pkg/read/localization"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Calculate text score component", func() {
	It("should return GO ID and score", func() {
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
			predictions:       predictions,
			predictionSummary: predictionSummary,
		}
		textAnnotations := map[string]map[string]float64{
			"prey1": map[string]float64{
				"GO:111111": 3.6,
				"GO:222222": 3.2,
			},
			"prey2": map[string]float64{
				"GO:111111": 3.8,
				"GO:333333": 3.4,
			},
			"prey3": map[string]float64{
				"GO:444444": 3.9,
			},
			"prey4": map[string]float64{
				"GO:111111": 3.1,
				"GO:333333": 3.3,
			},
		}

		expected := preyTextScore{
			"prey1": &textScoreComponents{
				GOID:  "GO:111111",
				score: 0.9,
			},
			"prey2": &textScoreComponents{
				GOID:  "GO:333333",
				score: 0.85,
			},
			"prey3": &textScoreComponents{
				GOID:  "",
				score: 0,
			},
			"prey4": &textScoreComponents{
				GOID:  "GO:111111",
				score: 0.775,
			},
		}

		actual := calculateTextScoreComponent(textAnnotations, inputfiles)
		for prey, scoreComponent := range *actual {
			Expect(scoreComponent.GOID).To(Equal(expected[prey].GOID), "should have best GO ID score")
			Expect(scoreComponent.score).To(BeNumerically("~", expected[prey].score, 0.00001))
		}
	})
})

var _ = Describe("Find best text annotation", func() {
	It("should return nil GO ID and 0 for score when there are no text annotations", func() {
		annotations := map[string]map[string]float64{}
		cellmapIDs := []string{"GO:111111", "GO:222222"}

		expectedGOID := ""
		expectedScore := float64(0)

		actualGOID, actualScore := findBestTextAnnotation(cellmapIDs, annotations["gene"])

		Expect(actualGOID).To(Equal(expectedGOID))
		Expect(actualScore).To(Equal(expectedScore))
	})

	It("should return GO ID and max score when there are matching text annotations", func() {
		annotations := map[string]map[string]float64{
			"gene": map[string]float64{
				"GO:111111": 3.4,
				"GO:222222": 3.8,
				"GO:333333": 2.5,
				"GO:444444": 4.5,
			},
		}
		cellmapIDs := []string{"GO:111111", "GO:222222", "GO:333333"}

		expectedGOID := "GO:222222"
		expectedScore := 3.8

		actualGOID, actualScore := findBestTextAnnotation(cellmapIDs, annotations["gene"])

		Expect(actualGOID).To(Equal(expectedGOID))
		Expect(actualScore).To(Equal(expectedScore))
	})

	It("should return nil GO ID and 0 for score when there are no matching text annotations", func() {
		annotations := map[string]map[string]float64{
			"gene": map[string]float64{
				"GO:111111": 3.4,
				"GO:222222": 3.8,
				"GO:333333": 2.5,
				"GO:444444": 4.5,
			},
		}
		cellmapIDs := []string{"GO:555555", "GO:666666"}

		expectedGOID := ""
		expectedScore := float64(0)

		actualGOID, actualScore := findBestTextAnnotation(cellmapIDs, annotations["gene"])

		Expect(actualGOID).To(Equal(expectedGOID))
		Expect(actualScore).To(Equal(expectedScore))
	})
})
