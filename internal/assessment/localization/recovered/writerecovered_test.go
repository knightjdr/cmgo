package recovered

import (
	"github.com/knightjdr/cmgo/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Write recovered genes", func() {
	It("should write all genes from compartment with boolean indicating if seen", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)

		summary := map[string]bool{
			"geneA": false,
			"geneB": true,
			"geneC": true,
			"geneD": false,
			"geneE": true,
		}

		expected := "gene\trecovered\n" +
			"geneA\tfalse\n" +
			"geneB\ttrue\n" +
			"geneC\ttrue\n" +
			"geneD\tfalse\n" +
			"geneE\ttrue\n"

		writeRecovered(summary, "test/out.txt")
		bytes, _ := afero.ReadFile(fs.Instance, "test/out.txt")
		Expect(string(bytes)).To(Equal(expected))
	})
})
