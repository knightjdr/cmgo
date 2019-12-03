// Package recovered reports the number of genes seen in a specific compartment.
package recovered

import (
	"log"

	"github.com/knightjdr/cmgo/internal/pkg/read/geneontology"
)

// AssessCompartment reports the number of genes seen from a compartment.
func AssessCompartment(fileOptions map[string]interface{}) {
	options, err := parseFlags(fileOptions)
	if err != nil {
		log.Fatalln(err)
	}

	annotations := geneontology.Annotations(options.goAnnotations)
	compartmentGenes := readCompartmentGenes((*annotations.Genes)["CC"], options.compartmentID)
	localizedGenes := readLocalizedGenes(options.genes, options.localizationID)

	summary := countRecoveredGenes(localizedGenes, compartmentGenes)
	writeRecovered(summary, options.outFile)
}
