package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortStringsCaseInsensitive(t *testing.T) {
	unsorted := []string{"b", "A", "C", "e", "F", "d"}
	wanted := []string{"A", "b", "C", "d", "e", "F"}
	assert.Equal(t, wanted, SortStringsCaseInsensitive(unsorted), "Should sort slice of strings in case insensitive fashion")
}
