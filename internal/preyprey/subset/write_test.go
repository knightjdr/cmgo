package subset

import (
	"github.com/knightjdr/cmgo/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Write results to file", func() {
	It("should write heatmap cluster to file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)

		lines := [][]string{
			{"B", "B", "1"},
			{"B", "C", "1"},
			{"C", "B", "1"},
			{"C", "C", "1"},
		}
		params := "{\"type\": \"heatmap\"}"

		write(lines, params, "test/out.txt")
		expected := "row\tcolumn\tvalue\tparams\n" +
			"B\tB\t1\t{\"type\": \"heatmap\"}\n" +
			"B\tC\t1\n" +
			"C\tB\t1\n" +
			"C\tC\t1\n"
		bytes, _ := afero.ReadFile(fs.Instance, "test/out.txt")
		Expect(string(bytes)).To(Equal(expected))
	})
})
