package flags

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertString(t *testing.T) {
	// TEST1: convert a value to a string
	assert.Equal(t, "a", ConvertString("a"), "Should convert an interface to a string")

	// TEST2: nil values return empty string
	assert.Equal(t, "", ConvertString(nil), "Should return an empty string")
}
