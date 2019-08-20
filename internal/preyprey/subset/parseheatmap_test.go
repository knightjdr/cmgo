package subset

import (
	"github.com/knightjdr/cmgo/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var heatmapText = `row	column	value	params
A	A	1	{"type": "heatmap"}
A	B	1 
A	C	1
B	A	1
B	B	1
B	C	1
C	A	1
C	B	1
C	C	1
`

var _ = Describe("Parse heatmap", func() {
	It("should read lines matching gene filter", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(
			fs.Instance,
			"test/heatmap.txt",
			[]byte(heatmapText),
			0444,
		)

		genes := []string{"B", "C"}

		actualLines, actualParams := parseHeatmap("test/heatmap.txt", genes)
		expectedLines := [][]string{
			{"B", "B", "1"},
			{"B", "C", "1"},
			{"C", "B", "1"},
			{"C", "C", "1"},
		}
		expectedParams := "{\"type\": \"heatmap\"}"
		Expect(actualLines).To(Equal(expectedLines), "should read lines matching gene filter")
		Expect(actualParams).To(Equal(expectedParams), "should read paramaters")
	})
})
