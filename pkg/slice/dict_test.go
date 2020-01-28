package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDict(t *testing.T) {
	wanted := map[string]bool{
		"a": true,
		"b": true,
		"c": true,
	}
	dict := Dict([]string{"a", "b", "c"})
	assert.Equal(t, wanted, dict, "Should convert slice to hash")
}

func TestDictInt(t *testing.T) {
	wanted := map[int]bool{
		1: true,
		2: true,
		4: true,
	}
	dict := DictInt([]int{1, 2, 4})
	assert.Equal(t, wanted, dict, "Should convert slice to hash")
}
