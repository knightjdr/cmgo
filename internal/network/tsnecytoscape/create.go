// Package tsnecytoscape generates a Cytoscape network from t-SNE coordinates.
package tsnecytoscape

import (
	"log"

	"github.com/knightjdr/cmgo/internal/pkg/read/list"
	readLocalization "github.com/knightjdr/cmgo/internal/pkg/read/localization"
	"github.com/knightjdr/cmgo/internal/pkg/read/tsne"
)

// Create generates a .cyjs file from t-SNE coordinates.
func Create(fileOptions map[string]interface{}) {
	options, err := parseFlags(fileOptions)
	if err != nil {
		log.Fatalln(err)
	}

	colors := list.ParseSlice(options.colorList)
	possibleLocalizations := list.ParseSlice(options.localizations)
	nodeCoordinates := tsne.Coordinates(options.nodeCoordinates)
	nodeLocalizations := readLocalization.Prey(options.nodeLocalizations)

	// Scale coordinates to defined plot width.
	transformation := networkTransformation(nodeCoordinates, options.width)
	writeJSON(nodeCoordinates, nodeLocalizations, possibleLocalizations, colors, transformation, options.outFile)
}
