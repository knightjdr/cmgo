// Package matrix creates files for visualizing LBA as a correlation network.
package matrix

import (
	"log"

	"github.com/knightjdr/cmgo/internal/pkg/read/list"
	readLocalization "github.com/knightjdr/cmgo/internal/pkg/read/localization"
	"github.com/knightjdr/cmgo/internal/pkg/read/matrix"
)

// Network creates a Cytoscape and txt of node pairs for visualizing a
// matrix as a network.
func Network(fileOptions map[string]interface{}) {
	options, err := parseFlags(fileOptions)
	if err != nil {
		log.Fatalln(err)
	}

	colors := list.ParseSlice(options.colorList)
	possibleLocalizations := list.ParseSlice(options.localizations)
	nodeLocalizations := readLocalization.Prey(options.nodeLocalizations)
	genes, _, profiles := matrix.Read(options.matrix)

	writeJSON(profiles, genes, options.cutoff, nodeLocalizations, possibleLocalizations, colors, options.outFileNetwork)
}
