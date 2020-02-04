package rankmetrics

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
				Expression:    []float64{2, 4, 6},
				Lysines:       []int{11, 16},
				TurnoverRates: []float64{1, 6},
			},
			2: &rankSummary{
				Expression:    []float64{5},
				Lysines:       []int{12, 17},
				TurnoverRates: []float64{2, 5, 7},
			},
			3: &rankSummary{
				Expression:    []float64{5, 10},
				Lysines:       []int{13},
				TurnoverRates: []float64{3},
			},
			4: &rankSummary{
				Expression:    []float64{2, 4},
				Lysines:       []int{19},
				TurnoverRates: []float64{9},
			},
		}

		expected := "prey rank\t" +
			"expression (mean)\texpression (SD)\tgenes with expression data\t" +
			"lysines (mean)\tlysines (SD)\tgenes with lysine data\t" +
			"turnover rate (mean)\tturnover rate(SD)\tgenes with turnover data\n" +
			"1\t4.0000\t2.0000\t3\t13.5000\t3.5355\t2\t3.5000\t3.5355\t2\n" +
			"2\t5.0000\t0.0000\t1\t14.5000\t3.5355\t2\t4.6667\t2.5166\t3\n" +
			"3\t7.5000\t3.5355\t2\t13.0000\t0.0000\t1\t3.0000\t0.0000\t1\n" +
			"4\t3.0000\t1.4142\t2\t19.0000\t0.0000\t1\t9.0000\t0.0000\t1\n"

		writeSummary(summary, "test/out.txt")
		bytes, _ := afero.ReadFile(fs.Instance, "test/out.txt")
		Expect(string(bytes)).To(Equal(expected))
	})
})
