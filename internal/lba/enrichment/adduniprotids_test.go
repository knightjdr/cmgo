package enrichment

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Add Uniprot IDs", func() {
	It("should add uniprot IDs to Refseq mapping", func() {
		refseqMapping := map[string]map[string]string{
			"NP_001263222.1": map[string]string{
				"Entrez":  "11188",
				"Symbol":  "NISCH",
				"UniProt": "",
			},
			"NP_001263218.1": map[string]string{
				"Entrez":  "5573",
				"Symbol":  "PRKAR1A",
				"UniProt": "",
			},
			"NP_00XXXXX.2": map[string]string{
				"Entrez":  "1234",
				"Symbol":  "TESTGENE",
				"UniProt": "",
			},
		}
		entrezToUniprot := map[string]string{
			"11188": "a",
			"5573":  "b",
		}
		expected := map[string]map[string]string{
			"NP_001263222.1": map[string]string{
				"Entrez":  "11188",
				"Symbol":  "NISCH",
				"UniProt": "a",
			},
			"NP_001263218.1": map[string]string{
				"Entrez":  "5573",
				"Symbol":  "PRKAR1A",
				"UniProt": "b",
			},
			"NP_00XXXXX.2": map[string]string{
				"Entrez":  "1234",
				"Symbol":  "TESTGENE",
				"UniProt": "",
			},
		}
		addUniprotIDs(&refseqMapping, entrezToUniprot)
		Expect(refseqMapping).To(Equal(expected))
	})
})
