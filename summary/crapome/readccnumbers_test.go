package crapome

import (
	"testing"

	"github.com/knightjdr/cmgo/fs"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

var ccNumbersText = `filename	cc
9522_BirAFLAG_April2017_Go_RIPA_07012018.raw	CC1100
9221_BirAFLAG_HighDen_Go_KA_RIPA_10112017_BR1.raw	CC1101
7909_BirAFLAG_April2017_Go_RIPA_06042017_TR1.raw	CC1102
7024_BirAFLAG_May2016_test_Go_RIPA_31102016_TR1	CC1103
7015_BirAFLAG_May2016_Go_RIPA_31102016_TR1.raw	CC1104
`

func TestReadCC(t *testing.T) {
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Create test directory and files.
	fs.Instance.MkdirAll("test", 0755)
	afero.WriteFile(
		fs.Instance,
		"test/cc.txt",
		[]byte(ccNumbersText),
		0444,
	)

	// TEST
	wanted := map[int]string{
		9522: "CC1100",
		9221: "CC1101",
		7909: "CC1102",
		7024: "CC1103",
		7015: "CC1104",
	}
	assert.Equal(t, wanted, readCC("test/cc.txt"), "Should read CC numbers from file to map")
}
