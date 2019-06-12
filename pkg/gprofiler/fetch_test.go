package gprofiler_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/knightjdr/cmgo/pkg/gprofiler"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var responseText = `{
  "result":[
    {
      "goshv":17431,
      "p_value":6.045452804836617e-17,
      "significant":true,
      "effective_domain_size":17710,
      "intersection_size":5,
      "term_size":5,
      "query_size":5,
      "precision":1.0,
      "recall":1.0,
      "name":"regulation of skeletal muscle contraction by regulation of release of sequestered calcium ion",
      "native":"GO:0014809",
      "source":"GO:BP",
      "intersections":[
        [
          "IBA"
        ],
        [
          "ISS",
          "IBA",
          "IEA"
        ],
        [
          "IC"
        ],
        [
          "ISS",
          "IEA"
        ],
        [
          "IC"
        ]
      ],
      "query":"query_1",
      "source_order":5299
    },
    {
      "goshv":17193,
      "p_value":7.823665557395747e-09,
      "significant":true,
      "effective_domain_size":17710,
      "intersection_size":4,
      "term_size":22,
      "query_size":5,
      "precision":0.8,
      "recall":0.18181818181818182,
      "name":"regulation of cardiac muscle contraction by regulation of the release of sequestered calcium ion",
      "native":"GO:0010881",
      "source":"GO:BP",
      "intersections":[
        [
          "IMP",
          "IEA"
        ],
        [],
        [
          "IC"
        ],
        [
          "ISS",
          "IEA"
        ],
        [
          "IC"
        ]
      ],
      "query":"query_1",
      "source_order":5061
    }
  ],
  "meta":{
    "genes_metadata":{
      "query":{
        "query_1":{
          "mapping":{
            "CASQ2":[
              "ENSG00000118729"
            ],
            "CASQ1":[
              "ENSG00000143318"
            ],
            "GSTO1":[
              "ENSG00000148834"
            ],
            "DMD":[
              "ENSG00000198947"
            ],
            "GSTM2":[
              "ENSG00000213366"
            ]
          },
          "ensgs":[
            "ENSG00000118729",
            "ENSG00000143318",
            "ENSG00000148834",
            "ENSG00000198947",
            "ENSG00000213366"
          ]
        }
      }
    }
  }
}`

var _ = Describe("Fetch from g:Profiler", func() {
	Context("successfully", func() {
		It("should return enriched terms for query genes", func() {
			var apiStub = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusCreated)
				w.Write([]byte(responseText))
			}))

			service := &gprofiler.Service{
				URL: apiStub.URL,
			}
			service.Body.Query = []string{"CASQ2", "CASQ1", "GSTO1", "DMD", "GSTM2"}
			terms := service.Fetch()

			expected := []gprofiler.EnrichedTerm{
				{
					Genes:            []string{"CASQ1", "CASQ2", "DMD", "GSTM2", "GSTO1"},
					ID:               "GO:0014809",
					Intersections:    [][]string{{"IBA"}, {"ISS", "IBA", "IEA"}, {"IC"}, {"ISS", "IEA"}, {"IC"}},
					IntersectionSize: 5,
					Name:             "regulation of skeletal muscle contraction by regulation of release of sequestered calcium ion",
					QuerySize:        5,
					Precision:        1.0,
					Pvalue:           6.045452804836617e-17,
					Recall:           1.0,
					Source:           "GO:BP",
					TermSize:         5,
				},
				{
					Genes:            []string{"CASQ2", "DMD", "GSTM2", "GSTO1"},
					ID:               "GO:0010881",
					Intersections:    [][]string{{"IMP", "IEA"}, {}, {"IC"}, {"ISS", "IEA"}, {"IC"}},
					IntersectionSize: 4,
					Name:             "regulation of cardiac muscle contraction by regulation of the release of sequestered calcium ion",
					QuerySize:        5,
					Precision:        0.8,
					Pvalue:           7.823665557395747e-09,
					Recall:           0.18181818181818182,
					Source:           "GO:BP",
					TermSize:         22,
				},
			}
			Expect(terms).To(Equal(expected))
		})
	})
})
