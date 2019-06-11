// Package overlap calculates overlap metrics between two lists of proteins
package overlap

import (
	"log"

	"github.com/knightjdr/cmgo/internal/organelle"
)

// Metrics uses a txt file with similarity scores between proteins and
// calculates metrics between and within two lists of proteins
func Metrics(fileOptions map[string]interface{}) {
	options, err := parseFlags(fileOptions)
	if err != nil {
		log.Fatalln(err)
	}

	compartments := organelle.ReadCompartments(options["compartmentFile"])
	similarity := readSimilarity(options["similarityFile"])

	compare(compartments, similarity, options["outFile"])
}
