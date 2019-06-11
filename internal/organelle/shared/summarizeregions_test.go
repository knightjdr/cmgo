package shared

import (
	"testing"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestSummarizeRegions(t *testing.T) {
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Create test directory and files.
	fs.Instance.MkdirAll("test", 0755)

	proteins := []string{"a", "b", "c", "d"}
	regions := map[string]map[string]bool{
		"a": {"regionA": true, "regionC": true, "regionD": true},
		"b": {"regionA": true, "regionB": true, "regionC": true},
		"c": {"regionC": true, "regionD": true},
		"d": {"regionC": true, "regionD": true},
	}
	wanted := "region\tno. preys\tpreys\tpreys not containing region\n-\t4\ta, b, c, d\t\nregionC\t4\ta, b, c, d\t\nregionD\t3\ta, c, d\tb\nregionA\t2\ta, b\tc, d\nregionB\t1\tb\ta, c, d\n"
	summarizeRegions(proteins, regions, "test/out.txt")
	bytes, _ := afero.ReadFile(fs.Instance, "test/out.txt")
	assert.Equal(t, wanted, string(bytes), "Should write region summary to file")
}
