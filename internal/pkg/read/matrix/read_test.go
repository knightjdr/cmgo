package matrix_test

import (
	"github.com/knightjdr/cmgo/internal/pkg/read/matrix"
	"github.com/knightjdr/cmgo/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var matrixText = `gene	GO:0005694	GO:0016604	GO:0005635
AAAS	0.0000	0.0000	0.2500
AAK1	0.0000	0.2200	0.6700
`

var _ = Describe("Read matrix", func() {
	It("should read a tab-delimited file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(
			fs.Instance,
			"test/matrix.txt",
			[]byte(matrixText),
			0444,
		)

		actualRows, actualColumns, actualMatrix := matrix.Read("test/matrix.txt")
		expectedColumns := []string{"GO:0005694", "GO:0016604", "GO:0005635"}
		expectedMatrix := [][]float64{
			{0.0000, 0.0000, 0.2500},
			{0.0000, 0.2200, 0.6700},
		}
		expectedRows := []string{"AAAS", "AAK1"}
		Expect(actualColumns).To(Equal(expectedColumns), "should parse column names")
		Expect(actualMatrix).To(Equal(expectedMatrix), "should parse data matrix")
		Expect(actualRows).To(Equal(expectedRows), "should parse row names")
	})
})
