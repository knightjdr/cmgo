package read

import (
	"testing"

	"github.com/knightjdr/cmgo/fs"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

var interDatText = `128_468	ACTB	BirA_R118G_H0QFJ5	1410	23
128_468	ACTB	NP_000029.2	4	3
128_468	ACTB	NP_000108.1	2	2
128_468	ACTB	NP_000280.1	3	2
128_468	ACTB	NP_000402.3	11	10
`

func TestMapInterDatLine(t *testing.T) {
	// TEST1: at least 21 elements
	line := []string{"128_468", "ACTB", "NP_000029.2", "4", "3"}
	wanted := InterDatRow{
		Bait: "ACTB",
		ID:   "128_468",
		Prey: "NP_000029.2",
		Spec: 4,
	}
	assert.Equal(t, wanted, mapInterDatLine(line), "Should map line from inter.dat file to struct")
}

func TestInterDat(t *testing.T) {
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Create test directory and files.
	fs.Instance.MkdirAll("test", 0755)
	afero.WriteFile(
		fs.Instance,
		"test/inter.dat",
		[]byte(interDatText),
		0444,
	)

	// TEST1: only filter by FDR
	wanted := []InterDatRow{
		{Bait: "ACTB", ID: "128_468", Prey: "BirA_R118G_H0QFJ5", Spec: 1410},
		{Bait: "ACTB", ID: "128_468", Prey: "NP_000029.2", Spec: 4},
		{Bait: "ACTB", ID: "128_468", Prey: "NP_000108.1", Spec: 2},
		{Bait: "ACTB", ID: "128_468", Prey: "NP_000280.1", Spec: 3},
		{Bait: "ACTB", ID: "128_468", Prey: "NP_000402.3", Spec: 11},
	}
	assert.Equal(t, wanted, InterDat("test/inter.dat"), "Should read bait.dat file")
}
