package interactions

import (
	"sort"

	"github.com/knightjdr/cmgo/pkg/mapfunc"
)

func mergeInteractors(biogrid, intact map[string][]string) map[string][]string {
	interactorsDict := make(map[string]map[string]bool, 0)
	addInteractors(biogrid, &interactorsDict)
	addInteractors(intact, &interactorsDict)

	interactors := make(map[string][]string, 0)
	for source, targets := range interactorsDict {
		interactors[source] = mapfunc.KeysStringBool(targets)
		sort.Strings(interactors[source])
	}
	return interactors
}

func addInteractors(interactorSubset map[string][]string, interactors *map[string]map[string]bool) {
	for source, targets := range interactorSubset {
		if _, ok := (*interactors)[source]; !ok {
			(*interactors)[source] = make(map[string]bool, 0)
		}
		for _, target := range targets {
			(*interactors)[source][target] = true
		}
	}
}
