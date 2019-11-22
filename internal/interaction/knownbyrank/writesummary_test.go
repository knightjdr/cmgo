package knownbyrank

import (
	"github.com/knightjdr/cmgo/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Write summary", func() {
	It("should write summary to file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)

		summary := map[int]*rankSummary{
			1: &rankSummary{
				BaitNumber: 3,
				Known:      3,
				Pairs:      []string{"geneA-gene1", "geneB-gene4", "geneC-gene6"},
			},
			2: &rankSummary{
				BaitNumber: 3,
				Known:      2,
				Pairs:      []string{"geneB-gene5", "geneC-gene7"},
			},
			3: &rankSummary{
				BaitNumber: 2,
				Known:      1,
				Pairs:      []string{"geneA-gene3"},
			},
			4: &rankSummary{
				BaitNumber: 1,
				Known:      0,
				Pairs:      []string{},
			},
		}

		expected := "prey rank\tproportion\tnumber of baits\tknown\tpairs\n" +
			"1\t1.00\t3\t3\tgeneA-gene1, geneB-gene4, geneC-gene6\n" +
			"2\t0.67\t3\t2\tgeneB-gene5, geneC-gene7\n" +
			"3\t0.50\t2\t1\tgeneA-gene3\n" +
			"4\t0.00\t1\t0\t\n"

		writeSummary(summary, "test/out.txt")
		bytes, _ := afero.ReadFile(fs.Instance, "test/out.txt")
		Expect(string(bytes)).To(Equal(expected))
	})
})
