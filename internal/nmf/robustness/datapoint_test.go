package robustness

import (
	"github.com/knightjdr/cmgo/pkg/gprofiler"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DataPoint", func() {
	It("", func() {
		originalFetch := fetch
		defer func() { fetch = originalFetch }()
		fetch = func(s *gprofiler.Service) {
			s.Result = []gprofiler.EnrichedTerm{
				{ID: "GO:a"},
				{ID: "GO:b"},
			}
		}

		geneIndices := []int{0, 1, 2, 3, 4}
		geneNames := []string{"v", "w", "x", "y", "z"}
		rankDefinition := []string{"GO:a", "GO:b", "GO:c", "GO:d", "GO:e"}
		percentiles := []float64{0.8, 0.6}
		persistence := 0.9
		replicates := 2

		actual := dataPoint(geneIndices, geneNames, rankDefinition, percentiles, persistence, replicates)
		expected := [][]float64{
			{0, 0},
			{0, 0},
		}
		Expect(actual).To(Equal(expected), "should return RBD data points for each percentile and with replicates")
	})
})
