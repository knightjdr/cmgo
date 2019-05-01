package dbgenes

import (
	"testing"

	"github.com/knightjdr/cmgo/fs"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestWriteGenes(t *testing.T) {
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Create test directory and files.
	fs.Instance.MkdirAll("test", 0755)

	genes := []string{"A", "B", "C", "e", "F", "D"}
	wanted := "A\nB\nC\nD\ne\nF\n"
	writeGenes(genes, "test/out.txt")
	bytes, _ := afero.ReadFile(fs.Instance, "test/out.txt")
	assert.Equal(t, wanted, string(bytes), "Should write genes to file sorted alphabetically")
}
