// Package turnoverbyrank calculates the proportion of known interactors for the Nth best prey across
// baits in a SAINT report with statistics on turnover rate
package turnoverbyrank

import (
	"log"

	"github.com/knightjdr/cmgo/internal/pkg/read/saint"
)

// CalculateTurnoverByRank reads a SAINT report and calculates the proportion of known interactors
// per prey rank and the rank turnover rate.
func CalculateTurnoverByRank(fileOptions map[string]interface{}) {
	options, err := parseFlags(fileOptions)
	if err != nil {
		log.Fatalln(err)
	}

	saint := saint.Read(options.saint, 1, 0)
	saint.LengthNormalizeSpectralCounts()
	saint.FilterByFDR(options.fdr)

	sortedPreysPerBait := saint.SortByPreyRank("NormalizedSpec")
	turnoverRates := readTurnoverRates(options.turnoverFile)

	summary := summarizeInteractions(sortedPreysPerBait, turnoverRates)
	writeSummary(summary, options.outFile)
}
