package recovered

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Count compartment genes successfully localized", func() {
	It("should extract genes with specified annotation", func() {
		compartmentGenes := []string{"geneA", "geneB", "geneC", "geneD", "geneE"}
		localizedGenes := []string{"geneB", "geneC", "geneE"}

		expected := map[string]bool{
			"geneA": false,
			"geneB": true,
			"geneC": true,
			"geneD": false,
			"geneE": true,
		}
		Expect(countRecoveredGenes(localizedGenes, compartmentGenes)).To(Equal(expected))
	})
})
