package nmf

import (
	"testing"

	"github.com/knightjdr/cmgo/fs"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

var localizationText = `gene	rank	score
AAAS	13	0.3326865	
AAK1	5	0.1396383	
AAR2	19	0.034698	
AARS2	6	0.2251458	
AASDH	2	0.09391108
`

func TestLocalization(t *testing.T) {
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Create test directory and files.
	fs.Instance.MkdirAll("test", 0755)
	afero.WriteFile(
		fs.Instance,
		"test/localization.txt",
		[]byte(localizationText),
		0444,
	)

	// TEST
	wanted := NMFlocalization{
		"AAAS": GeneLocalization{
			Rank:  13,
			Score: 0.3326865,
		},
		"AAK1": GeneLocalization{
			Rank:  5,
			Score: 0.1396383,
		},
		"AAR2": GeneLocalization{
			Rank:  19,
			Score: 0.034698,
		},
		"AARS2": GeneLocalization{
			Rank:  6,
			Score: 0.2251458,
		},
		"AASDH": GeneLocalization{
			Rank:  2,
			Score: 0.09391108,
		},
	}
	assert.Equal(t, wanted, Localization("test/localization.txt"), "Should read an NMF localization file")
}
