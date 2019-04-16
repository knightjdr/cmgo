package organelle

import (
	"testing"

	"github.com/knightjdr/cmgo/fs"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

var jsonText = `[
	{
		"name": "compartment 1",
		"proteins": ["a", "b", "c"]
	},
	{
		"name": "compartment 2",
		"proteins": ["d", "e", "f"]
	}
]`

func TestReadCompartments(t *testing.T) {
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Create test directory and files.
	fs.Instance.MkdirAll("test", 0755)
	afero.WriteFile(
		fs.Instance,
		"test/testfile.json",
		[]byte(jsonText),
		0444,
	)

	wanted := Compartments{
		{
			Name:     "compartment 1",
			Proteins: []string{"a", "b", "c"},
		},
		{
			Name:     "compartment 2",
			Proteins: []string{"d", "e", "f"},
		},
	}
	compartments := ReadCompartments("test/testfile.json")
	assert.Equal(t, wanted, compartments, "Should read JSON into struct")
}
