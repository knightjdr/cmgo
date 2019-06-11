package crapome

import (
	"testing"

	"github.com/knightjdr/cmgo/internal/pkg/read/saint"
	"github.com/stretchr/testify/assert"
)

func TestRemoveNonControls(t *testing.T) {
	baits := []saint.BaitDatRow{
		{Control: false, ID: "128_468", Name: "ACTB"},
		{Control: false, ID: "128_590", Name: "ATP2A1"},
		{Control: true, ID: "128_737", Name: "737_BirAFLAG"},
	}
	interactions := []saint.InterDatRow{
		{Bait: "ACTB", ID: "128_468", Prey: "A", Spec: 1410},
		{Bait: "ACTB", ID: "128_468", Prey: "B", Spec: 4},
		{Bait: "ATP2A1", ID: "128_590", Prey: "C", Spec: 2},
		{Bait: "ATP2A1", ID: "128_590", Prey: "D", Spec: 3},
		{Bait: "737_BirAFLAG", ID: "128_737", Prey: "E", Spec: 11},
		{Bait: "737_BirAFLAG", ID: "128_737", Prey: "F", Spec: 11},
	}
	wantedBaits := []saint.BaitDatRow{
		{Control: true, ID: "128_737", Name: "737_BirAFLAG"},
	}
	wantedInteractions := []saint.InterDatRow{
		{Bait: "737_BirAFLAG", ID: "128_737", Prey: "E", Spec: 11},
		{Bait: "737_BirAFLAG", ID: "128_737", Prey: "F", Spec: 11},
	}
	resultBaits, resultInteractions := removeNonControls(baits, interactions)
	assert.Equal(t, wantedBaits, resultBaits, "Should remove non control samples from baits")
	assert.Equal(t, wantedInteractions, resultInteractions, "Should remove non control samples from interactions")
}
