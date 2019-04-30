package crapome

import (
	"github.com/knightjdr/cmgo/read"
)

type preyDefinition struct {
	GeneID int
	Name   string
}

func preyDict(preys []read.PreyDatRow) map[string]*preyDefinition {
	preyMap := make(map[string]*preyDefinition, len(preys))

	for _, prey := range preys {
		preyMap[prey.Accession] = &preyDefinition{
			GeneID: prey.GeneID,
			Name:   prey.Name,
		}
	}
	return preyMap
}
