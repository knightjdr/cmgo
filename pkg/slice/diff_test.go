package slice

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiff(t *testing.T) {
	sliceA := []string{"a", "b", "c", "d", "e"}
	sliceB := []string{"b", "d", "e", "f", "g"}
	wanted := []string{"a", "c", "f", "g"}
	result := Diff(sliceA, sliceB)
	sort.Strings(result)
	assert.Equal(t, wanted, result, "Should return a slice of different elements")
}
