package slice

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersect(t *testing.T) {
	sliceA := []string{"a", "b", "c", "d", "e"}
	sliceB := []string{"b", "d", "e", "f", "g"}
	wanted := []string{"b", "d", "e"}
	result := Intersect(sliceA, sliceB)
	sort.Strings(result)
	assert.Equal(t, wanted, result, "Should return a slice of shared elements")
}
