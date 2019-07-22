// Package evaluate localizes preys using LBA
package evaluate

import (
	"log"

	"github.com/knightjdr/cmgo/internal/pkg/localization"
	"github.com/knightjdr/cmgo/internal/pkg/read/geneontology"
	readLocalization "github.com/knightjdr/cmgo/internal/pkg/read/localization"
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

	genes := make(map[string]map[string]map[string]interface{}, len(localizations))
	for gene, prediction := range localizations {
		genes[gene] = make(map[string]map[string]interface{}, 0)
		for id, name := range prediction {
			genes[gene][id] = map[string]interface{}{
				"known": localization.IsKnown(gene, []string{id}, (*goAnnotations.Genes)[options.namespace], (*goHierarchy)[options.namespace]),
				"name":  name,
			}
		}

	}

	knownList(genes, options.outFile)
	summarize(genes, options.outFileSummary)
}
