package crapome

import (
	"testing"

	"github.com/knightjdr/cmgo/read"
	"github.com/stretchr/testify/assert"
)

func TestPreyDict(t *testing.T) {
	preys := []read.PreyDatRow{
		{Accession: "BirA_R118G_H0QFJ5", Length: 321, Name: "BirA_R118G_H0QFJ5"},
		{Accession: "NP_000029.2", Length: 2843, Name: "APC"},
		{Accession: "NP_000108.1", Length: 254, Name: "EMD"},
		{Accession: "NP_000280.1", Length: 780, Name: "PFKM"},
		{Accession: "NP_000402.3", Length: 726, Name: "HLCS"},
	}
	wanted := map[string]string{
		"BirA_R118G_H0QFJ5": "BirA_R118G_H0QFJ5",
		"NP_000029.2":       "APC",
		"NP_000108.1":       "EMD",
		"NP_000280.1":       "PFKM",
		"NP_000402.3":       "HLCS",
	}
	assert.Equal(t, wanted, preyDict(preys), "Should convert an array of prey data to a map")
}
