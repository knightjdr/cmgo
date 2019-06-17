package robustness

import (
	"github.com/knightjdr/cmgo/pkg/gprofiler"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DefineRanks", func() {
	It("", func() {
		originalFetch := fetch
		defer func() { fetch = originalFetch }()
		fetch = func(s *gprofiler.Service) {
			s.Result = []gprofiler.EnrichedTerm{
				{ID: "GO:1"},
				{ID: "GO:2"},
				{ID: "GO:3"},
			}
		}

		characterizingGenes := [][]int{
			{0, 1},
			{2, 3},
		}
		geneNames := []string{"a", "b", "c", "d"}
		actual := defineRanks(characterizingGenes, geneNames)
		expected := [][]string{
			{"GO:1", "GO:2", "GO:3"},
			{"GO:1", "GO:2", "GO:3"},
		}
		Expect(actual).To(Equal(expected), "should return GO definitions for each rank")
	})
})
