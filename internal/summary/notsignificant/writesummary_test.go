package notsignificant

import (
	"testing"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestWriteSummary(t *testing.T) {
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Create test directory and files.
	fs.Instance.MkdirAll("test", 0755)

	summary := map[string]*preySummary{
		"prey-1": &preySummary{baits: []string{"a", "b", "c"}, bestFDR: 0.01, ctrlAvg: 4, maxSpec: 47},
		"prey-2": &preySummary{baits: []string{"a", "c"}, bestFDR: 0, ctrlAvg: 15, maxSpec: 15},
		"prey-3": &preySummary{baits: []string{"a", "b"}, bestFDR: 0.01, ctrlAvg: 25, maxSpec: 5},
		"prey-4": &preySummary{baits: []string{"c"}, bestFDR: 0, ctrlAvg: 0, maxSpec: 35.5},
	}
	wanted := "prey\tbaits\tmaximum spectral count\taverage control count\tbest FDR\nprey-1\ta, b, c\t47.00\t4.00\t0.01\nprey-2\ta, c\t15.00\t15.00\t0.00\nprey-3\ta, b\t5.00\t25.00\t0.01\nprey-4\tc\t35.50\t0.00\t0.00\n"
	writeSummary(summary, "test/out.txt")
	bytes, _ := afero.ReadFile(fs.Instance, "test/out.txt")
	assert.Equal(t, wanted, string(bytes), "Should write summary to file")
}
