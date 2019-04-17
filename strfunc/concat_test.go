package strfunc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcat(t *testing.T) {
	slice := []string{"a", "bc", "d\ne"}

	// TEST1: concatenates strings
	want := "abcd\ne"
	assert.Equal(t, want, Concat(slice), "Should concatenate slice of strings")

	// TEST2: empty slice
	want = ""
	assert.Equal(t, want, Concat([]string{}), "Should handle empty slice")
}
