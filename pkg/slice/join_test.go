package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJoin(t *testing.T) {
	slice := []int{1, 2, 3}

	expected := "1, 2, 3"
	assert.Equal(t, expected, JoinInts(slice, ", "), "should join a slice of ints")
}
