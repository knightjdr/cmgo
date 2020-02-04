package saint

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniProtMap(t *testing.T) {
	oldMapRefseqToUniProt := mapRefseqToUniProt
	mapRefseqToUniProt = func([]string, string) map[string][]string {
		return map[string][]string{
			"NP_000001": []string{"P00001", "Q00001"},
			"NP_000002": []string{"P00002"},
			"NP_000003": []string{"P00003", "Q00003"},
		}
	}
	defer func() {
		mapRefseqToUniProt = oldMapRefseqToUniProt
	}()

	saint := &SAINT{
		Row{Prey: "NP_000001", PreyGene: "ACADVL"},
		Row{Prey: "NP_000002", PreyGene: "ACAT1"},
		Row{Prey: "NP_000003", PreyGene: "DLD"},
		Row{Prey: "NP_000001", PreyGene: "ACADVL"},
		Row{Prey: "NP_000002", PreyGene: "ACAT1"},
	}

	expected := map[string]string{
		"P00001": "ACADVL",
		"Q00001": "ACADVL",
		"P00002": "ACAT1",
		"P00003": "DLD",
		"Q00003": "DLD",
	}
	actual := saint.GetUniProtMapping()
	assert.Equal(t, expected, actual, "should return a map of uniprot accessions to gene names")
}
