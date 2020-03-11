// Package goenrich finds enriched GO terms via gProfiler
package goenrich

import (
	"log"

	"github.com/knightjdr/cmgo/internal/pkg/read/saint"
)

// Enrich finds enriched GO terms via gProfiler
func Enrich(fileOptions map[string]interface{}) {
	options, err := parseFlags(fileOptions)
	if err != nil {
		log.Fatalln(err)
	}

	saint := saint.Read(options.saint, 0.01, 0)
	enrichment := enrichWithGprofiler(saint, options)
	writeEnrichment(enrichment, options)
}
