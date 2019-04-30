package crapome

import (
	"github.com/knightjdr/cmgo/read"
)

func parseInteractions(interactions []read.InterDatRow) map[string]map[string]int {
	parsed := make(map[string]map[string]int, 0)

	for _, interaction := range interactions {
		id := interaction.ID
		prey := interaction.Prey
		spec := interaction.Spec
		if _, ok := parsed[prey]; !ok {
			parsed[prey] = make(map[string]int, 0)
		}

		parsed[prey][id] = spec
	}

	return parsed
}
