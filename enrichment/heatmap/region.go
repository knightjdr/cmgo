// Package heatmap creates a heat map to display enriched regions across compartments.
package heatmap

import (
	"log"

	"github.com/knightjdr/cmgo/image/svg"
	"github.com/knightjdr/hclust"
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

	parameters := svg.HeatmapSettings{
		Filename:     options.outFile,
		FillColor:    "blueBlack",
		AbundanceCap: 50,
		MinAbundance: 0,
		InvertColor:  false,
	}
	svg.Heatmap(sortedAbundance, colTree.Order, rowTree.Order, parameters)
}
