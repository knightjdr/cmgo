package crapome

import (
	"github.com/knightjdr/cmgo/internal/pkg/read/saint"
)

func removeNonControls(baits []saint.BaitDatRow, interactions []saint.InterDatRow) ([]saint.BaitDatRow, []saint.InterDatRow) {
	filteredBaits := make([]saint.BaitDatRow, 0, len(baits))
	filteredInteractions := make([]saint.InterDatRow, 0, len(interactions))

	controls := make(map[string]bool, 0)
	for _, baitRow := range baits {
		if baitRow.Control {
			filteredBaits = append(filteredBaits, baitRow)
			controls[baitRow.ID] = true
		}
	}

	for _, interRow := range interactions {
		if _, ok := controls[interRow.ID]; ok {
			filteredInteractions = append(filteredInteractions, interRow)
		}
	}

	return filteredBaits, filteredInteractions
}
