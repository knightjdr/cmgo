package saint

import (
	"testing"

	"github.com/knightjdr/cmgo/fs"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

var preyDatText = `BirA_R118G_H0QFJ5	321	BirA_R118G_H0QFJ5
NP_000029.2	2843	APC
NP_000108.1	254	EMD
NP_000280.1	780	PFKM
NP_000402.3	726	HLCS
`

func TestMapPreyDatLine(t *testing.T) {
	// TEST1: at least 21 elements
	line := []string{"NP_000029.2", "2843", "APC"}
	wanted := PreyDatRow{
		Accession: "NP_000029.2",
		Length:    2843,
		Name:      "APC",
	}
	assert.Equal(t, wanted, mapPreyDatLine(line), "Should map line from prey.dat file to struct")
}

func TestPreyDat(t *testing.T) {
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Create test directory and files.
	fs.Instance.MkdirAll("test", 0755)
	afero.WriteFile(
		fs.Instance,
		"test/prey.dat",
		[]byte(preyDatText),
		0444,
	)

	// TEST1: only filter by FDR
	wanted := []PreyDatRow{
		{Accession: "BirA_R118G_H0QFJ5", Length: 321, Name: "BirA_R118G_H0QFJ5"},
		{Accession: "NP_000029.2", Length: 2843, Name: "APC"},
		{Accession: "NP_000108.1", Length: 254, Name: "EMD"},
		{Accession: "NP_000280.1", Length: 780, Name: "PFKM"},
		{Accession: "NP_000402.3", Length: 726, Name: "HLCS"},
	}
	assert.Equal(t, wanted, PreyDat("test/prey.dat"), "Should read bait.dat file")
}
