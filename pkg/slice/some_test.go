package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSome(t *testing.T) {
	slice := []float64{0.01, 0.2, 0.05}
	testThresholds := []float64{0.15, 0.3}

	wanted := []bool{true, false}
	for i, threshold := range testThresholds {
		testFunc := generateTestFunc(threshold)
		assert.Equal(t, wanted[i], SomeFloat(slice, testFunc), "Should return boolean indicating if slice member passes testing function")
	}
}

func generateTestFunc(threshold float64) func(value float64) bool {
	return func(value float64) bool {
		if value >= threshold {
			return true
		}
		return false
	}
}
