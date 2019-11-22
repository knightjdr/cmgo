package turnoverbyrank

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Summarize interactions", func() {
	It("should count known interactions and total baits having a prey at the Nth rank", func() {
		sortedPreysPerBait := map[string][]string{
			"geneA": []string{"gene1", "gene2", "gene3"},
			"geneB": []string{"gene1", "gene5"},
			"geneC": []string{"gene6", "gene7", "gene8", "gene9"},
		}
		turnoverRates := map[string]float64{
			"gene1": 1.00,
			"gene2": 2.00,
			"gene3": 3.00,
			"gene5": 5.00,
			"gene6": 6.00,
			"gene7": 7.00,
			"gene9": 9.00,
		}

		expected := map[int]*rankSummary{
			1: &rankSummary{
				TurnoverPreys: map[string]bool{"gene1": true, "gene6": true},
				TurnoverRates: []float64{1, 6},
			},
			2: &rankSummary{
				TurnoverPreys: map[string]bool{"gene2": true, "gene5": true, "gene7": true},
				TurnoverRates: []float64{2, 5, 7},
			},
			3: &rankSummary{
				TurnoverPreys: map[string]bool{"gene3": true},
				TurnoverRates: []float64{3},
			},
			4: &rankSummary{
				TurnoverPreys: map[string]bool{"gene9": true},
				TurnoverRates: []float64{9},
			},
		}
		Expect(summarizeInteractions(sortedPreysPerBait, turnoverRates)).To(Equal(expected))
	})
})
