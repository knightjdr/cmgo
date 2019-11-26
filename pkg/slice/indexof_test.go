package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexOfString(t *testing.T) {
	// TEST1: item exists
	slice := []string{"a", "b", "c", "d", "e"}
	assert.Equal(t, 2, IndexOfString("c", slice), "Should return index of element")

	// TEST2: item does not exist
	assert.Equal(t, -1, IndexOfString("f", slice), "Should return -1")
}
