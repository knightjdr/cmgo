package recovered

import (
	"sort"

	"github.com/knightjdr/cmgo/internal/pkg/read/geneontology"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parse genes with annotation", func() {
	It("should extract genes with specified annotation", func() {
		annotations := map[string]map[string]*geneontology.GOannotation{
			"geneA": map[string]*geneontology.GOannotation{
				"id1": &geneontology.GOannotation{},
			},
			"geneB": map[string]*geneontology.GOannotation{
				"id2": &geneontology.GOannotation{},
				"id3": &geneontology.GOannotation{},
			},
			"geneD": map[string]*geneontology.GOannotation{
				"id1": &geneontology.GOannotation{},
				"id5": &geneontology.GOannotation{},
			},
			"geneE": map[string]*geneontology.GOannotation{
				"id6": &geneontology.GOannotation{},
			},
		}

		expected := []string{"geneA", "geneD"}
		actual := readCompartmentGenes(annotations, "id1")
		sort.Strings(actual)
		Expect(actual).To(Equal(expected))
	})
})
