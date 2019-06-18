package gene

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var responseText = `{
	"responseHeader":{"status":0,"QTime":18},
	"response":{
		"numFound":41851,
		"start":0,
		"docs":[
			{
				"hgnc_id":"HGNC:5",
				"symbol":"A1BG",
				"name":"alpha-1-B glycoprotein",
				"entrez_id":"1",
				"ensembl_gene_id":"ENSG00000121410",
				"uniprot_ids":["P04217"]
			},
			{
				"hgnc_id":"HGNC:37133",
				"symbol":"A1BG-AS1",
				"name":"A1BG antisense RNA 1",
				"entrez_id":"503538",
				"ensembl_gene_id":"ENSG00000268895"
			}
		]
	}
}`

var _ = Describe("Fetch from HGNC", func() {
	Context("successfully", func() {
		It("should return gene IDS for mapping", func() {
			var apiStub = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusCreated)
				w.Write([]byte(responseText))
			}))

			service := &hgncService{
				URL: apiStub.URL,
			}
			fetchHGNC(service)

			expected := []map[string]string{
				{
					"EnsemblGene": "ENSG00000121410",
					"Entrez":      "1",
					"HGNC":        "HGNC:5",
					"Name":        "alpha-1-B glycoprotein",
					"Symbol":      "A1BG",
					"UniProt":     "P04217",
				},
				{
					"EnsemblGene": "ENSG00000268895",
					"Entrez":      "503538",
					"HGNC":        "HGNC:37133",
					"Name":        "A1BG antisense RNA 1",
					"Symbol":      "A1BG-AS1",
					"UniProt":     "",
				},
			}
			Expect(service.Result).To(Equal(expected))
		})
	})
})
