package moonlighting

import (
	"github.com/knightjdr/cmgo/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Write rank moonlighting matrix", func() {
	It("should write between rank moonlighting counts to file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)

		matrix := [][]int{
			{0, 1, 0, 0, 0},
			{0, 0, 2, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 1, 0, 0},
			{1, 0, 0, 0, 0},
		}

		expected := "\t1\t2\t3\t4\t5\n" +
			"1\t0\t1\t0\t0\t0\n" +
			"2\t0\t0\t2\t0\t0\n" +
			"3\t0\t0\t0\t0\t0\n" +
			"4\t0\t0\t1\t0\t0\n" +
			"5\t1\t0\t0\t0\t0\n"

		writeRankMoonlightingMatrix(matrix, "test/out.txt")
		bytes, _ := afero.ReadFile(fs.Instance, "test/out.txt")
		Expect(string(bytes)).To(Equal(expected))
	})
})
