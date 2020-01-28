package isolation

import (
	"github.com/knightjdr/cmgo/internal/pkg/read/nmf"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Compartment sharing", func() {
	It("should count edges shared with other compartments", func() {
		genes := []string{"geneA", "geneB", "geneC", "geneD"}
		localizations := nmf.NMFlocalization{
			"geneA": nmf.GeneLocalization{
				Compartment: 1,
			},
			"geneB": nmf.GeneLocalization{
				Compartment: 2,
			},
			"geneC": nmf.GeneLocalization{
				Compartment: 2,
			},
			"geneD": nmf.GeneLocalization{
				Compartment: 1,
			},
		}
		scores := &isolationScores{
			1: &isolationScore{
				edgesOutside:       3,
				edgesWithin:        1,
				isolation:          0.25,
				nodesOutside:       []int{1, 2},
				sharedCompartments: []int{0, 0, 0},
			},
			2: &isolationScore{
				edgesOutside:       3,
				edgesWithin:        0,
				isolation:          0,
				nodesOutside:       []int{0, 3},
				sharedCompartments: []int{0, 0, 0},
			},
		}

		expected := &isolationScores{
			1: &isolationScore{
				edgesOutside:       3,
				edgesWithin:        1,
				isolation:          0.25,
				nodesOutside:       []int{1, 2},
				sharedCompartments: []int{0, 0, 2},
			},
			2: &isolationScore{
				edgesOutside:       3,
				edgesWithin:        0,
				isolation:          0,
				nodesOutside:       []int{0, 3},
				sharedCompartments: []int{0, 2, 0},
			},
		}
		calculateCompartmentSharing(scores, genes, localizations)
		Expect(scores).To(Equal(expected))
	})
})
