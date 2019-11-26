package transmembrane

import (
	"github.com/knightjdr/cmgo/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Write summary", func() {
	It("should write prey summary information to file to file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)

		summary := map[string]preySummary{
			"preyA": preySummary{
				cytosolicBaits: 3,
				cytosolicScore: 0.3,
				localization:   "cytosolic",
				lumenalBaits:   1,
				lumenalScore:   0.2,
			},
			"preyC": preySummary{
				cytosolicBaits: 2,
				cytosolicScore: 0.15,
				localization:   "lumenal",
				lumenalBaits:   2,
				lumenalScore:   0.6,
			},
		}

		expected := "prey\tlocalization\tcytosolic baits\tlumenal baits\tcytosolic score\tlumenal score\n" +
			"preyA\tcytosolic\t3\t1\t0.300\t0.200\n" +
			"preyC\tlumenal\t2\t2\t0.150\t0.600\n"

		writeSummary(summary, "test/out.txt")
		bytes, _ := afero.ReadFile(fs.Instance, "test/out.txt")
		Expect(string(bytes)).To(Equal(expected))
	})
})
