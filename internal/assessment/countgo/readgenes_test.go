package countgo

import (
	"github.com/knightjdr/cmgo/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var geneText = `gene	localization
geneA	nucleus
geneC	nucleus
geneB	nucleus
`

var _ = Describe("Read genes from a txt file", func() {
	It("should read gene names", func() {
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

		expected := []string{"geneA", "geneC", "geneB"}
		Expect(readGenes("test/genes.txt")).To(Equal(expected))
	})
})
