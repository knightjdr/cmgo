package bioplex_test

import (
	"github.com/knightjdr/cmgo/internal/pkg/read/bioplex"
	"github.com/knightjdr/cmgo/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var bioplexText = `GeneA	GeneB	UniprotA	UniprotB	SymbolA	SymbolB	p(Wrong)	p(No Interaction)	p(Interaction)
100	728378	P00813	A5A3E0	ADA	POTEF	2.38085788908859e-9	0.000331855941652957	0.999668141677489
222389	708	Q8N7W2	Q07021	BEND7	C1QBP	2.96221526059856e-17	0.00564451180569955	0.994355488194301
222389	4038	Q8N7W2	O75096	BEND7	LRP4	3.30299445393738e-10	0.000280259555661228	0.999719740114039
645121	3312	Q6ZMN8	P11142	CCNI2	HSPA8	2.06028533960837e-16	0.0362347656743182	0.963765234325682
`

var _ = Describe("Read", func() {
	It("should read BioPlex file as list of interactions", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(
			fs.Instance,
			"test/bioplex.txt",
			[]byte(bioplexText),
			0444,
		)
		expected := []bioplex.Entry{
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
		Expect(bioplex.Read("test/bioplex.txt", true)).To(Equal(expected))
	})
})
