// Package subset gets a cluster from a prey-prey interactive file.
package subset

import (
	"log"

	"github.com/knightjdr/cmgo/internal/pkg/read/list"
)

// Heatmap takes a prey-prey interactive file for ProHits-viz and
// subsets a cluster based on an input list of genes.
func Heatmap(fileOptions map[string]interface{}) {
	options, err := parseFlags(fileOptions)
	if err != nil {
		log.Fatalln(err)
	}

	genes := list.ParseSlice(options.genes)
	lines, params := parseHeatmap(options.heatmap, genes)

	write(lines, params, options.outFile)
}
