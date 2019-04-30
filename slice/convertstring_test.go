package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertStringToFloat(t *testing.T) {
	slice := []string{"1", "2", "2", "7"}
	wanted := []float64{1, 2, 2, 7}
	assert.Equal(t, wanted, ConvertStringToFloat(slice), "Should convert slice of strings to floats")
}

func TestConvertStringToInt(t *testing.T) {
	slice := []string{"1", "2", "2", "7"}
	wanted := []int{1, 2, 2, 7}
	assert.Equal(t, wanted, ConvertStringToInt(slice), "Should convert slice of strings to ints")
}
