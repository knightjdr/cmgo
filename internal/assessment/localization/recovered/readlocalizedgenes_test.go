package recovered

import (
	"github.com/knightjdr/cmgo/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var geneText = `gene	localization
geneA	nucleus
geneC	cytosol
geneB	nucleus
`

var _ = Describe("Read genes from a txt file", func() {
	It("should read gene names that localize to specified compartment", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(
			fs.Instance,
			"test/genes.txt",
			[]byte(geneText),
			0444,
		)

		expected := []string{"geneA", "geneB"}
		Expect(readLocalizedGenes("test/genes.txt", "nucleus")).To(Equal(expected))
	})
})
