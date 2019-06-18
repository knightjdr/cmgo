package robustness

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/spf13/afero"
)

var _ = Describe("Write stats", func() {
	It("should write stats to tsv file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)

		data := [][]meanSD{
			{
				{Mean: 0.5, SD: 0.25},
				{Mean: 1.5, SD: 0.6},
			},
			{
				{Mean: 2.5, SD: 0.7},
				{Mean: 2, SD: 0.4},
			},
		}
		percentile := []float64{0.9, 0.8}
		rankNames := []string{"1", "2"}
		expected := "rank\tpercentile\tmean\tSD\n" +
			"1\t0.90\t0.50000\t0.25000\n" +
			"1\t0.80\t1.50000\t0.60000\n" +
			"2\t0.90\t2.50000\t0.70000\n" +
			"2\t0.80\t2.00000\t0.40000\n"
		writeStats(data, rankNames, percentile, "test/out.txt")
		bytes, _ := afero.ReadFile(fs.Instance, "test/out.txt")
		Expect(string(bytes)).To(Equal(expected))
	})
})
