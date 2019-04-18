// Package heatmap creates a heat map to display enriched regions across compartments.
package heatmap

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/knightjdr/cmgo/fs"
	"github.com/knightjdr/cmgo/image/svg/heatmap"
	"github.com/knightjdr/cmgo/image/svg/legend"
	"github.com/knightjdr/hclust"
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

	colDist := hclust.Distance(matrix, options.distanceMetric, true)
	rowDist := hclust.Distance(matrix, options.distanceMetric, false)

	colClust, err := hclust.Cluster(colDist, options.clusteringMethod)
	if err != nil {
		log.Fatalln(err)
	}
	rowClust, err := hclust.Cluster(rowDist, options.clusteringMethod)
	if err != nil {
		log.Fatalln(err)
	}

	// Optimize clustering.
	colClust = hclust.Optimize(colClust, colDist, 0)
	rowClust = hclust.Optimize(rowClust, rowDist, 0)

	// Create tree and get clustering order.
	colTree, err := hclust.Tree(colClust, columns)
	rowTree, err := hclust.Tree(rowClust, rows)

	// Sort matrices.
	sortedAbundance, _ := hclust.Sort(matrix, columns, colTree.Order, "column")
	sortedAbundance, _ = hclust.Sort(sortedAbundance, rows, rowTree.Order, "row")

	parameters := heatmap.Settings{
		AbundanceCap: options.abundanceCap,
		Filename:     options.outFile,
		FillColor:    "blueBlack",
		InvertColor:  false,
		MinAbundance: options.minAbundance,
	}
	heatmap.Draw(sortedAbundance, colTree.Order, rowTree.Order, parameters)

	// legend
	dir := filepath.Dir(options.outFile)
	outFile := filepath.Base(options.outFile)

	legendTitle := fmt.Sprintf("Fold enrichment (log2) - %s", outFile)
	distanceLegend := legend.Gradient("blueBlack", legendTitle, 101, options.minAbundance, options.abundanceCap, false)

	legendFileName := fmt.Sprintf("%s/legend-%s", dir, outFile)
	afero.WriteFile(fs.Instance, legendFileName, []byte(distanceLegend), 0644)
}
