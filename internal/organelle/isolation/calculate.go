// Package isolation calculates percent isolation of each NMF compartment.
package isolation

import (
	"log"

	"github.com/knightjdr/cmgo/internal/pkg/read/localization"
	"github.com/knightjdr/cmgo/internal/pkg/read/nmf"
	"github.com/knightjdr/cmgo/pkg/correlation"
)

// Calculate the percent isolation of each NMF compartment.
func Calculate(fileOptions map[string]interface{}) {
	options, err := parseFlags(fileOptions)
	if err != nil {
		log.Fatalln(err)
	}

	basis, _, genes := nmf.ReadBasis(options.basisMatrix)
	corr := correlation.CoefficientMatrix(basis, true, "pearson")

	localizations := nmf.Localization(options.nmfLocalization)
	preysPerCompartment := getPreysPerCompartment(localizations, genes)

	scores := calculateIsolation(corr, options.correlationCutoff, preysPerCompartment)
	calculateCompartmentSharing(scores, genes, localizations)

	nmfSummary := localization.SummaryFile(options.nmfSummary)
	writeScores(scores, nmfSummary, options.outFile)
	writeHeatmap(scores, nmfSummary, options)
}
