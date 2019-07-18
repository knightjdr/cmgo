package localize

import (
	"github.com/knightjdr/cmgo/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var enrichmentText = `symbol	Entrez	Refseq	UniProt	GO ID	GO term	p-value	recall	precision	query size	term size
RBM15	64783	NP_001188474.1	Q96T37	GO:0044428	nuclear part	4.243467e-19	0.01	0.88	50	4521
RBM15	64783	NP_001188474.1	Q96T37	GO:0031981	nuclear lumen	4.508720e-18	0.01	0.84	50	4142
RBM15	64783	NP_001188474.1	Q96T37	GO:0070013	intracellular organelle lumen	6.906764e-14	0.01	0.84	50	5283
RBM15	64783	NP_001188474.1	Q96T37	GO:0043233	organelle lumen	6.906764e-14	0.01	0.84	50	5283
CSNK1G3	1456	NP_001026982.1	Q9Y6M4	GO:0005886	plasma membrane	1.838539e-19	0.01	0.94	50	5539
CSNK1G3	1456	NP_001026982.1	Q9Y6M4	GO:0071944	cell periphery	5.044889e-19	0.01	0.94	50	5662
`

var _ = Describe("Read enrichment data", func() {
	It("should read enrichment data", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(
			fs.Instance,
			"test/enrichment.txt",
			[]byte(enrichmentText),
			0444,
		)
		expected := map[string][]Enrichment{
			"RBM15": []Enrichment{
				{Entrez: "64783", ID: "GO:0044428", Precision: 0.88, Pvalue: 4.243467e-19},
				{Entrez: "64783", ID: "GO:0031981", Precision: 0.84, Pvalue: 4.508720e-18},
				{Entrez: "64783", ID: "GO:0070013", Precision: 0.84, Pvalue: 6.906764e-14},
				{Entrez: "64783", ID: "GO:0043233", Precision: 0.84, Pvalue: 6.906764e-14},
			},
			"CSNK1G3": []Enrichment{
				{Entrez: "1456", ID: "GO:0005886", Precision: 0.94, Pvalue: 1.838539e-19},
				{Entrez: "1456", ID: "GO:0071944", Precision: 0.94, Pvalue: 5.044889e-19},
			},
		}
		Expect(readEnrichment("test/enrichment.txt")).To(Equal(expected))
	})
})
