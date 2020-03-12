package transmembrane

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Summarize prey information", func() {
	It("should return a summary for each prey", func() {
		basis := [][]float64{
			{0.1, 0.2, 0.3},
			{0.4, 0.45, 0.35},
			{0.05, 0.6, 0.15},
		}
		summOptions := summaryOptions{
			basis:                 basis,
			cytosolicCompartments: []string{"1", "3"},
			cytosolicPreys:        []string{"preyA", "preyB"},
			lumenalCompartments:   []string{"2"},
			lumenalPreys:          []string{"preyC"},
			organelleBaitsPerPrey: map[string]map[string]int{
				"preyA": map[string]int{
					"cytosolic": 3,
					"lumenal":   1,
				},
				"preyC": map[string]int{
					"cytosolic": 2,
					"lumenal":   2,
				},
			},
			rows:               []string{"preyA", "preyB", "preyC"},
			transmembranePreys: []string{"preyA", "preyC"},
			transmembranePreyData: map[string]orientationData{
				"preyA": orientationData{
					Cytosolic: 64,
					Length:    100,
					Lumenal:   25,
					UniProt:   "id1",
				},
				"preyC": orientationData{
					Cytosolic: 78,
					Length:    200,
					Lumenal:   56,
					UniProt:   "id2",
				},
			},
		}

		expected := map[string]preySummary{
			"preyA": preySummary{
				cytosolicBaits:    3,
				cytosolicFraction: 0.6400,
				cytosolicScore:    0.3,
				length:            100,
				localization:      "cytosolic",
				lumenalBaits:      1,
				lumenalFraction:   0.2500,
				lumenalScore:      0.2,
				maxCytosolicScore: 0.35,
				maxLumenalScore:   0.6,
				uniprotID:         "id1",
			},
			"preyC": preySummary{
				cytosolicBaits:    2,
				cytosolicFraction: 0.3900,
				cytosolicScore:    0.15,
				length:            200,
				localization:      "lumenal",
				lumenalBaits:      2,
				lumenalFraction:   0.2800,
				lumenalScore:      0.6,
				maxCytosolicScore: 0.35,
				maxLumenalScore:   0.6,
				uniprotID:         "id2",
			},
		}
		Expect(summarize(summOptions)).To(Equal(expected))
	})
})
