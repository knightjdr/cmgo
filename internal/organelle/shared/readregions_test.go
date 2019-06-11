package shared

import (
	"testing"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

var regionsText = `A1BG	disorder
A1BG	sig_p
A1CF	disorder
A1CF	low_complexity
A2M	disorder
A2M	low_complexity
A2M	sig_p
A2ML1	disorder
A2ML1	low_complexity
A2ML1	sig_p
A3GALT2	low_complexity
A3GALT2	transmembrane`

func TestReadRegions(t *testing.T) {
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Create test directory and files.
	fs.Instance.MkdirAll("test", 0755)
	afero.WriteFile(
		fs.Instance,
		"test/regions.txt",
		[]byte(regionsText),
		0444,
	)

	wanted := map[string]map[string]bool{
		"A1BG":    {"disorder": true, "sig_p": true},
		"A1CF":    {"disorder": true, "low_complexity": true},
		"A2M":     {"disorder": true, "low_complexity": true, "sig_p": true},
		"A2ML1":   {"disorder": true, "low_complexity": true, "sig_p": true},
		"A3GALT2": {"low_complexity": true, "transmembrane": true},
	}
	regions := readRegions("test/regions.txt")
	assert.Equal(t, wanted, regions, "Should read regions file to map")
}
