package list

import (
	"testing"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

var listText = `A
BB
D
F
`

func TestParseSlice(t *testing.T) {
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Create test directory and files.
	fs.Instance.MkdirAll("test", 0755)
	afero.WriteFile(
		fs.Instance,
		"test/list.txt",
		[]byte(listText),
		0444,
	)

	wanted := []string{"A", "BB", "D", "F"}
	assert.Equal(t, wanted, ParseSlice("test/list.txt"), "Should parse a list to a string slice")
}
