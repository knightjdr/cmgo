package uv

import (
	readNMF "github.com/knightjdr/cmgo/internal/pkg/read/nmf"
	"github.com/knightjdr/cmgo/pkg/slice"
)

func defineNonCharacterizingGenes(characterizingGenes [][]string, nmfLocalizations readNMF.NMFlocalization) [][]string {
	// Create a map of all characterizing genes for each rank for
	// quickly checking for existence.
	dict := make([]map[string]bool, len(characterizingGenes))
	for i, genes := range characterizingGenes {
		dict[i] = make(map[string]bool, 0)
		dict[i] = slice.Dict(genes)
	}

	ncg := make([][]string, len(characterizingGenes))
	for gene, details := range nmfLocalizations {
		rank := details.Compartment
		rankIndex := rank - 1
		if _, ok := dict[rankIndex][gene]; !ok {
			if len(ncg[rankIndex]) == 0 {
				ncg[rankIndex] = make([]string, 0)
			}
			ncg[rankIndex] = append(ncg[rankIndex], gene)
		}
	}
	return ncg
}
