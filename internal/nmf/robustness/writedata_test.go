package robustness

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/spf13/afero"
)

var _ = Describe("Writedata", func() {
	It("should write data to tsv file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)

		data := [][][]float64{
			{
				{1, 2, 3},
				{1, 2, 3},
			},
			{
				{1, 2, 3},
				{1, 2, 3},
			},
		}
		percentile := []float64{0.9, 0.8}
		rankNames := []string{"1", "2"}
		replicates := 3
		expected := "rank\tpercentile\treplicate\tRBD\n" +
			"1\t0.90\t1\t1.00000\n" +
			"1\t0.90\t2\t2.00000\n" +
			"1\t0.90\t3\t3.00000\n" +
			"1\t0.80\t1\t1.00000\n" +
			"1\t0.80\t2\t2.00000\n" +
			"1\t0.80\t3\t3.00000\n" +
			"2\t0.90\t1\t1.00000\n" +
			"2\t0.90\t2\t2.00000\n" +
			"2\t0.90\t3\t3.00000\n" +
			"2\t0.80\t1\t1.00000\n" +
			"2\t0.80\t2\t2.00000\n" +
			"2\t0.80\t3\t3.00000\n"
		writeData(data, rankNames, percentile, replicates, "test/out.txt")
		bytes, _ := afero.ReadFile(fs.Instance, "test/out.txt")
		Expect(string(bytes)).To(Equal(expected))
	})
})
