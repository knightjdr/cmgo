// Package evaluate localizes preys using LBA
package evaluate

import (
	"log"

	"github.com/knightjdr/cmgo/internal/pkg/localization"
	"github.com/knightjdr/cmgo/internal/pkg/read/geneontology"
	readLocalization "github.com/knightjdr/cmgo/internal/pkg/read/localization"
	"github.com/knightjdr/cmgo/pkg/mapfunc"
)

// Evaluate determines whether localizations are previously known
func Evaluate(fileOptions map[string]interface{}) {
	options, err := parseFlags(fileOptions)
	if err != nil {
		log.Fatalln(err)
	}

	// Read GO files.
	goAnnotations := geneontology.Annotations(options.goAnnotations)
	goHierarchy := geneontology.OBO(options.goHierarchy)
	goHierarchy.GetChildren(options.namespace)
	goHierarchy.GetParents(options.namespace)

	// Read and assess localizations.
	localizations := readLocalization.Prey(options.localization)

	genes := make(map[string]bool, len(localizations))
	for gene, prediction := range localizations {
		ids := mapfunc.KeysStringString(prediction)
		genes[gene] = localization.IsKnown(gene, ids, (*goAnnotations.Genes)[options.namespace], (*goHierarchy)[options.namespace])
	}

	summarize(genes, options.outFileSummary)
}
