package gradient

import (
	"testing"

	"github.com/knightjdr/cmgo/read/localization"
	"github.com/stretchr/testify/assert"
)

func TestAdjacentSimilarity(t *testing.T) {
	baits := []string{"a", "b", "c", "d", "e", "f"}
	expected := localization.ExpectedLocalizations{
		"a": {Terms: []string{"localization1"}},
		"b": {Terms: []string{"localization1"}},
		"c": {Terms: []string{"localization2"}},
		"d": {Terms: []string{"localization2", "localization3"}},
		"e": {Terms: []string{"localization3"}},
		"f": {Terms: []string{"localization4"}},
	}
	wanted := []int{1, 1, 1, 2, 1, 0}
	assert.Equal(t, wanted, adjacentSimilarity(baits, expected), "Should calculate adjacent similarity for each bait")
}
