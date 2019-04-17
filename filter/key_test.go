package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKey(t *testing.T) {
	occurrence := map[string]int{
		"a": 2,
		"b": 1,
		"c": 2,
		"d": 10,
	}
	wanted := []string{"a", "c", "d"}
	assert.Equal(t, wanted, Key(occurrence, 2), "Should return a slice with map keys that have values passing filter")
}
