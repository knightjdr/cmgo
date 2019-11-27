package uniprot

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var responseText = `{
	"accession": "Q9BUL8",
	"id": "PDC10_HUMAN",
	"organism": {
		"taxonomy": 9606,
		"names": [
			{ "type": "scientific", "value": "Homo sapiens" },
			{ "type": "common", "value": "Human" }
		]
	},
	"secondaryAccession": ["A8K515", "D3DNN5", "O14811"],
	"gene": [
		{
			"name": { "value": "PDCD10" },
			"synonyms": [
				{ "value": "CCM3" },
				{ "value": "TFAR15" }
			]
		}
	],
	"features": [
		{
			"type": "CHAIN",
			"category": "MOLECULE_PROCESSING",
			"ftId": "PRO_0000187562",
			"description": "Programmed cell death protein 10",
			"begin": "1",
			"end": "212"
		},
		{
			"type": "MOD_RES",
			"category": "PTM",
			"description": "N6-acetyllysine",
			"begin": "179",
			"end": "179",
			"evidences": []
		}
	],
	"sequence": {
			"version": 1,
			"length": 212,
			"mass": 24702,
			"modified": "2001-06-01",
			"sequence": "MRMTMEEMKNEAETTSMVSMPLYAVMYPVFNELERVNLSAAQTLRAAFIKAEKENPGLTQDIIMKILEKKSVEVNFTESLLRMAADDVEEYMIERPEPEFQDLNEKARALKQILSKIPDEINDRVRFLQTIKDIASAIKELLDTVNNVFKKYQYQNRRALEHQKKEFVKYSKSFSDTLKTYFKDGKAINVFVSANRLIHQTNLILQTFKTVA"
	}
}`

var _ = Describe("Fetch from Uniprot", func() {
	Context("successfully", func() {
		It("should return regions", func() {
			var apiStub = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusCreated)
				w.Write([]byte(responseText))
			}))

			ids := []string{"Q9BUL8"}
			service := &uniprotService{
				URL: apiStub.URL,
			}
			fetchUniprot(service, ids)

			expected := &Entries{
				"Q9BUL8": Entry{
					Accession: "Q9BUL8",
					Features: []Feature{
						Feature{
							Begin:       1,
							Category:    "MOLECULE_PROCESSING",
							Description: "Programmed cell death protein 10",
							End:         212,
							Type:        "CHAIN",
						},
						Feature{
							Begin:       179,
							Category:    "PTM",
							Description: "N6-acetyllysine",
							End:         179,
							Type:        "MOD_RES",
						},
					},
					Gene: Gene{
						Symbol:   "PDCD10",
						Synonyms: []string{"CCM3", "TFAR15"},
					},
					ID: "PDC10_HUMAN",
					Organism: Organism{
						Common: "Human",
						Names: []OrganismName{
							OrganismName{Type: "scientific", Value: "Homo sapiens"},
							OrganismName{Type: "common", Value: "Human"},
						},
						Scientific: "Homo sapiens",
						Taxonomy:   9606,
					},
					SecondaryAccession: []string{"A8K515", "D3DNN5", "O14811"},
					Sequence:           Sequence{Length: 212, Mass: 24702},
				},
			}
			Expect(service.Result).To(Equal(expected))
		})
	})

	Context("missing ID", func() {
		It("should return empty struct", func() {
			ids := []string{""}
			service := &uniprotService{
				URL: "",
			}
			fetchUniprot(service, ids)

			expected := &Entries{}
			Expect(service.Result).To(Equal(expected))
		})
	})
})
