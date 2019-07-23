// Package correlation creates files for visualizing LBA as a correlation network.
package correlation

import (
	"log"

	"github.com/knightjdr/cmgo/internal/pkg/read/list"
	readLocalization "github.com/knightjdr/cmgo/internal/pkg/read/localization"
	"github.com/knightjdr/cmgo/internal/pkg/read/matrix"
	"github.com/knightjdr/cmgo/pkg/correlation"
)

// Network creates a Cytoscape and txt of node pairs for visualizing an LBA
// profile as a network.
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
	cutoff := calculateCutoff(corr, options.edgesPerNode)
	writeJSON(corr, genes, cutoff, nodeLocalizations, possibleLocalizations, colors, options.outFileNetwork)
}
