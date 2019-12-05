package rankaverage

import (
	"github.com/knightjdr/cmgo/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Write rank averages", func() {
	It("should write averages to file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)

		summary := map[string]preySummary{
			"preyA": preySummary{
				mean:  1,
				ranks: []int{1, 1, 1},
				sd:    0,
			},
			"preyB": preySummary{
				mean:  1.667,
				ranks: []int{1, 2, 2},
				sd:    0.577,
			},
			"preyC": preySummary{
				mean:  2.500,
				ranks: []int{2, 3},
				sd:    0.707,
			},
		}

		expected := "prey\tmean\tsd\tinteraction ranks\n" +
			"preyA\t1.000\t0.000\t\"1,1,1\"\n" +
			"preyB\t1.667\t0.577\t\"1,2,2\"\n" +
			"preyC\t2.500\t0.707\t\"2,3\"\n"

		writeAverages(summary, "test/out.txt")
		bytes, _ := afero.ReadFile(fs.Instance, "test/out.txt")
		Expect(string(bytes)).To(Equal(expected))
	})
})
