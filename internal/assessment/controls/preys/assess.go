// Package preys outputs statistics on prey proteins in controls.
package preys

import (
	"log"
)

type analysis struct {
	parameters parameters
}

// Assess prey metrics for each prey interaction rank,
// such as the turnover rate, cellular abundance and number of lysines.
func Assess(fileOptions map[string]interface{}) {
	data := analysis{}
	var err error

	data.parameters, err = parseFlags(fileOptions)
	if err != nil {
		log.Fatalln(err)
	}

	baits, interactions := readSaintFiles(data.parameters)
	summary := summarize(baits, interactions)

	writeSummary(summary, data.parameters.outFile)
	goEnrichment(summary, data.parameters)
}
