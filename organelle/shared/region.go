// Package shared outputs metrics on the regions found in preys shared by two lists of proteins
package shared

import (
	"log"

	"github.com/knightjdr/cmgo/organelle"
	"github.com/knightjdr/cmgo/read"
)

// Region summarizes regions in preys shared between two lists of proteins
func Region(fileOptions map[string]interface{}) {
	options, err := parseFlags(fileOptions)
	if err != nil {
		log.Fatalln(err)
	}

	compartments := organelle.ReadCompartments(options.compartmentFile)
	regions := readRegions(options.regionFile)
	saint := read.Saint(options.saintFile, options.fdr, 1)

	shared := overlapPreys(compartments, saint, options.minPreyOccurrence)

	summarizeRegions(shared, regions, options.outFile)
}
