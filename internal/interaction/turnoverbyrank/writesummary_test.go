package turnoverbyrank

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
				TurnoverRates: []float64{1, 6},
			},
			2: &rankSummary{
				TurnoverRates: []float64{2, 5, 7},
			},
			3: &rankSummary{
				TurnoverRates: []float64{3},
			},
			4: &rankSummary{
				TurnoverRates: []float64{9},
			},
		}

		expected := "prey rank\tturnover rate (mean)\tturnover rate(SD)\tgenes with turnover data\n" +
			"1\t3.5000\t3.5355\t2\n" +
			"2\t4.6667\t2.5166\t3\n" +
			"3\t3.0000\t0.0000\t1\n" +
			"4\t9.0000\t0.0000\t1\n"

		writeSummary(summary, "test/out.txt")
		bytes, _ := afero.ReadFile(fs.Instance, "test/out.txt")
		Expect(string(bytes)).To(Equal(expected))
	})
})
