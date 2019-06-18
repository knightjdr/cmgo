package gene

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Dict", func() {
	It("should generate a map from Entrez to Ensembl", func() {
		ids := []map[string]string{
			{
				"EnsemblGene": "ENSG00000121410",
				"Entrez":      "1",
				"HGNC":        "HGNC:5",
				"Name":        "alpha-1-B glycoprotein",
				"Symbol":      "A1BG",
				"UniProt":     "P04217",
			},
			{
				"EnsemblGene": "ENSG00000268895",
				"Entrez":      "503538",
				"HGNC":        "HGNC:37133",
				"Name":        "A1BG antisense RNA 1",
				"Symbol":      "A1BG-AS1",
				"UniProt":     "",
			},
		}
		expected := map[string]string{
			"1":      "ENSG00000121410",
			"503538": "ENSG00000268895",
		}
		Expect(dict(ids, "Entrez", "EnsemblGene")).To(Equal(expected))
	})

	It("should generate a map from UniProt to HGNC", func() {
		ids := []map[string]string{
			{
				"EnsemblGene": "ENSG00000121410",
				"Entrez":      "1",
				"HGNC":        "HGNC:5",
				"Name":        "alpha-1-B glycoprotein",
				"Symbol":      "A1BG",
				"UniProt":     "P04217",
			},
			{
				"EnsemblGene": "ENSG00000268895",
				"Entrez":      "503538",
				"HGNC":        "HGNC:37133",
				"Name":        "A1BG antisense RNA 1",
				"Symbol":      "A1BG-AS1",
				"UniProt":     "",
			},
		}
		expected := map[string]string{
			"P04217": "HGNC:5",
		}
		Expect(dict(ids, "UniProt", "HGNC")).To(Equal(expected))
	})
})
