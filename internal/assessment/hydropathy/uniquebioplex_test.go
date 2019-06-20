package hydropathy

import (
	"sort"

	"github.com/knightjdr/cmgo/internal/pkg/read/bioplex"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Unique BioPlex interactors", func() {
	It("should return a slice of unique preys from SAINT data", func() {
		bioplexData := []bioplex.Entry{
			{
				Entrez:  bioplex.Interaction{Source: "100", Target: "728378"},
				Symbol:  bioplex.Interaction{Source: "ADA", Target: "POTEF"},
				UniProt: bioplex.Interaction{Source: "P00813", Target: "A5A3E0"},
			},
			{
				Entrez:  bioplex.Interaction{Source: "222389", Target: "708"},
				Symbol:  bioplex.Interaction{Source: "BEND7", Target: "C1QBP"},
				UniProt: bioplex.Interaction{Source: "Q8N7W2", Target: "Q07021"},
			},
			{
				Entrez:  bioplex.Interaction{Source: "222389", Target: "4038"},
				Symbol:  bioplex.Interaction{Source: "BEND7", Target: "LRP4"},
				UniProt: bioplex.Interaction{Source: "Q8N7W2", Target: "O75096"},
			},
			{
				Entrez:  bioplex.Interaction{Source: "645121", Target: "3312"},
				Symbol:  bioplex.Interaction{Source: "CCNI2", Target: "HSPA8"},
				UniProt: bioplex.Interaction{Source: "Q6ZMN8", Target: "P11142"},
			},
		}
		actual := uniqueBioplex(bioplexData)
		sort.Strings(actual)
		expected := []string{"100", "222389", "3312", "4038", "645121", "708", "728378"}
		Expect(actual).To(Equal(expected))
	})
})
