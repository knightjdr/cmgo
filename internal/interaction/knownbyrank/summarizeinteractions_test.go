package knownbyrank

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Summarize interactions", func() {
	It("should count known interactions and total baits having a prey at the Nth rank", func() {
		knownInteractions := map[string][]string{
			"geneA": []string{"gene1", "gene3"},
			"geneB": []string{"gene4", "gene5"},
			"geneC": []string{"gene6", "gene7"},
			"gene1": []string{"geneA"},
			"gene3": []string{"geneA"},
			"gene4": []string{"geneB"},
			"gene5": []string{"geneB"},
			"gene6": []string{"geneC"},
			"gene7": []string{"geneC"},
		}
		sortedPreysPerBait := map[string][]string{
			"geneA": []string{"gene1", "gene2", "gene3"},
			"geneB": []string{"gene4", "gene5"},
			"geneC": []string{"gene6", "gene7", "gene8", "gene9"},
		}

		expected := map[int]*rankSummary{
			1: &rankSummary{
				BaitNumber: 3,
				Known:      3,
				Pairs:      []string{"geneA-gene1", "geneB-gene4", "geneC-gene6"},
			},
			2: &rankSummary{
				BaitNumber: 3,
				Known:      2,
				Pairs:      []string{"geneB-gene5", "geneC-gene7"},
			},
			3: &rankSummary{
				BaitNumber: 2,
				Known:      1,
				Pairs:      []string{"geneA-gene3"},
			},
			4: &rankSummary{
				BaitNumber: 1,
				Known:      0,
				Pairs:      []string{},
			},
		}
		Expect(summarizeInteractions(sortedPreysPerBait, knownInteractions)).To(Equal(expected))
	})
})
