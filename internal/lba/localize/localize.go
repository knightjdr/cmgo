// Package localize localizes preys using LBA
package localize

import (
	"log"
)

// Localize assigns genes to a compartment and generates a localization
// profile from enriched terms and a list of valid localizations.
func Localize(fileOptions map[string]interface{}) {
	options, err := parseFlags(fileOptions)
	if err != nil {
		log.Fatalln(err)
	}

	enrichment := readEnrichment(options.enrichment)
	goIDs := readLocalizations(options.localization)

	// Remove invalid localizations.
	enrichment = filterEnrichment(enrichment, goIDs)

	writePrimary(enrichment, goIDs, options.outFilePrimary)
}
