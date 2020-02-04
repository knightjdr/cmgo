package gene

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Map Refseq to UniProt", func() {
	Context("successfully", func() {
		It("should return a map of Refseq IDs to UniProt", func() {
			var apiStub = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusCreated)
				w.Write([]byte(uniprotResponseText))
			}))

			ids := []string{"NP_065068", "NP_000001", "NP_000002"}

			expected := map[string][]string{
				"NP_065068": []string{"A8K4L6", "Q9ULK5"},
				"NP_000001": []string{"P00001"},
				"NP_000002": []string{"P00002", "Q00002"},
			}
			Expect(RefseqToUniProt(ids, apiStub.URL)).To(Equal(expected))
		})
	})
})
