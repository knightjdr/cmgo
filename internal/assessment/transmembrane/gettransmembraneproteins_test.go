package transmembrane

import (
	"sort"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/cmgo/pkg/pfam"
)

var _ = Describe("Filter transmembrane proteins", func() {
	It("should return a slice of proteins that have a transmembrane motif", func() {
		features := &pfam.Features{
			"id1": pfam.Regions{
				Motifs: []pfam.Motif{
					{Name: "disorder"},
				},
			},
			"id2": pfam.Regions{
				Motifs: []pfam.Motif{
					{Name: "disorder"},
					{Name: "transmembrane"},
					{Name: "disorder"},
				},
			},
			"id3": pfam.Regions{
				Motifs: []pfam.Motif{
					{Name: "transmembrane"},
				},
			},
			"id4": pfam.Regions{
				Motifs: []pfam.Motif{
					{Name: "disorder"},
				},
			},
		}

		actual := filterTransmembrane(features)
		sort.Strings(actual)
		expected := []string{"id2", "id3"}
		Expect(actual).To(Equal(expected))
	})
})

var _ = Describe("Map UniProtID to gene Symbol", func() {
	It("should return a slice gene symbols", func() {
		symbolToUniProt := map[string]string{
			"symbolA": "id2",
			"symbolB": "id1",
			"symbolC": "id3",
			"symbolD": "id4",
		}
		uniprotIDs := []string{"id1", "id2", "id3"}

		expected := []string{"symbolB", "symbolA", "symbolC"}
		Expect(mapUniprotToSymbol(symbolToUniProt, uniprotIDs)).To(Equal(expected))
	})
})
