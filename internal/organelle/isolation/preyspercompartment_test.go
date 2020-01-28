package isolation

import (
	"github.com/knightjdr/cmgo/internal/pkg/read/nmf"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Preys per compartment", func() {
	It("should find preys assigned to each compartment", func() {
		genes := []string{"geneB", "geneA", "geneC"}
		localizations := nmf.NMFlocalization{
			"geneA": nmf.GeneLocalization{
				Compartment: 1,
			},
			"geneB": nmf.GeneLocalization{
				Compartment: 2,
			},
			"geneC": nmf.GeneLocalization{
				Compartment: 1,
			},
		}

		expected := map[int]map[int]bool{
			1: map[int]bool{
				1: true,
				2: true,
			},
			2: map[int]bool{
				0: true,
			},
		}
		Expect(getPreysPerCompartment(localizations, genes)).To(Equal(expected))
	})
})
