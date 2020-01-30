package rankmetrics

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Summarize metrics", func() {
	It("should calculate metrics for preys at the Nth rank", func() {
		lysines := map[string]int{
			"gene1": 11,
			"gene2": 12,
			"gene3": 13,
			"gene6": 16,
			"gene7": 17,
			"gene9": 19,
		}
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
				LysinePreys:   map[string]bool{"gene1": true, "gene6": true},
				Lysines:       []int{11, 16},
				TurnoverPreys: map[string]bool{"gene1": true, "gene6": true},
				TurnoverRates: []float64{1, 6},
			},
			2: &rankSummary{
				LysinePreys:   map[string]bool{"gene2": true, "gene7": true},
				Lysines:       []int{12, 17},
				TurnoverPreys: map[string]bool{"gene2": true, "gene5": true, "gene7": true},
				TurnoverRates: []float64{2, 5, 7},
			},
			3: &rankSummary{
				LysinePreys:   map[string]bool{"gene3": true},
				Lysines:       []int{13},
				TurnoverPreys: map[string]bool{"gene3": true},
				TurnoverRates: []float64{3},
			},
			4: &rankSummary{
				LysinePreys:   map[string]bool{"gene9": true},
				Lysines:       []int{19},
				TurnoverPreys: map[string]bool{"gene9": true},
				TurnoverRates: []float64{9},
			},
		}
		Expect(summarizeMetrics(sortedPreysPerBait, lysines, turnoverRates)).To(Equal(expected))
	})
})
