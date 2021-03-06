package crapome

import (
	"sort"
)

func orderPreys(data map[string]map[string]int, preyMap map[string]string) []string {
	// Order data preys alphabetically by name. For each prey in dataset, get gene name
	// from map. Sort by gene name, then map that back to accession.
	preyOrder := make([]string, 0)
	reverseMap := make(map[string]string, 0)
	for prey, name := range preyMap {
		if _, ok := data[prey]; ok {
			preyOrder = append(preyOrder, name)
			reverseMap[name] = prey
		}
	}
	sort.Strings(preyOrder)

	for i, name := range preyOrder {
		preyOrder[i] = reverseMap[name]
	}

	return preyOrder
}
