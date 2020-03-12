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
				cytosolicBaits:    3,
				cytosolicFraction: 0.6400,
				cytosolicScore:    0.3,
				length:            100,
				localization:      "cytosolic",
				lumenalBaits:      1,
				lumenalFraction:   0.2500,
				lumenalScore:      0.2,
				maxCytosolicScore: 0.35,
				maxLumenalScore:   0.6,
				uniprotID:         "id1",
			},
			"preyC": preySummary{
				cytosolicBaits:    2,
				cytosolicFraction: 0.3900,
				cytosolicScore:    0.15,
				length:            200,
				localization:      "lumenal",
				lumenalBaits:      2,
				lumenalFraction:   0.2800,
				lumenalScore:      0.6,
				maxCytosolicScore: 0.35,
				maxLumenalScore:   0.6,
				uniprotID:         "id2",
			},
		}

		expected := "prey\tuniprot\tlocalization\tcytosolic baits\t" +
			"lumenal baits\tcytosolic score\tlumenal score\tmax cytosolic score\tmax lumenal score\t" +
			"sequence length\tcytosolic fraction\tlumenal fraction\t" +
			"NMF orientation score\tsequence orientation score\n" +
			"preyA\tid1\tcytosolic\t3\t1\t0.300\t0.200\t0.350\t0.600\t100\t0.6400\t0.2500\t0.5238\t0.3900\n" +
			"preyC\tid2\tlumenal\t2\t2\t0.150\t0.600\t0.350\t0.600\t200\t0.3900\t0.2800\t-0.5714\t0.1100\n"

		writeSummary(summary, "test/out.txt")
		bytes, _ := afero.ReadFile(fs.Instance, "test/out.txt")
		Expect(string(bytes)).To(Equal(expected))
	})
})
