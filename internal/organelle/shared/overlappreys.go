package shared

import (
	"github.com/knightjdr/cmgo/internal/organelle"
	"github.com/knightjdr/cmgo/internal/pkg/read/saint"
	"github.com/knightjdr/cmgo/pkg/filter"
	"github.com/knightjdr/cmgo/pkg/slice"
	"github.com/knightjdr/cmgo/pkg/stats"
)

func overlapPreys(compartments organelle.Compartments, saint *saint.SAINT, minOccurence int) []string {
	compartmentDictA := slice.Dict(compartments[0].Proteins)
	compartmentDictB := slice.Dict(compartments[1].Proteins)

	// Get all preys for each compartment (including duplicates).
	preysCompartmentA := make([]string, 0)
	preysCompartmentB := make([]string, 0)
	for _, row := range *saint {
		if _, ok := compartmentDictA[row.Bait]; ok {
			preysCompartmentA = append(preysCompartmentA, row.PreyGene)
		} else if _, ok := compartmentDictB[row.Bait]; ok {
			preysCompartmentB = append(preysCompartmentB, row.PreyGene)
		}
	}

	// Count prey occurrences.
	preyOccurrencesA := stats.Occurrence(preysCompartmentA)
	preyOccurrencesB := stats.Occurrence(preysCompartmentB)

	// Keey preys passing minOccurence cutoff
	preysCompartmentA = filter.Key(preyOccurrencesA, minOccurence)
	preysCompartmentB = filter.Key(preyOccurrencesB, minOccurence)

	intersection := slice.Intersect(preysCompartmentA, preysCompartmentB)

	return intersection
}
