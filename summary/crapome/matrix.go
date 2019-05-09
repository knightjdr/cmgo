// Package crapome generates a CRAPome matrix from SAINT input files.
package crapome

import (
	"log"

	"github.com/knightjdr/cmgo/read/saint"
)

// Matrix reads SAINT input files and generates a CRAPome matrix.
func Matrix(fileOptions map[string]interface{}) {
	options, err := parseFlags(fileOptions)
	if err != nil {
		log.Fatalln(err)
	}

	// Read and merge .dat files.
	baits := make([]saint.BaitDatRow, 0)
	for _, filename := range options.baitFiles {
		baits = append(baits, saint.BaitDat(filename)...)
	}
	interactions := make([]saint.InterDatRow, 0)
	for _, filename := range options.interactionFiles {
		interactions = append(interactions, saint.InterDat(filename)...)
	}
	preys := make([]saint.PreyDatRow, 0)
	for _, filename := range options.preyFiles {
		preys = append(preys, saint.PreyDat(filename)...)
	}

	// Remove non-controls.
	baits, interactions = removeNonControls(baits, interactions)

	// Create map for prey accessions.
	preyMap := preyDict(preys)

	// Parse interactions and determine preys to output and order for output.
	parsed := parseInteractions(interactions)
	preyOrder := orderPreys(parsed, preyMap)

	writeMatrix(parsed, baits, preyMap, preyOrder, options.outFile)
}
