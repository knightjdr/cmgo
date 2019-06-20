package hydropathy

import (
	"github.com/knightjdr/cmgo/internal/pkg/read/database"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sequence mapping", func() {
	It("should return mapping of Entrez and refseq to sequence", func() {
		data := []database.Fasta{
			{Entrez: "a", Refseq: "nm_a", Sequence: "abc"},
			{Entrez: "b", Refseq: "nm_b", Sequence: "bcd"},
			{Entrez: "c", Refseq: "nm_c", Sequence: "cde"},
		}
		actualEntrezMapping, actualRefseqMapping, actualRefseqEntez := sequenceMapping(data)
		expectedEntrezMapping := map[string]string{
			"a": "abc",
			"b": "bcd",
			"c": "cde",
		}
		expectedRefseqMapping := map[string]string{
			"nm_a": "abc",
			"nm_b": "bcd",
			"nm_c": "cde",
		}
		expectedRefseqEntez := map[string]string{
			"nm_a": "a",
			"nm_b": "b",
			"nm_c": "c",
		}
		Expect(actualEntrezMapping).To(Equal(expectedEntrezMapping), "should create mapping of Entrez to sequence")
		Expect(actualRefseqMapping).To(Equal(expectedRefseqMapping), "should create mapping of Refseq to sequence")
		Expect(actualRefseqEntez).To(Equal(expectedRefseqEntez), "should create mapping of Refseq to Entrez")
	})
})
