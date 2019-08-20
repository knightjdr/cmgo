package genes

import (
	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/knightjdr/cmgo/pkg/gprofiler"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Write results to file", func() {
	It("should write enrichment results to file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)

		enrichment := []gprofiler.EnrichedTerm{
			{
				Genes:     []string{"1", "2"},
				ID:        "GO:1",
				Name:      "Term 1",
				Pvalue:    1.08e-17,
				Recall:    0.01,
				Precision: 0.50,
				QuerySize: 2,
				TermSize:  100,
			},
			{
				Genes:     []string{"2"},
				ID:        "GO:2",
				Name:      "Term 2",
				Pvalue:    9.08e-13,
				Recall:    0.20,
				Precision: 1.00,
				QuerySize: 1,
				TermSize:  5,
			},
		}

		write(enrichment, "test/out.txt")
		expected := "GO ID\tGO term\tp-value\trecall\tprecision\tquery size\tterm size\tno. genes\tgenes\n" +
			"GO:1\tTerm 1\t1.08e-17\t0.0100\t0.5000\t2\t100\t2\t1, 2\n" +
			"GO:2\tTerm 2\t9.08e-13\t0.2000\t1.0000\t1\t5\t1\t2\n"
		bytes, _ := afero.ReadFile(fs.Instance, "test/out.txt")
		Expect(string(bytes)).To(Equal(expected))
	})
})
