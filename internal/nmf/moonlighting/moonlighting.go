// Package moonlighting scores preys localized by NMF for moonlighting in more than one compartment
package moonlighting

import (
	"log"

	"github.com/knightjdr/cmgo/internal/pkg/nmf"
	readLocalization "github.com/knightjdr/cmgo/internal/pkg/read/localization"
	"github.com/knightjdr/cmgo/internal/pkg/read/matrix"
	readNMF "github.com/knightjdr/cmgo/internal/pkg/read/nmf"
)

type outputOptions struct {
	localization readLocalization.Summary
	minRankValue float64
	outfile      string
	preyNames    []string
}

// Calculate calculates prey moonlighting scores from an NMF basis file
// and a dissimilarity matrix.
func Calculate(fileOptions map[string]interface{}) {
	options, err := parseFlags(fileOptions)
	if err != nil {
		log.Fatalln(err)
	}

	basis, _, rowNames := readNMF.ReadBasis(options.basisMatrix)
	basis, rowNames = nmf.FilterBasisByTreshold(basis, rowNames, options.minRankValue)
	nmfSummary := readLocalization.SummaryFile(options.nmfSummary)

	_, _, dissimilarityMatrix := matrix.Read(options.dissimilarityMatrix)
	compatibleRanks := defineCompatibleRanks(dissimilarityMatrix)
	moonlightingScores := calculateMoonlightingScores(basis, compatibleRanks)

	outputOpts := outputOptions{
		localization: nmfSummary,
		minRankValue: options.minRankValue,
		outfile:      options.outFileScores,
		preyNames:    rowNames,
	}
	writeMoonlightingScores(moonlightingScores, outputOpts)

	rankMoonlightingMatrix := countRankMoonlighting(moonlightingScores, len(compatibleRanks), options.minRankValue)
	writeRankMoonlightingMatrix(rankMoonlightingMatrix, options.outFileMatrix)
	writeHeatmap(rankMoonlightingMatrix, options.outFileHeatmap)
}
