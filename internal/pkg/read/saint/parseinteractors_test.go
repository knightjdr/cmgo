package saint

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInteractors(t *testing.T) {
	saint := &SAINT{
		Row{Bait: "AARS2", PreyGene: "ACADVL", FDR: 0},
		Row{Bait: "AARS2", PreyGene: "ACAT1", FDR: 0.05},
		Row{Bait: "AARS2", PreyGene: "DLD", FDR: 0.02},
		Row{Bait: "ABCC1", PreyGene: "ACADVL", FDR: 0.01},
	}

	wanted := map[string][]string{
		"AARS2": []string{"ACADVL", "DLD"},
		"ABCC1": []string{"ACADVL"},
	}
	actual := saint.ParseInteractors(0.02)
	assert.Equal(t, wanted, actual, "Should return significant interactors for each bait")
}
