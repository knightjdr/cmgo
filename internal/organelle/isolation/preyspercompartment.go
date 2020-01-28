package isolation

import (
	"github.com/knightjdr/cmgo/internal/pkg/read/nmf"
	"github.com/knightjdr/cmgo/pkg/slice"
)

func getPreysPerCompartment(localizations nmf.NMFlocalization, genes []string) map[int]map[int]bool {
	preysPerCompartment := make(map[int]map[int]bool, 0)

	for gene, localization := range localizations {
		geneIndex := slice.IndexOfString(gene, genes)
		if _, ok := preysPerCompartment[localization.Compartment]; !ok {
			preysPerCompartment[localization.Compartment] = make(map[int]bool)
		}

		preysPerCompartment[localization.Compartment][geneIndex] = true
	}

	return preysPerCompartment
}
