package crapome

import (
	"github.com/knightjdr/cmgo/read"
)

func removeNonControls(baits []read.BaitDatRow, interactions []read.InterDatRow) ([]read.BaitDatRow, []read.InterDatRow) {
	filteredBaits := make([]read.BaitDatRow, 0, len(baits))
	filteredInteractions := make([]read.InterDatRow, 0, len(interactions))

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
