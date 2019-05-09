package saint

import (
	"testing"

	"github.com/knightjdr/cmgo/fs"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

var baitDatText = `128_468	ACTB	T
128_492	ACTB	T
128_590	ATP2A1	T
128_546	ATP2A1	T
128_737	737_BirAFLAG	C
128_825	825_BirAFLAG	C
`

func TestMapBaitDatLine(t *testing.T) {
	// TEST1: at least 21 elements
	line := []string{"128_737", "737_BirAFLAG", "C"}
	wanted := BaitDatRow{
		Control: true,
		ID:      "128_737",
		Name:    "737_BirAFLAG",
	}
	assert.Equal(t, wanted, mapBaitDatLine(line), "Should map line from bait.dat file to struct")
}

func TestBaitDat(t *testing.T) {
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Create test directory and files.
	fs.Instance.MkdirAll("test", 0755)
	afero.WriteFile(
		fs.Instance,
		"test/bait.dat",
		[]byte(baitDatText),
		0444,
	)

	// TEST1: only filter by FDR
	wanted := []BaitDatRow{
		{Control: false, ID: "128_468", Name: "ACTB"},
		{Control: false, ID: "128_492", Name: "ACTB"},
		{Control: false, ID: "128_590", Name: "ATP2A1"},
		{Control: false, ID: "128_546", Name: "ATP2A1"},
		{Control: true, ID: "128_737", Name: "737_BirAFLAG"},
		{Control: true, ID: "128_825", Name: "825_BirAFLAG"},
	}
	assert.Equal(t, wanted, BaitDat("test/bait.dat"), "Should read bait.dat file")
}
