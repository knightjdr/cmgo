package gene_test

import (
	"github.com/knightjdr/cmgo/pkg/gene"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Response", func() {
	Describe("Define UniProt", func() {
		It("should add primary UniProt ID to gene struct", func() {
			result := gene.Response{
				Docs: []gene.ID{
					{
						EnsemblGene: "ENSG00000121410",
						Entrez:      "1",
						HGNC:        "HGNC:5",
						Name:        "alpha-1-B glycoprotein",
						Symbol:      "A1BG",
						UniProt:     "",
						UniProtList: []string{"P04217"},
					},
					{
						EnsemblGene: "ENSG00000268895",
						Entrez:      "503538",
						HGNC:        "HGNC:37133",
						Name:        "A1BG antisense RNA 1",
						Symbol:      "A1BG-AS1",
						UniProt:     "",
						UniProtList: nil,
					},
				},
			}
			result.DefineUniProt()

			expected := []gene.ID{
				{
					EnsemblGene: "ENSG00000121410",
					Entrez:      "1",
					HGNC:        "HGNC:5",
					Name:        "alpha-1-B glycoprotein",
					Symbol:      "A1BG",
					UniProt:     "P04217",
					UniProtList: []string{"P04217"},
				},
				{
					EnsemblGene: "ENSG00000268895",
					Entrez:      "503538",
					HGNC:        "HGNC:37133",
					Name:        "A1BG antisense RNA 1",
					Symbol:      "A1BG-AS1",
					UniProt:     "",
					UniProtList: nil,
				},
			}
			Expect(result.Docs).To(Equal(expected))
		})
	})

	Describe("ParseIDtoMap", func() {
		It("should return a map of strings", func() {
			result := gene.Response{
				Docs: []gene.ID{
					{
						EnsemblGene: "ENSG00000121410",
						Entrez:      "1",
						HGNC:        "HGNC:5",
						Name:        "alpha-1-B glycoprotein",
						Symbol:      "A1BG",
						UniProt:     "P04217",
					},
					{
						EnsemblGene: "ENSG00000268895",
						Entrez:      "503538",
						HGNC:        "HGNC:37133",
						Name:        "A1BG antisense RNA 1",
						Symbol:      "A1BG-AS1",
						UniProt:     "",
					},
				},
			}

			expected := []map[string]string{
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
			Expect(result.ParseIDtoMap()).To(Equal(expected))
		})
	})
})
