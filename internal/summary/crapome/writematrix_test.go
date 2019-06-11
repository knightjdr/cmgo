package crapome

import (
	"testing"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/knightjdr/cmgo/internal/pkg/read/saint"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestWriteMatrix(t *testing.T) {
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Create test directory and files.
	fs.Instance.MkdirAll("test", 0755)

	baits := []saint.BaitDatRow{
		{ID: "128_468"},
		{ID: "128_590"},
		{ID: "128_737"},
	}
	data := map[string]map[string]int{
		"B": map[string]int{
			"128_468": 4,
			"128_737": 8,
		},
		"A": map[string]int{
			"128_468": 5,
			"128_737": 10,
		},
		"C": map[string]int{
			"128_590": 2,
		},
		"D": map[string]int{
			"128_590": 3,
		},
	}
	idToCCmap := map[int]string{
		468: "CC1",
		590: "CC2",
		737: "CC3",
	}
	preyMap := map[string]string{
		"A": "preyA",
		"B": "preyB",
		"C": "preyC",
		"D": "preyD",
	}
	preyOrder := []string{"A", "B", "C", "D"}
	wanted := "GENE\tREFSEQ_ID\tAVE_SC\tNUM_EXPT\tCC1\tCC2\tCC3\n" +
		"preyA\tA\t7.50\t2\t5\t0\t10\n" +
		"preyB\tB\t6.00\t2\t4\t0\t8\n" +
		"preyC\tC\t2.00\t1\t0\t2\t0\n" +
		"preyD\tD\t3.00\t1\t0\t3\t0\n"
	writeMatrix(data, baits, preyMap, preyOrder, idToCCmap, "test/out.txt")
	bytes, _ := afero.ReadFile(fs.Instance, "test/out.txt")
	assert.Equal(t, wanted, string(bytes), "Should write matrix to file")
}
