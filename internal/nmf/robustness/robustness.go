// Package robustness assesses the sensitivity of NMF rank assignments
package robustness

import (
	"log"

	"github.com/knightjdr/cmgo/internal/pkg/nmf"
	readNMF "github.com/knightjdr/cmgo/internal/pkg/read/nmf"
)

// Evaluate performs GO enrichments on each NMF rank and tests how
// sensitive these are to the genes used.
func Evaluate(fileOptions map[string]interface{}) {
	options, err := parseFlags(fileOptions)
	if err != nil {
		log.Fatalln(err)
	}

	basis, columns, rows := readNMF.Basis(options.basisMatrix)
	characterizingGenes := nmf.FilterBasis(basis, options.maxGenesPerRank, options.minRankValue, options.withinRankMax)
	rankDefinitions := defineRanks(characterizingGenes, rows)

	dataPoints := make([][][]float64, len(columns))
	for rank, geneIndices := range characterizingGenes {
		dataPoints[rank] = dataPoint(geneIndices, rows, rankDefinitions[rank], options.percentiles, options.persistence, options.replicates)
	}
	summary := summaryStats(dataPoints, options.percentiles)

	writeData(dataPoints, columns, options.percentiles, options.replicates, options.outFile)
	writeStats(summary, columns, options.percentiles, options.outFileSummary)
}
