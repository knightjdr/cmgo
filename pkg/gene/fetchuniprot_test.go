package gene

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var uniprotResponseText = `From	To
NP_065068	A8K4L6
NP_065068	Q9ULK5
NP_000001	P00001
NP_000002	P00002
NP_000002	Q00002
`

var _ = Describe("Map Refseq Ids from UniProt", func() {
	Context("successfully", func() {
		It("should return mapping from UniProt to Refseq", func() {
			var apiStub = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusCreated)
				w.Write([]byte(uniprotResponseText))
			}))

			service := &uniprotService{
				Query: []string{"NP_065068", "NP_000001", "NP_000002"},
				URL:   apiStub.URL,
			}
			fetchUniProt(service)

			expected := map[string][]string{
				"NP_065068": []string{"A8K4L6", "Q9ULK5"},
				"NP_000001": []string{"P00001"},
				"NP_000002": []string{"P00002", "Q00002"},
			}
			Expect(service.Result).To(Equal(expected))
		})
	})
})

var _ = Describe("Subset query string", func() {
	It("should subset first two elements from query", func() {
		oldLength := maxQueryLength
		maxQueryLength = 10
		defer func() {
			maxQueryLength = oldLength
		}()

		query := []string{"aaaaa", "bbbbb", "ccccc", "dddddd", "eeeee"}

		expectedStart := 2
		expectedSubset := []string{"aaaaa", "bbbbb"}

		actualSubset, actualStart := getNextQuerySubset(query, 0)
		Expect(actualStart).To(Equal(expectedStart), "should return start index for next subset")
		Expect(actualSubset).To(Equal(expectedSubset), "should subset query")
	})

	It("should subset third element from query", func() {
		oldLength := maxQueryLength
		maxQueryLength = 10
		defer func() {
			maxQueryLength = oldLength
		}()

		query := []string{"aaaaa", "bbbbb", "ccccc", "dddddd", "eeeee"}

		expectedStart := 3
		expectedSubset := []string{"ccccc"}

		actualSubset, actualStart := getNextQuerySubset(query, 2)
		Expect(actualStart).To(Equal(expectedStart), "should return start index for next subset")
		Expect(actualSubset).To(Equal(expectedSubset), "should subset query")
	})

	It("should subset fourth element from query", func() {
		oldLength := maxQueryLength
		maxQueryLength = 10
		defer func() {
			maxQueryLength = oldLength
		}()

		query := []string{"aaaaa", "bbbbb", "ccccc", "dddddd", "eeeee"}

		expectedStart := 4
		expectedSubset := []string{"dddddd"}

		actualSubset, actualStart := getNextQuerySubset(query, 3)
		Expect(actualStart).To(Equal(expectedStart), "should return start index for next subset")
		Expect(actualSubset).To(Equal(expectedSubset), "should subset query")
	})

	It("should subset fifth and last element from query", func() {
		oldLength := maxQueryLength
		maxQueryLength = 10
		defer func() {
			maxQueryLength = oldLength
		}()

		query := []string{"aaaaa", "bbbbb", "ccccc", "dddddd", "eeeee"}

		expectedStart := 5
		expectedSubset := []string{"eeeee"}

		actualSubset, actualStart := getNextQuerySubset(query, 4)
		Expect(actualStart).To(Equal(expectedStart), "should return start index for next subset")
		Expect(actualSubset).To(Equal(expectedSubset), "should subset query")
	})
})
