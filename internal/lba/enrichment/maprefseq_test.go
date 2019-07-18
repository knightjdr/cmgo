package enrichment

import (
	"github.com/knightjdr/cmgo/internal/pkg/read/database"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Map Refseq", func() {
	It("should create mapping from Refseq IDs", func() {
		databaseData := []database.Fasta{
			{
				Entrez:   "11188",
				Refseq:   "NP_001263222.1",
				Sequence: "",
				Symbol:   "NISCH",
			},
			{
				Entrez:   "5573",
				Refseq:   "NP_001263218.1",
				Sequence: "",
				Symbol:   "PRKAR1A",
			},
			{
				Entrez:   "1234",
				Refseq:   "NP_00XXXXX.2",
				Sequence: "",
				Symbol:   "TESTGENE",
			},
		}
		preys := []string{"NP_001263222.1", "NP_00XXXXX.2"}
		actualMapping, actualGeneIDs := mapRefseq(databaseData, preys)
		expectedGeneIDs := []string{"11188", "1234"}
		expectedMapping := map[string]map[string]string{
			"NP_001263222.1": map[string]string{
				"Entrez":  "11188",
				"Symbol":  "NISCH",
				"UniProt": "",
			},
			"NP_00XXXXX.2": map[string]string{
				"Entrez":  "1234",
				"Symbol":  "TESTGENE",
				"UniProt": "",
			},
		}
		Expect(actualGeneIDs).To(Equal(expectedGeneIDs), "Should return a list of gene IDs")
		Expect(actualMapping).To(Equal(expectedMapping), "Should return mapping from Refseq to other IDs")
	})
})
