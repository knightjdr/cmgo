package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasIntersect(t *testing.T) {
	// TEST1: slices have a shared element
	sliceA := []string{"a", "b", "c", "d", "e"}
	sliceB := []string{"b", "d", "e", "f", "g"}
	assert.True(t, HasIntersect(sliceA, sliceB), "Should return true when slices share an element")

	// TEST2: slices have a shared element
	sliceA = []string{"a", "b", "c", "d", "e"}
	sliceB = []string{"w", "x", "y", "z"}
	assert.False(t, HasIntersect(sliceA, sliceB), "Should return false when two slices do not share an element")
}
