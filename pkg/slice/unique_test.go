package slice

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniqueInts(t *testing.T) {
	wanted := []int{1, 2, 4}
	unique := UniqueInts([]int{1, 2, 4, 4, 2, 2})
	unique = sort.IntSlice(unique)
	assert.Equal(t, wanted, unique, "Should remove duplicates from a slice of ints")
}

func TestUniqueStrings(t *testing.T) {
	wanted := []string{"a", "b", "c"}
	unique := UniqueStrings([]string{"a", "b", "c", "c", "c", "b"})
	sort.Strings(unique)
	assert.Equal(t, wanted, unique, "Should remove duplicates from a slice of strings")
}
