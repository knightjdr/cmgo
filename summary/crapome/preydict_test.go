package crapome

import (
	"testing"

	"github.com/knightjdr/cmgo/read"
	"github.com/stretchr/testify/assert"
)

func TestPreyDict(t *testing.T) {
	preys := []read.PreyDatRow{
		{Accession: "BirA_R118G_H0QFJ5", GeneID: 321, Name: "BirA_R118G_H0QFJ5"},
		{Accession: "NP_000029.2", GeneID: 2843, Name: "APC"},
		{Accession: "NP_000108.1", GeneID: 254, Name: "EMD"},
		{Accession: "NP_000280.1", GeneID: 780, Name: "PFKM"},
		{Accession: "NP_000402.3", GeneID: 726, Name: "HLCS"},
	}
	wanted := map[string]*preyDefinition{
		"BirA_R118G_H0QFJ5": &preyDefinition{GeneID: 321, Name: "BirA_R118G_H0QFJ5"},
		"NP_000029.2":       &preyDefinition{GeneID: 2843, Name: "APC"},
		"NP_000108.1":       &preyDefinition{GeneID: 254, Name: "EMD"},
		"NP_000280.1":       &preyDefinition{GeneID: 780, Name: "PFKM"},
		"NP_000402.3":       &preyDefinition{GeneID: 726, Name: "HLCS"},
	}
	assert.Equal(t, wanted, preyDict(preys), "Should convert an array of prey data to a map")
}
