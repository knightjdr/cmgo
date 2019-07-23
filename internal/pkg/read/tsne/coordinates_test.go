package tsne_test

import (
	"github.com/knightjdr/cmgo/internal/pkg/read/tsne"
	"github.com/knightjdr/cmgo/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var coordinateText = `gene	x	y
AAAS	68.5533816047702	-33.9816908466039
AAK1	-28.8672128420763	-43.6402491976645
AAR2	2.2808526181511	-6.97152015675554
`

var _ = Describe("Read", func() {
	It("should read t-SNE coordinate file", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		// Create test directory and files.
		fs.Instance.MkdirAll("test", 0755)
		afero.WriteFile(
			fs.Instance,
			"test/coordinates.txt",
			[]byte(coordinateText),
			0444,
		)
		expected := map[string]tsne.Coordinate{
			"AAAS": tsne.Coordinate{X: 68.5533816047702, Y: -33.9816908466039},
			"AAK1": tsne.Coordinate{X: -28.8672128420763, Y: -43.6402491976645},
			"AAR2": tsne.Coordinate{X: 2.2808526181511, Y: -6.97152015675554},
		}
		Expect(tsne.Coordinates("test/coordinates.txt")).To(Equal(expected))
	})
})
