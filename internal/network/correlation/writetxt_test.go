package correlation

import (
	"github.com/knightjdr/cmgo/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Write txt", func() {
	It("should write data to tab-delimited file to file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)

		genes := []string{"a", "b", "c", "d"}
		pairs := map[string][]edgePair{
			"a": []edgePair{
				{Target: "c", Weight: 0.71},
				{Target: "d", Weight: 0.82},
			},
			"b": []edgePair{
				{Target: "c", Weight: 0.95},
			},
			"c": []edgePair{
				{Target: "d", Weight: 0.71},
			},
			"d": []edgePair{},
		}

		expected := "source\ttarget\tweight\n" +
			"a\tc\t0.7100\n" +
			"a\td\t0.8200\n" +
			"b\tc\t0.9500\n" +
			"c\td\t0.7100\n"
		writeTXT(genes, pairs, "test/out.txt")
		bytes, _ := afero.ReadFile(fs.Instance, "test/out.txt")
		Expect(string(bytes)).To(Equal(expected))
	})
})
