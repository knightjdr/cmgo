package uniprot_test

import (
	"github.com/knightjdr/cmgo/internal/pkg/read/uniprot"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Create map from gene name to ID", func() {
	It("should return entrys", func() {
		uniprotData := &uniprot.Entries{
			"P31946": uniprot.Entry{
				Symbol: "YWHAB",
			},
			"P62257": uniprot.Entry{
				Symbol: "",
			},
			"P62258": uniprot.Entry{
				Symbol: "YWHAE",
			},
		}

		expected := map[string]string{
			"YWHAB": "P31946",
			"YWHAE": "P62258",
		}
		Expect(uniprotData.CreateGeneNameMap()).To(Equal(expected))
	})
})
