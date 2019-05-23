// Package nmfsafe checks NMF and SAFE localizations for agreement.
package nmfsafe

import (
	"log"

	"github.com/knightjdr/cmgo/localization"
	"github.com/knightjdr/cmgo/read/geneontology"
	readLocalization "github.com/knightjdr/cmgo/read/localization"
	"github.com/knightjdr/cmgo/read/nmf"
	"github.com/knightjdr/cmgo/read/safe"
)

type localizationInfo struct {
	Concordant bool
	Domain     int
	NMFids     []string
	NMFinSAFE  bool
	NMFknown   bool
	NMFterms   []string
	Rank       int
	SAFEids    []string
	SAFEinNMF  bool
	SAFEknown  bool
	SAFEterms  []string
}

// Concordance checks NMF and SAFE localizations for agreement.
func Concordance(fileOptions map[string]interface{}) {
	options, err := parseFlags(fileOptions)
	if err != nil {
		log.Fatalln(err)
	}

	// Read GO files.
	goAnnotations := geneontology.Annotations(options.goAnnotations)
	goHierarchy := geneontology.OBO(options.goHierarchy)
	goHierarchy.GetChildren(options.namespace)
	goHierarchy.GetParents(options.namespace)

	// Read and assess NMF localizations.
	nmfLocalizations := nmf.Localization(options.nmfLocalization)
	nmfSummary := readLocalization.SummaryFile(options.nmfSummary)

	genes := make(map[string]*localizationInfo, len(nmfLocalizations))

	for gene, nmfPrediction := range nmfLocalizations {
		genes[gene] = &localizationInfo{
			NMFids:   nmfSummary[nmfPrediction.Compartment].GOid,
			NMFknown: localization.IsKnown(gene, nmfSummary[nmfPrediction.Compartment].GOid, (*goAnnotations.Genes)[options.namespace], (*goHierarchy)[options.namespace]),
			NMFterms: nmfSummary[nmfPrediction.Compartment].GOterms,
			Rank:     nmfPrediction.Compartment,
		}
	}

	// Read and assess SAFE localizations.
	safeLocalizations := safe.Localization(options.safeLocalization)
	safeSummary := readLocalization.SummaryFile(options.safeSummary)

	for gene, safePrediction := range safeLocalizations {
		genes[gene].Domain = safePrediction.Compartment
		genes[gene].SAFEids = safeSummary[safePrediction.Compartment].GOid
		genes[gene].SAFEknown = localization.IsKnown(gene, safeSummary[safePrediction.Compartment].GOid, (*goAnnotations.Genes)[options.namespace], (*goHierarchy)[options.namespace])
		genes[gene].SAFEterms = safeSummary[safePrediction.Compartment].GOterms
	}

	compare(genes, (*goHierarchy)[options.namespace])

	outputTable(genes, options.outFile)
	summarize(genes, options.outSummaryFile)
}
