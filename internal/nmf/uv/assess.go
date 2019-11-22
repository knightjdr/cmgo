// Package uv validates the localizations of prey genes not used for defining NMF ranks
package uv

import (
	"log"

	"github.com/knightjdr/cmgo/internal/pkg/nmf"
	"github.com/knightjdr/cmgo/internal/pkg/read/geneontology"
	readLocalization "github.com/knightjdr/cmgo/internal/pkg/read/localization"
	readNMF "github.com/knightjdr/cmgo/internal/pkg/read/nmf"
)

// Assess the localizations of preys that do not characterize
// an NMF rank.
func Assess(fileOptions map[string]interface{}) {
	options, err := parseFlags(fileOptions)
	if err != nil {
		log.Fatalln(err)
	}

	// Define genes that could characterize an NMF rank but fall outside the
	// options.maxGenesPerRank level.
	nmfLocalizations := readNMF.Localization(options.nmfLocalization)
	basis, _, rows := readNMF.ReadBasis(options.basisMatrix)
	characterizingGeneIndices := nmf.FilterBasisForRankDefiningPreys(basis, options.maxGenesPerRank, options.minRankValue, options.withinRankMax)
	characterizingGenes := geneFromIndex(characterizingGeneIndices, rows)
	nonCharacterizingGenes := defineNonCharacterizingGenes(characterizingGenes, nmfLocalizations)

	// Read GO files.
	goAnnotations := geneontology.Annotations(options.goAnnotations)
	goHierarchy := geneontology.OBO(options.goHierarchy)
	goHierarchy.GetChildren(options.namespace)
	goHierarchy.GetParents(options.namespace)

	nmfSummary := readLocalization.SummaryFile(options.nmfSummary)

	assessment := assessGenes(nonCharacterizingGenes, nmfSummary, (*goAnnotations.Genes)[options.namespace], (*goHierarchy)[options.namespace])
	writeSummary(assessment, options.outFile)
}
