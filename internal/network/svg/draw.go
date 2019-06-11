// Package svg creates a network in svg format from NMF or SAFE data.
package svg

import (
	"log"

	"github.com/knightjdr/cmgo/internal/pkg/read/list"
	"github.com/knightjdr/cmgo/internal/pkg/read/nmf"
)

// Draw creates an svg version of a network.
func Draw(fileOptions map[string]interface{}) {
	options, err := parseFlags(fileOptions)
	if err != nil {
		log.Fatalln(err)
	}

	colors := list.ParseSlice(options.colorList)
	localization := nmf.Localization(options.localizations)
	coordinates := readCoordinates(options.nodeCoordinates)

	// Scale coordinates to defined plot width.
	plotWidth := float64(1000)
	coordinates, plotHeight := scaleCoordinates(coordinates, plotWidth, 4, 1)

	writeSVG(coordinates, colors, localization, plotWidth, plotHeight, options.outFile)
}
