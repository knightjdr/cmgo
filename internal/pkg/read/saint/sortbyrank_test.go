package saint

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortByPreyRank(t *testing.T) {
	saint := &SAINT{
		Row{Bait: "AARS2", PreyGene: "ACADVL", AvgSpec: 4.5, NormalizedSpec: 4},
		Row{Bait: "AARS2", PreyGene: "ACAT1", AvgSpec: 7, NormalizedSpec: 12},
		Row{Bait: "AARS2", PreyGene: "DLD", AvgSpec: 16, NormalizedSpec: 8},
		Row{Bait: "ABCC1", PreyGene: "ACADVL", AvgSpec: 8, NormalizedSpec: 17},
		Row{Bait: "ABCC1", PreyGene: "ACAT1", AvgSpec: 3, NormalizedSpec: 13},
		Row{Bait: "ABCC1", PreyGene: "DLD", AvgSpec: 6, NormalizedSpec: 9},
	}

	// TEST1: sort by AvgSpec
	wanted := map[string][]string{
		"AARS2": []string{"DLD", "ACAT1", "ACADVL"},
		"ABCC1": []string{"ACADVL", "DLD", "ACAT1"},
	}
	actual := saint.SortByPreyRank("AvgSpec")
	assert.Equal(t, wanted, actual, "Should returns preys sorted by AvgSpec")

	// TEST2: sort by NormalizedSpec
	wanted = map[string][]string{
		"AARS2": []string{"ACAT1", "DLD", "ACADVL"},
		"ABCC1": []string{"ACADVL", "ACAT1", "DLD"},
	}
	actual = saint.SortByPreyRank("NormalizedSpec")
	assert.Equal(t, wanted, actual, "Should returns preys sorted by NormalizedSpec")
}
