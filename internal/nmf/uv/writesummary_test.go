package uv

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/spf13/afero"
)

var _ = Describe("Write summary", func() {
	It("should write data to tsv file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)

		data := []map[string][]string{
			{
				"known":   []string{"a", "b"},
				"unknown": []string{"c", "d", "e"},
			},
			{
				"known":   []string{"f", "g"},
				"unknown": []string{"i", "h"},
			},
		}
		expected := "rank\tknown\tunknown\tfraction\tgenes known\tgenes unknown\n" +
			"1\t2\t3\t0.40\ta, b\tc, d, e\n" +
			"2\t2\t2\t0.50\tf, g\th, i\n"
		writeSummary(data, "test/out.txt")
		bytes, _ := afero.ReadFile(fs.Instance, "test/out.txt")
		Expect(string(bytes)).To(Equal(expected))
	})
})
