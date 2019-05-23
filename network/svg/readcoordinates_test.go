package svg

import (
	"testing"

	"github.com/knightjdr/cmgo/fs"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

var coordinateText = `gene	x	y
AAAS	68.5533816047702	-33.9816908466039
AAK1	-28.8672128420763	-43.6402491976645
AAR2	2.2808526181511	-6.97152015675554
`

func TestReadCoordinates(t *testing.T) {
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

	wanted := map[string]coordinate{
		"AAAS": coordinate{ X: 68.5533816047702, Y: -33.9816908466039 },
		"AAK1": coordinate{ X: -28.8672128420763, Y: -43.6402491976645 },
		"AAR2": coordinate{ X: 2.2808526181511, Y: -6.97152015675554 },
	}
	assert.Equal(t, wanted, readCoordinates("test/coordinates.txt"), "Should read coordinates from file")
}
