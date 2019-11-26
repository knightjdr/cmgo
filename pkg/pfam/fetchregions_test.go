package pfam

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var responseText = `[
	{
		"length": 212,
		"motifs": [
			{
				"end": 11,
				"start": 1,
				"type": "disorder"
			}
		],
		"regions": [
			{
				"end": 161,
				"start": 14,
				"metadata": {
					"identifier": "Domain 1"
				}
			},
			{
				"end": 200,
				"start": 170,
				"metadata": {
					"identifier": "Domain 2"
				}
			}
		]
	}
]`

var _ = Describe("Fetch from Pfam", func() {
	Context("successfully", func() {
		It("should return regions", func() {
			var apiStub = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusCreated)
				w.Write([]byte(responseText))
			}))

			ids := []string{"Q9BUL8"}
			service := &pfamService{
				URL: apiStub.URL,
			}
			fetchRegions(service, ids)

			expected := &Features{
				"Q9BUL8": Regions{
					Domains: []Domain{
						{
							End:      161,
							Metadata: Metadata{Identified: "Domain 1"},
							Name:     "Domain 1",
							Start:    14,
						},
						{
							End:      200,
							Metadata: Metadata{Identified: "Domain 2"},
							Name:     "Domain 2",
							Start:    170,
						},
					},
					Motifs: []Motif{
						{End: 11, Name: "disorder", Start: 1},
					},
				},
			}
			Expect(service.Result).To(Equal(expected))
		})
	})

	Context("missing ID", func() {
		It("should return empty struct", func() {
			var apiStub = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusCreated)
				w.Write([]byte(""))
			}))

			ids := []string{"Q9BUL8"}
			service := &pfamService{
				URL: apiStub.URL,
			}
			fetchRegions(service, ids)

			expected := &Features{
				"Q9BUL8": Regions{},
			}
			Expect(service.Result).To(Equal(expected))
		})
	})
})
