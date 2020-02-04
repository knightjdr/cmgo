package gene

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Map IDs", func() {
	Context("successfully", func() {
		It("should return a map of requested IDs to target", func() {
			var apiStub = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusCreated)
				w.Write([]byte(hgncResponseText))
			}))

			ids := []string{"503538", "1"}

			expected := map[string]string{
				"1":      "P04217",
				"503538": "",
			}
			Expect(MapIDs(ids, "Entrez", "UniProt", apiStub.URL)).To(Equal(expected))
		})
	})
})
