package customsort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestByMapValueInt(t *testing.T) {
	m := map[string]int{
		"a": 5,
		"b": 20,
		"c": 7,
	}

	// TEST1: sort in descending order
	wanted := []KeyValue{
		{Key: "b", Value: 20},
		{Key: "c", Value: 7},
		{Key: "a", Value: 5},
	}
	assert.Equal(t, wanted, ByMapValueInt(m, "descending"), "Should return descending sort order for a map by integer values")

	// TEST1: sort in descending order
	wanted = []KeyValue{
		{Key: "a", Value: 5},
		{Key: "c", Value: 7},
		{Key: "b", Value: 20},
	}
	assert.Equal(t, wanted, ByMapValueInt(m, "ascending"), "Should return ascending sort order for a map by integer values")
}
