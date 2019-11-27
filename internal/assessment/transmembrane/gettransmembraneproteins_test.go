package transmembrane

import (
	"sort"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/cmgo/pkg/pfam"
	"github.com/knightjdr/cmgo/pkg/uniprot"
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

var _ = Describe("Extract orientation", func() {
	It("should extract orientation", func() {
		entries := &uniprot.Entries{
			"id1": uniprot.Entry{
				Features: []uniprot.Feature{
					uniprot.Feature{
						Begin:       1,
						Description: "Extracellular",
						End:         25,
					},
					uniprot.Feature{
						Begin:       26,
						Description: "Transmembrane",
						End:         36,
					},
					uniprot.Feature{
						Begin:       37,
						Description: "Cytoplasmic",
						End:         100,
					},
				},
				Sequence: uniprot.Sequence{Length: 100},
			},
			"id2": uniprot.Entry{
				Features: []uniprot.Feature{
					uniprot.Feature{
						Begin:       1,
						Description: "PTM",
						End:         1,
					},
					uniprot.Feature{
						Begin:       37,
						Description: "Extracellular",
						End:         85,
					},
					uniprot.Feature{
						Begin:       86,
						Description: "Transmembrane",
						End:         95,
					},
					uniprot.Feature{
						Begin:       95,
						Description: "Cytoplasmic",
						End:         150,
					},
					uniprot.Feature{
						Begin:       151,
						Description: "Transmembrane",
						End:         171,
					},
					uniprot.Feature{
						Begin:       172,
						Description: "Lumenal",
						End:         200,
					},
				},
				Sequence: uniprot.Sequence{Length: 200},
			},
		}

		expected := map[string]orientationData{
			"id1": orientationData{
				Cytosolic: 64,
				Length:    100,
				Lumenal:   25,
				UniProt:   "id1",
			},
			"id2": orientationData{
				Cytosolic: 56,
				Length:    200,
				Lumenal:   78,
				UniProt:   "id2",
			},
		}
		Expect(extractOrientation(entries)).To(Equal(expected))
	})
})

var _ = Describe("Map protein data to gene syumbol", func() {
	It("should return a slice gene symbols", func() {
		orientation := map[string]orientationData{
			"id1": orientationData{
				Cytosolic: 25,
				Length:    100,
				Lumenal:   65,
				UniProt:   "id1",
			},
			"id3": orientationData{
				Cytosolic: 75,
				Length:    200,
				Lumenal:   10,
				UniProt:   "id3",
			},
		}
		symbolToUniProt := map[string]string{
			"symbolA": "id2",
			"symbolB": "id1",
			"symbolC": "id3",
			"symbolD": "id4",
		}

		expectedData := map[string]orientationData{
			"symbolB": orientationData{
				Cytosolic: 25,
				Length:    100,
				Lumenal:   65,
				UniProt:   "id1",
			},
			"symbolC": orientationData{
				Cytosolic: 75,
				Length:    200,
				Lumenal:   10,
				UniProt:   "id3",
			},
		}
		expectedGenes := []string{"symbolB", "symbolC"}
		acutalGenes, actualData := mapProteinToGene(symbolToUniProt, orientation)
		sort.Strings(acutalGenes)
		Expect(actualData).To(Equal(expectedData), "should return membrane orientation data")
		Expect(acutalGenes).To(Equal(expectedGenes), "should return transmembrane genes")
	})
})
