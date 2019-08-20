// Package genes performs a GO enrichment on a list of genes.
package genes

import (
	"log"

	"github.com/knightjdr/cmgo/internal/pkg/read/list"
)

// Enrich performs a gProfiler enrichment on a list of genes.
func Enrich(fileOptions map[string]interface{}) {
	options, err := parseFlags(fileOptions)
	if err != nil {
		log.Fatalln(err)
	}

	background := list.ParseSlice(options.background)
	genes := list.ParseSlice(options.genes)

	enrichment := profile(genes, background, options.namespace)
	write(enrichment, options.outFile)
}
