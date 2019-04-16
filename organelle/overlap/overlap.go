// Package overlap calculates overlap metrics between two lists of proteins
package overlap

import "log"

// Metrics uses a txt file with similarity scores between proteins and
// calculates metrics between and within two lists of proteins
func Metrics(fileOptions map[string]interface{}) {
	options, err := parseFlags(fileOptions)
	if err != nil {
		log.Fatalln(err)
	}

	compartments := readCompartments(options["compartmentFile"])
	similarity := readSimilarity(options["similarityFile"])

	compare(compartments, similarity, options["outFile"])
}
