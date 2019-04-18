// Package cluster clusters a matrix
package cluster

import (
	"log"

	"github.com/knightjdr/hclust"
)

// Process receives a matrix and parameters and returns a cluster matrix
func Process(matrix [][]float64, columns, rows []string, metric, method string) ([][]float64, []string, []string) {
	colDist := hclust.Distance(matrix, metric, true)
	rowDist := hclust.Distance(matrix, metric, false)

	colClust, err := hclust.Cluster(colDist, method)
	if err != nil {
		log.Fatalln(err)
	}
	rowClust, err := hclust.Cluster(rowDist, method)
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

	return sortedAbundance, colTree.Order, rowTree.Order
}
