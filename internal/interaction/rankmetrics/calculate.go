// Package rankmetrics calculates prey metrics for each prey interaction rank.
package rankmetrics

import (
	"log"

	"github.com/knightjdr/cmgo/internal/pkg/read/saint"
)

type analysis struct {
	parameters     parameters
	uniprotMapping map[string]string
}

// Calculate prey metrics for each prey interaction rank,
// such as the turnover rate, cellular abundance and number of lysines.
func Calculate(fileOptions map[string]interface{}) {
	data := analysis{}
	var err error

	data.parameters, err = parseFlags(fileOptions)
	if err != nil {
		log.Fatalln(err)
	}

	saint := saint.Read(data.parameters.saint, 1, 0)
	saint.LengthNormalizeSpectralCounts()
	saint.FilterByFDR(data.parameters.fdr)

	data.uniprotMapping = fetchUniprotIDs(saint)

	/* lysines := countLysines(data.parameters.fasta)
	turnoverRates := readTurnoverRates(data.parameters.turnoverFile)

	sortedPreysPerBait := saint.SortByPreyRank("NormalizedSpec")
	summary := summarizeMetrics(sortedPreysPerBait, lysines, turnoverRates)
	writeSummary(summary, data.parameters.outFile) */
}
