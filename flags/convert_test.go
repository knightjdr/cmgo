package flags

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertFloat(t *testing.T) {
	// TEST1: convert a value to a float
	assert.Equal(t, 0.01, ConvertFloat("0.01"), "Should convert an interface to a float")

	// TEST2: nil values return empty string
	assert.Equal(t, float64(0), ConvertFloat(nil), "Should return an nil float")
}

func TestConvertString(t *testing.T) {
	// TEST1: convert a value to a string
	assert.Equal(t, "a", ConvertString("a"), "Should convert an interface to a string")

	// TEST2: nil values return empty string
	assert.Equal(t, "", ConvertString(nil), "Should return an empty string")
}
