// Package knownbyrank calculates the proportion of known interactors for the Nth best prey across
// baits in a SAINT report
package knownbyrank

import (
	"log"

	"github.com/knightjdr/cmgo/internal/pkg/read/interactions"
	"github.com/knightjdr/cmgo/internal/pkg/read/saint"
)

// CalculateKnownByRank reads a SAINT report and calculates the proportion of known interactors
// per prey rank.
func CalculateKnownByRank(fileOptions map[string]interface{}) {
	options, err := parseFlags(fileOptions)
	if err != nil {
		log.Fatalln(err)
	}

	saint := saint.Read(options.saint, 1, 0)
	saint.LengthNormalizeSpectralCounts()
	saint.FilterByFDR(options.fdr)

	sortedPreysPerBait := saint.SortByPreyRank("NormalizedSpec")
	interactors := interactions.Read(options.biogrid, options.intact, options.species)

}
