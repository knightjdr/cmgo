package crapome

import (
	"github.com/knightjdr/cmgo/internal/pkg/read/saint"
)

func preyDict(preys []saint.PreyDatRow) map[string]string {
	preyMap := make(map[string]string, len(preys))

	for _, prey := range preys {
		preyMap[prey.Accession] = prey.Name
	}
	return preyMap
}
