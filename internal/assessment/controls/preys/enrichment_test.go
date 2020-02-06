package preys

import (
	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/knightjdr/cmgo/pkg/gprofiler"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Write GO enrichment", func() {
	It("should write results", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("test", 0755)

		results := []gprofiler.EnrichedTerm{
			gprofiler.EnrichedTerm{
				Name:             "Term1",
				ID:               "1",
				Source:           "CC",
				QuerySize:        10,
				TermSize:         100,
				IntersectionSize: 5,
				Precision:        0.500,
				Recall:           0.050,
				Pvalue:           0.001,
				Genes:            []string{"geneA", "geneB"},
			},
			gprofiler.EnrichedTerm{
				Name:             "Term2",
				ID:               "2",
				Source:           "BP",
				QuerySize:        20,
				TermSize:         200,
				IntersectionSize: 10,
				Precision:        0.500,
				Recall:           0.050,
				Pvalue:           0.01,
				Genes:            []string{"geneC", "geneD"},
			},
		}

		expected := "term\tid\tsource\tquery size\tterm size\tintersection size\tprecision\trecall\tp-value\tgenes\n" +
			"Term1\t1\tCC\t10\t100\t5\t0.500\t0.050\t1.000000e-03\tgeneA, geneB\n" +
			"Term2\t2\tBP\t20\t200\t10\t0.500\t0.050\t1.000000e-02\tgeneC, geneD\n"

		writeEnrichment(results, "test/out.txt")
		bytes, _ := afero.ReadFile(fs.Instance, "test/out.txt")
		Expect(string(bytes)).To(Equal(expected))
	})
})
