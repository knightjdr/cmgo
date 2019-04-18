package subset

import (
	"testing"

	"github.com/knightjdr/cmgo/fs"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

var basisText = `variable,1,2,3
AAAS,0.0,0.183,0.0
AAK1,0.139,0.0,0.0
AAR2,0.0,0.016,0.034
AARS2,0.0,0.0,0.002
`

func TestReadBasis(t *testing.T) {
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Create test directory and files.
	fs.Instance.MkdirAll("test", 0755)
	afero.WriteFile(
		fs.Instance,
		"test/basis.csv",
		[]byte(basisText),
		0444,
	)

	wantedColumns := []string{"1", "2", "3"}
	wantedMatrix := [][]float64{
		{0.0, 0.183, 0.0},
		{0.139, 0.0, 0.0},
		{0.0, 0.016, 0.034},
		{0.0, 0.0, 0.002},
	}
	wantedRows := []string{"AAAS", "AAK1", "AAR2", "AARS2"}
	matrix, columns, rows := readBasis("test/basis.csv")
	assert.Equal(t, wantedColumns, columns, "Should return columns")
	assert.Equal(t, wantedMatrix, matrix, "Should read basis matrix")
	assert.Equal(t, wantedRows, rows, "Should return rows")
}
