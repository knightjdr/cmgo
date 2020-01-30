// Package rankmetrics calculates prey metrics for each prey interaction rank.
package rankmetrics

import (
	"log"

	"github.com/knightjdr/cmgo/internal/pkg/read/saint"
)

// Calculate prey metrics for each prey interaction rank,
// such as the turnover rate, cellular abundance and number of lysines.
func Calculate(fileOptions map[string]interface{}) {
	options, err := parseFlags(fileOptions)
	if err != nil {
		log.Fatalln(err)
	}

	saint := saint.Read(options.saint, 1, 0)
	saint.LengthNormalizeSpectralCounts()
	saint.FilterByFDR(options.fdr)

	lysines := countLysines(options.fasta)
	turnoverRates := readTurnoverRates(options.turnoverFile)

	sortedPreysPerBait := saint.SortByPreyRank("NormalizedSpec")
	summary := summarizeMetrics(sortedPreysPerBait, lysines, turnoverRates)
	writeSummary(summary, options.outFile)
}
