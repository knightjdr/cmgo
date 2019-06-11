package overlap

import (
	"testing"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

var similarityText = `source	target	distance
a	a	1
a	b	0.5
a	c	0.25
b	a	0.5
b	b	1
b	c	0.63
c	a	0.25
c	b	0.63
c	c	1`

func TestReadSimilarity(t *testing.T) {
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Create test directory and files.
	fs.Instance.MkdirAll("test", 0755)
	afero.WriteFile(
		fs.Instance,
		"test/testfile.txt",
		[]byte(similarityText),
		0444,
	)

	wanted := map[string]map[string]float64{
		"a": {"b": 0.5, "c": 0.25},
		"b": {"c": 0.63},
	}
	similarity := readSimilarity("test/testfile.txt")
	assert.Equal(t, wanted, similarity, "Should read similarity file to map")
}
