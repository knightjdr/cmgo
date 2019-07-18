package localize

import (
	"github.com/knightjdr/cmgo/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Write primary localization", func() {
	It("should write primary localization", func() {
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
		goIDs := map[string]string{
			"GO:1": "Term A",
			"GO:2": "Term B",
			"GO:3": "Term C",
			"GO:4": "Term D",
		}

		writePrimary(enrichment, goIDs, "test/out.txt")
		expected := "gene\tterm(s)\tID(s)\tprecision\n" +
			"A\tTerm A\tGO:1\t0.20\n" +
			"B\tTerm C;Term D\tGO:3;GO:4\t0.67;0.45\n"
		bytes, _ := afero.ReadFile(fs.Instance, "test/out.txt")
		Expect(string(bytes)).To(Equal(expected))
	})
})
