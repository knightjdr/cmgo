package countgo

import (
	"github.com/knightjdr/cmgo/internal/pkg/read/geneontology"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Count genes with annotations", func() {
	It("should count the number of genes with at least one annotation", func() {
		annotations := map[string]map[string]*geneontology.GOannotation{
			"geneA": map[string]*geneontology.GOannotation{
				"id1": &geneontology.GOannotation{},
			},
			"geneB": map[string]*geneontology.GOannotation{
				"id2": &geneontology.GOannotation{},
				"id3": &geneontology.GOannotation{},
			},
			"geneD": map[string]*geneontology.GOannotation{
				"id4": &geneontology.GOannotation{},
				"id5": &geneontology.GOannotation{},
			},
			"geneE": map[string]*geneontology.GOannotation{
				"id6": &geneontology.GOannotation{},
			},
		}
		genes := []string{"geneA", "geneB", "geneC", "geneD", "geneE", "geneF"}

		expectedWith := []string{"geneA", "geneB", "geneD", "geneE"}
		expectedWithout := []string{"geneC", "geneF"}

		actualCount, actualWith, actualWithout := countGenesWithAnnotation(genes, annotations)
		Expect(actualCount).To(Equal(4), "Should count the number of genes with annotations")
		Expect(actualWith).To(Equal(expectedWith), "Should return a list of genes with annotations")
		Expect(actualWithout).To(Equal(expectedWithout), "Should return a list of genes without annotations")
	})
})
