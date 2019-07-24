// Package correlation creates files for visualizing LBA as a correlation network.
package correlation

import (
	"log"

	"github.com/knightjdr/cmgo/internal/pkg/read/list"
	readLocalization "github.com/knightjdr/cmgo/internal/pkg/read/localization"
	"github.com/knightjdr/cmgo/internal/pkg/read/matrix"
	"github.com/knightjdr/cmgo/pkg/correlation"
)

// Network creates a Cytoscape and txt of node pairs for visualizing localization
// profiles as a network.
func Network(fileOptions map[string]interface{}) {
	options, err := parseFlags(fileOptions)
	if err != nil {
		log.Fatalln(err)
	}

	colors := list.ParseSlice(options.colorList)
	possibleLocalizations := list.ParseSlice(options.localizations)
	nodeLocalizations := readLocalization.Prey(options.nodeLocalizations)
	genes, _, profiles := matrix.Read(options.nodeProfiles)

	corr := correlation.CoefficientMatrix(profiles, true, "Pearson")
	cutoff := options.cutoff
	if cutoff == 0 {
		cutoff = calculateCutoff(corr, options.edgesPerNode)
	}
	pairs := filterPairs(corr, genes, cutoff, options.maxEdges)
	writeJSON(genes, pairs, nodeLocalizations, possibleLocalizations, colors, options.outFileNetwork)
	writeTXT(genes, pairs, options.outFile)
}
