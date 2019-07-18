package evaluate

import (
	"github.com/knightjdr/cmgo/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Write summary", func() {
	It("should write to file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)

		genes := map[string]bool{
			"a": true,
			"b": true,
			"c": false,
			"d": true,
			"e": true,
			"f": true,
			"g": true,
			"h": false,
		}

		expected := "total genes: 8\n" +
			"known genes: 6\n" +
			"fraction known: 0.7500\n"
		summarize(genes, "test/out.txt")
		bytes, _ := afero.ReadFile(fs.Instance, "test/out.txt")
		Expect(string(bytes)).To(Equal(expected))
	})
})
