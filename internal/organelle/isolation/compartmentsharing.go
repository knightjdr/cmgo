package isolation

import "github.com/knightjdr/cmgo/internal/pkg/read/nmf"

func calculateCompartmentSharing(scores *isolationScores, genes []string, localizations nmf.NMFlocalization) {
	for _, score := range *scores {
		for _, node := range score.nodesOutside {
			targetGene := genes[node]
			targetCompartment := localizations[targetGene].Compartment
			(*score).sharedCompartments[targetCompartment]++
		}
	}
}
