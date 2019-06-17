// Package robustness assesses the sensitivity of NMF rank assignments
package robustness

import (
	"log"

	"github.com/knightjdr/cmgo/internal/pkg/read/nmf"
)

// Evaluate performs GO enrichments on each NMF rank and tests how
// sensitive these are to the genes used.
func Evaluate(fileOptions map[string]interface{}) {
	options, err := parseFlags(fileOptions)
	if err != nil {
		log.Fatalln(err)
	}

	basis, columns, rows := nmf.Basis(options.basisMatrix)
	characterizingGenes := filterBasis(basis, options.maxGenesPerRank, options.minRankValue, options.withinRankMax)
	rankDefinitions := defineRanks(characterizingGenes, rows)

	dataPoints := make([][][]float64, len(columns))
	for rank, geneIndices := range characterizingGenes {
		dataPoints[rank] = dataPoint(geneIndices, rows, rankDefinitions[rank], options.percentiles, options.persistence, options.replicates)
	}
	writeData(dataPoints, columns, options.percentiles, options.replicates, options.outFile)
}
