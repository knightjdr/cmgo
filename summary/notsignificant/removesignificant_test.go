package notsignificant

import (
	"testing"

	"github.com/knightjdr/cmgo/read"
	"github.com/stretchr/testify/assert"
)

func TestRemoveSignificant(t *testing.T) {
	data := []read.SaintRow{
		{Bait: "a", PreyGene: "prey-1", FDR: 0.01},
		{Bait: "a", PreyGene: "prey-2", FDR: 0},
		{Bait: "a", PreyGene: "prey-3", FDR: 0.05},
		{Bait: "b", PreyGene: "prey-1", FDR: 0.03},
		{Bait: "b", PreyGene: "prey-5", FDR: 0.01},
		{Bait: "c", PreyGene: "prey-6", FDR: 0.09},
		{Bait: "c", PreyGene: "prey-7", FDR: 0.1},
		{Bait: "c", PreyGene: "prey-3", FDR: 0},
	}
	wanted := []read.SaintRow{
		{Bait: "c", PreyGene: "prey-6", FDR: 0.09},
		{Bait: "c", PreyGene: "prey-7", FDR: 0.1},
	}
	assert.Equal(t, wanted, removeSignificant(data, 0.01), "Should remove significant rows from SAINT data")
}
