package shared

import (
	"sort"
	"testing"

	"github.com/knightjdr/cmgo/organelle"
	"github.com/knightjdr/cmgo/read/saint"
	"github.com/stretchr/testify/assert"
)

func TestOverlapPreys(t *testing.T) {
	compartments := organelle.Compartments{
		{Name: "compartmentA", Proteins: []string{"a", "b", "c"}},
		{Name: "compartmentB", Proteins: []string{"d", "e", "f"}},
	}
	saint := []saint.Row{
		{Bait: "a", PreyGene: "1"},
		{Bait: "a", PreyGene: "2"},
		{Bait: "a", PreyGene: "3"},
		{Bait: "b", PreyGene: "1"},
		{Bait: "b", PreyGene: "4"},
		{Bait: "c", PreyGene: "5"},
		{Bait: "c", PreyGene: "6"},
		{Bait: "d", PreyGene: "1"},
		{Bait: "d", PreyGene: "7"},
		{Bait: "d", PreyGene: "3"},
		{Bait: "e", PreyGene: "6"},
		{Bait: "f", PreyGene: "1"},
		{Bait: "f", PreyGene: "2"},
	}

	// TEST1: prey only needs to be seen with one bait in each compartment
	wanted := []string{"1", "2", "3", "6"}
	overlap := overlapPreys(compartments, saint, 1)
	sort.Strings(overlap)
	assert.Equal(t, wanted, overlap, "Should return a slice of the overlapping preys between two compartments")

	// TEST2: prey needs to be seen with at least two baits in each compartment
	wanted = []string{"1"}
	overlap = overlapPreys(compartments, saint, 2)
	sort.Strings(overlap)
	assert.Equal(t, wanted, overlap, "Should return a slice of the overlapping preys seen with two baits between two compartments")
}
