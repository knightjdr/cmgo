// Package heatmap creates a heat map to display enriched regions across compartments.
package heatmap

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/knightjdr/cmgo/internal/pkg/cluster"
	"github.com/knightjdr/cmgo/internal/pkg/image/svg/heatmap"
	"github.com/knightjdr/cmgo/internal/pkg/image/svg/legend"
	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/spf13/afero"
)

// Region a heat map from enrichment data.
func Region(fileOptions map[string]interface{}) {
	options, err := parseFlags(fileOptions)
	if err != nil {
		log.Fatalln(err)
	}

	enrichment := readEnrichment(options.enrichmentFile, options.pValue)
	matrix, columns, rows := enrichmentMatrix(enrichment)

	// Clustering
	matrix, columns, rows = cluster.Process(matrix, columns, rows, options.distanceMetric, options.clusteringMethod)

	parameters := heatmap.Settings{
		AbundanceCap: options.abundanceCap,
		Filename:     options.outFile,
		FillColor:    "blueBlack",
		InvertColor:  false,
		MinAbundance: options.minAbundance,
	}
	heatmap.Draw(matrix, columns, rows, parameters)

	// legend
	dir := filepath.Dir(options.outFile)
	outFile := filepath.Base(options.outFile)

	legendTitle := fmt.Sprintf("Fold enrichment (log2) - %s", outFile)
	distanceLegend := legend.Gradient("blueBlack", legendTitle, 101, options.minAbundance, options.abundanceCap, false)

	legendFileName := fmt.Sprintf("%s/legend-%s", dir, outFile)
	afero.WriteFile(fs.Instance, legendFileName, []byte(distanceLegend), 0644)
}
