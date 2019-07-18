package localize

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Filter enrichment data", func() {
	It("should remove invalid localizations", func() {
		enrichment := map[string][]Enrichment{
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
		goIDs := map[string]string{
			"GO:0031981": "a",
			"GO:0070013": "b",
			"GO:0071944": "c",
		}

		expected := map[string][]Enrichment{
			"RBM15": []Enrichment{
				{Entrez: "64783", ID: "GO:0031981", Precision: 0.84, Pvalue: 4.508720e-18},
				{Entrez: "64783", ID: "GO:0070013", Precision: 0.84, Pvalue: 6.906764e-14},
			},
			"CSNK1G3": []Enrichment{
				{Entrez: "1456", ID: "GO:0071944", Precision: 0.94, Pvalue: 5.044889e-19},
			},
		}
		
		Expect(filterEnrichment(enrichment, goIDs)).To(Equal(expected))
	})
})
