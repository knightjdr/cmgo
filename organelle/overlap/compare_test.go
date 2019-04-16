package overlap

import (
	"testing"

	"github.com/knightjdr/cmgo/fs"
	"github.com/knightjdr/cmgo/organelle"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestCompartmentDict(t *testing.T) {
	wanted := map[string]bool{
		"a": true,
		"b": true,
		"c": true,
	}
	dict := compartmentDict([]string{"a", "b", "c"})
	assert.Equal(t, wanted, dict, "Should convert slice to hash")
}

func TestRangeIndex(t *testing.T) {
	dict1 := map[string]bool{
		"a": true,
		"b": true,
	}
	dict2 := map[string]bool{
		"d": true,
		"e": true,
	}

	// TEST1
	result := rangeIndex("a", "b", dict1, dict2)
	assert.Equal(t, 0, result, "Should return 0 if source and target in the first map")

	// TEST2
	result = rangeIndex("d", "e", dict1, dict2)
	assert.Equal(t, 1, result, "Should return 1 if source and target in the second map")

	// TEST3
	result = rangeIndex("a", "d", dict1, dict2)
	assert.Equal(t, 2, result, "Should return 2 if source and target in different maps")

	// TEST4
	result = rangeIndex("a", "x", dict1, dict2)
	assert.Equal(t, -1, result, "Should return -1 is target is not in either map")

	// TEST5
	result = rangeIndex("x", "d", dict1, dict2)
	assert.Equal(t, -1, result, "Should return -1 if source is not in either map")

	// TEST6
	result = rangeIndex("x", "y", dict1, dict2)
	assert.Equal(t, -1, result, "Should return -1 is neither source or target in either map")
}

func TestCompare(t *testing.T) {
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Create test directory and files.
	fs.Instance.MkdirAll("test", 0755)

	compartments := organelle.Compartments{
		{
			Name:     "compartment 1",
			Proteins: []string{"a", "b", "c"},
		},
		{
			Name:     "compartment 2",
			Proteins: []string{"d", "e", "f"},
		},
	}
	similarity := map[string]map[string]float64{
		"a": {"b": 0.5, "c": 0.25, "d": 0.1, "e": 0.4, "f": 0.2},
		"b": {"c": 0.6, "d": 0.3, "e": 0.15, "f": 0.7},
		"c": {"d": 0.5, "e": 0.5, "f": 0.25},
		"d": {"e": 0.4, "f": 0.2},
		"e": {"f": 0.2},
	}
	want := "\tmedian\tmean\tmin\tmax\ncompartment 1\t0.500\t0.450\t0.250\t0.600\ncompartment 2\t0.200\t0.267\t0.200\t0.400\nbetween\t0.300\t0.344\t0.100\t0.700\n"
	compare(compartments, similarity, "test/out.txt")
	bytes, _ := afero.ReadFile(fs.Instance, "test/out.txt")
	assert.Equal(t, want, string(bytes), "Should write summary metrics to file")
}
