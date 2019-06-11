package list

import (
	"testing"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

var csvText = `field1	field2	field3
A	1	C
D	2	F
G	3	I
`

func TestCSV(t *testing.T) {
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Create test directory and files.
	fs.Instance.MkdirAll("test", 0755)
	afero.WriteFile(
		fs.Instance,
		"test/csv.txt",
		[]byte(csvText),
		0444,
	)

	wanted := []map[string]string{
		{"field1": "A", "field2": "1", "field3": "C"},
		{"field1": "D", "field2": "2", "field3": "F"},
		{"field1": "G", "field2": "3", "field3": "I"},
	}
	assert.Equal(t, wanted, CSV("test/csv.txt", '\t'), "Should parse csv file")
}
