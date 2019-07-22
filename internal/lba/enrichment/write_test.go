package enrichment

import (
	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/knightjdr/cmgo/pkg/gprofiler"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Find Refseq ID from Entrez", func() {
	It("should return a Refseq ID", func() {
		refseqMap := map[string]map[string]string{
			"A": map[string]string{
				"Entrez": "3",
			},
			"B": map[string]string{
				"Entrez": "1",
			},
			"C": map[string]string{
				"Entrez": "2",
			},
		}
		tests := []string{"1", "2", "3"}
		expected := []string{"B", "C", "A"}
		for i := range tests {
			Expect(findRefseq(tests[i], refseqMap)).To(Equal(expected[i]))
		}
	})

	It("should return nil string", func() {
		refseqMap := map[string]map[string]string{
			"A": map[string]string{
				"Entrez": "3",
			},
			"B": map[string]string{
				"Entrez": "1",
			},
			"C": map[string]string{
				"Entrez": "2",
			},
		}
		tests := []string{"4", "5", "6"}
		for i := range tests {
			Expect(findRefseq(tests[i], refseqMap)).To(Equal(""))
		}
	})
})

var _ = Describe("Write results to file", func() {
	It("should write enrichment results to file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)

		enrichment := map[string][]gprofiler.EnrichedTerm{
			"A": []gprofiler.EnrichedTerm{
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
			},
			"B": []gprofiler.EnrichedTerm{
				{
					Genes:     []string{"3", "2"},
					ID:        "GO:3",
					Name:      "Term 3",
					Pvalue:    4.56e-2,
					Recall:    0.01,
					Precision: 0.50,
					QuerySize: 2,
					TermSize:  100,
				},
			},
			"C": []gprofiler.EnrichedTerm{
				{
					Genes:     []string{"3", "1"},
					ID:        "GO:4",
					Name:      "Term 4",
					Pvalue:    7.11e-7,
					Recall:    0.01,
					Precision: 0.50,
					QuerySize: 2,
					TermSize:  100,
				},
			},
		}
		refseqMap := map[string]map[string]string{
			"A": map[string]string{
				"Entrez":  "3",
				"Symbol":  "a",
				"UniProt": "QA",
			},
			"B": map[string]string{
				"Entrez":  "1",
				"Symbol":  "b",
				"UniProt": "QB",
			},
			"C": map[string]string{
				"Entrez":  "2",
				"Symbol":  "c",
				"UniProt": "QC",
			},
		}

		write(enrichment, refseqMap, "test/out.txt")
		expected := "symbol\tEntrez\tRefseq\tUniProt\tGO ID\tGO term\tp-value\trecall\tprecision\tquery size\tterm size\tgenes\n" +
			"a\t3\tA\tQA\tGO:1\tTerm 1\t1.08e-17\t0.0100\t0.5000\t2\t100\tb, c\n" +
			"a\t3\tA\tQA\tGO:2\tTerm 2\t9.08e-13\t0.2000\t1.0000\t1\t5\tc\n" +
			"b\t1\tB\tQB\tGO:3\tTerm 3\t4.56e-02\t0.0100\t0.5000\t2\t100\ta, c\n" +
			"c\t2\tC\tQC\tGO:4\tTerm 4\t7.11e-07\t0.0100\t0.5000\t2\t100\ta, b\n"
		bytes, _ := afero.ReadFile(fs.Instance, "test/out.txt")
		Expect(string(bytes)).To(Equal(expected))
	})
})
