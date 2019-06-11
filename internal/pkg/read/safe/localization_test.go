package safe

import (
	"testing"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

var localizationText = `## 
## This file lists the properties of all nodes in the network.
## 

Node label	Node label ORF	Domain (predominant)	Neighborhood score [max=1, min=0] (predominant)	Total number of enriched domains	Number of enriched attributes per domain
VAMP3	VAMP3	20	0.600	1	0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,5,0,0,0,0
SNAP29	SNAP29	1	0.263	0	0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0
CDCA3	CDCA3	20	1.000	1	0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,61,0,0,0,0
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
	wanted := SAFElocalization{
		"VAMP3": GeneLocalization{
			Compartment: 20,
			Score:       0.600,
		},
		"SNAP29": GeneLocalization{
			Compartment: 1,
			Score:       0.263,
		},
		"CDCA3": GeneLocalization{
			Compartment: 20,
			Score:       1.000,
		},
	}
	assert.Equal(t, wanted, Localization("test/localization.txt"), "Should read a SAFE localization file")
}
