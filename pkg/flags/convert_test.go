package flags

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertFloat(t *testing.T) {
	// TEST1: return an input float
	assert.Equal(t, 0.01, convertFloat(0.01), "Should return input float")

	// TEST2: convert an int to a float
	assert.Equal(t, float64(1), convertFloat(1), "Should convert an int to a float")

	// TEST3: convert a string to a float
	assert.Equal(t, 0.01, convertFloat("0.01"), "Should convert a string to a float")

	// TEST4: unhandled type
	assert.Equal(t, float64(0), convertFloat(int64(1)), "Should return 0 for unrecognized type")

	// TEST5: nil values return nil float value
	assert.Equal(t, float64(0), convertFloat(nil), "Should return 0")
}

func TestConvertInt(t *testing.T) {
	// TEST1: return an input integer
	assert.Equal(t, 1, convertInt(1), "Should return input integer")

	// TEST2: convert a float to an int
	assert.Equal(t, 1, convertInt(1.4), "Should convert a float to an int")

	// TEST3: convert a string to a int
	assert.Equal(t, 1, convertInt("1"), "Should convert a string to a int")

	// TEST4: unhandled type
	assert.Equal(t, 0, convertInt(int64(1)), "Should return 0 for unrecognized type")

	// TEST5: nil values return nil int value
	assert.Equal(t, 0, convertInt(nil), "Should return 0")
}

func TestConvertString(t *testing.T) {
	// TEST1: convert a value to a string
	assert.Equal(t, "a", convertString("a"), "Should convert an interface to a string")

	// TEST2: nil values return empty string
	assert.Equal(t, "", convertString(nil), "Should return an empty string")
}
