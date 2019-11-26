package transmembrane

import (
	"github.com/knightjdr/cmgo/internal/pkg/read/saint"
	"github.com/knightjdr/cmgo/pkg/slice"
)

func findBaitsPerPrey(preys []string, saint *saint.SAINT) map[string]map[string]bool {
	preyDict := slice.Dict(preys)

	baitsPerPrey := make(map[string]map[string]bool, 0)
	for _, row := range *saint {
		if _, ok := preyDict[row.PreyGene]; ok {
			allocateMemoryBaitsPerPrey(baitsPerPrey, row.PreyGene)
			baitsPerPrey[row.PreyGene][row.Bait] = true
		}
	}

	return baitsPerPrey
}

func allocateMemoryBaitsPerPrey(baitsPerPrey map[string]map[string]bool, key string) {
	if _, ok := baitsPerPrey[key]; !ok {
		baitsPerPrey[key] = make(map[string]bool, 0)
	}
}
