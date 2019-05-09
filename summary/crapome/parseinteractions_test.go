package crapome

import (
	"testing"

	"github.com/knightjdr/cmgo/read/saint"
	"github.com/stretchr/testify/assert"
)

func TestParseInteractions(t *testing.T) {
	interactions := []saint.InterDatRow{
		{Bait: "ACTB", ID: "128_468", Prey: "A", Spec: 1410},
		{Bait: "ACTB", ID: "128_468", Prey: "B", Spec: 4},
		{Bait: "ATP2A1", ID: "128_590", Prey: "C", Spec: 2},
		{Bait: "ATP2A1", ID: "128_590", Prey: "D", Spec: 3},
		{Bait: "737_BirAFLAG", ID: "128_737", Prey: "A", Spec: 11},
		{Bait: "737_BirAFLAG", ID: "128_737", Prey: "B", Spec: 11},
	}
	wanted := map[string]map[string]int{
		"A": map[string]int{
			"128_468": 1410,
			"128_737": 11,
		},
		"B": map[string]int{
			"128_468": 4,
			"128_737": 11,
		},
		"C": map[string]int{
			"128_590": 2,
		},
		"D": map[string]int{
			"128_590": 3,
		},
	}
	assert.Equal(t, wanted, parseInteractions(interactions), "Should produce result")
}
