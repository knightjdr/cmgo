package gprofiler_test

import (
	"github.com/knightjdr/cmgo/pkg/gprofiler"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Response", func() {
	Describe("AddInteractionGenes", func() {
		It("should add a list of intersecting/matching genes to each enriched term in response", func() {
			response := gprofiler.Response{
				MetaData: gprofiler.MetaData{
					GenesMetaData: gprofiler.GenesMetaData{
						Query: map[string]gprofiler.Query{
							"query_1": gprofiler.Query{
								ENSG: []string{"1", "3", "7"},
								Mapping: map[string][]string{
									"a": []string{"1"},
									"d": []string{"7"},
									"x": []string{"3"},
								},
							},
						},
					},
				},
				Result: []gprofiler.EnrichedTerm{
					{
						Intersections: [][]string{
							{"IEA"},
							{},
							{"IMP", "ISS"},
						},
					},
					{
						Intersections: [][]string{
							{},
							{"IEA"},
							{"IMP", "ISS"},
						},
					},
				},
			}
			response.AddIntersectionGenes("query_1")

			expected := []gprofiler.EnrichedTerm{
				{
					Genes: []string{"a", "d"},
					Intersections: [][]string{
						{"IEA"},
						{},
						{"IMP", "ISS"},
					},
				},
				{
					Genes: []string{"d", "x"},
					Intersections: [][]string{
						{},
						{"IEA"},
						{"IMP", "ISS"},
					},
				},
			}
			Expect(response.Result).To(Equal(expected))
		})
	})
})
