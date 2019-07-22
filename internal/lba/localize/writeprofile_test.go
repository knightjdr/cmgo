package localize

import (
	"github.com/knightjdr/cmgo/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Write localization profile", func() {
	It("should write to file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)

		enrichment := map[string][]Enrichment{
			"A": []Enrichment{
				{
					ID:        "GO:1",
					Precision: 0.2,
					Pvalue:    1.08e-17,
				},
				{
					ID:        "GO:2",
					Precision: 0.4,
					Pvalue:    9.08e-13,
				},
			},
			"B": []Enrichment{
				{
					ID:        "GO:3",
					Precision: 0.67,
					Pvalue:    4.56e-2,
				},
				{
					ID:        "GO:4",
					Precision: 0.447,
					Pvalue:    4.56e-2,
				},
			},
		}
		orderIDs := []string{
			"GO:2",
			"GO:1",
			"GO:3",
			"GO:4",
		}

		writeProfile(enrichment, orderIDs, "test/out.txt")
		expected := "gene\tGO:2\tGO:1\tGO:3\tGO:4\n" +
			"A\t0.4000\t0.2000\t0.0000\t0.0000\n" +
			"B\t0.0000\t0.0000\t0.6700\t0.4470\n"
		bytes, _ := afero.ReadFile(fs.Instance, "test/out.txt")
		Expect(string(bytes)).To(Equal(expected))
	})
})
