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
		}

		expected := map[string]preySummary{
			"preyA": preySummary{
				cytosolicBaits: 3,
				cytosolicScore: 0.3,
				localization:   "cytosolic",
				lumenalBaits:   1,
				lumenalScore:   0.2,
			},
			"preyC": preySummary{
				cytosolicBaits: 2,
				cytosolicScore: 0.15,
				localization:   "lumenal",
				lumenalBaits:   2,
				lumenalScore:   0.6,
			},
		}
		Expect(summarize(summOptions)).To(Equal(expected))
	})
})
