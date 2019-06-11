package slice

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniqueStrings(t *testing.T) {
	wanted := []string{"a", "b", "c"}
	unique := UniqueStrings([]string{"a", "b", "c", "c", "c", "b"})
	sort.Strings(unique)
	assert.Equal(t, wanted, unique, "Should remove duplicates from a slice of strings")
}
