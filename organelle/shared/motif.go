// Package shared outputs metrics on the regions found in preys shared by two lists of proteins
package shared

import (
	"log"

	"github.com/knightjdr/cmgo/organelle"
	"github.com/knightjdr/cmgo/read"
)

// Motif calculates shared motifs between two lists of proteins
func Motif(fileOptions map[string]interface{}) {
	options, err := parseFlags(fileOptions)
	if err != nil {
		log.Fatalln(err)
	}

	compartments := organelle.ReadCompartments(options.compartmentFile)
	saint := read.Saint(options.saintFile, options.fdr)

	regions := readRegions(options.regionFile)
}
