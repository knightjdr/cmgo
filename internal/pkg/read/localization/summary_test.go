package localization

import (
	"github.com/knightjdr/cmgo/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var summaryText = `rank	term	displayname	go	synonyms	ic
1	[cell junction]	[cell junction]	[GO:0030054]	[]	[1.166]
2	[chromosome]	[chromatin]	[GO:0005694]	"[[chromatid, interphase chromosome, prophase chromosome]]"	[1.256]
3	"[actin cytoskeleton, cytosol]"	"[actin cytoskeleton, cytosol]"	"[GO:0015629, GO:0005829]"	"[[], []]"	"[1.579, 0.569]"
`

var _ = Describe("Read summary file", func() {
	It("should read file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(
			fs.Instance,
			"test/summary.txt",
			[]byte(summaryText),
			0444,
		)

		expected := Summary{
			1: CompartmentSummary{
				DisplayTerms: []string{"cell junction"},
				GOid:         []string{"GO:0030054"},
				GOterms:      []string{"cell junction"},
				IC:           []float64{1.166},
			},
			2: CompartmentSummary{
				DisplayTerms: []string{"chromatin"},
				GOid:         []string{"GO:0005694"},
				GOterms:      []string{"chromosome"},
				IC:           []float64{1.256},
			},
			3: CompartmentSummary{
				DisplayTerms: []string{"actin cytoskeleton", "cytosol"},
				GOid:         []string{"GO:0015629", "GO:0005829"},
				GOterms:      []string{"actin cytoskeleton", "cytosol"},
				IC:           []float64{1.579, 0.569},
			},
		}
		Expect(SummaryFile("test/summary.txt")).To(Equal(expected))
	})
})
