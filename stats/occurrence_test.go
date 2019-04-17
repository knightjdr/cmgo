package stats

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOccurence(t *testing.T) {
	slice := []string{"a", "a", "b", "c", "c", "d", "c", "d", "d", "d"}
	wanted := map[string]int{
		"a": 2,
		"b": 1,
		"c": 3,
		"d": 4,
	}
	assert.Equal(t, wanted, Occurrence(slice), "Should return a map with the number of times each key occurs in the input slice")
}
