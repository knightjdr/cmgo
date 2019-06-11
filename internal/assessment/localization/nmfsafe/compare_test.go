package nmfsafe

import (
	"sort"
	"testing"

	"github.com/knightjdr/cmgo/internal/pkg/read/geneontology"
	"github.com/stretchr/testify/assert"
)

func TestGetAllChildren(t *testing.T) {
	ids := []string{"GO:1", "GO:2"}
	hierarchy := map[string]*geneontology.GOterm{
		"GO:1": &geneontology.GOterm{
			Children: []string{"GO:3", "GO:4"},
		},
		"GO:2": &geneontology.GOterm{
			Children: []string{"GO:3", "GO:5"},
		},
	}

	wanted := []string{"GO:1", "GO:2", "GO:3", "GO:4", "GO:5"}
	result := getAllChildren(ids, hierarchy)
	sort.Strings(result)
	assert.Equal(t, wanted, result, "Should return input IDs and their children")
}

func TestLocalizationAggreement(t *testing.T) {
	// TEST1: one of the first slice's IDs is in the children of the second
	idsA := []string{"GO:1", "GO:2"}
	idsB := []string{"GO:3", "GO:4"}
	childrenA := []string{"GO:1", "GO:2"}
	childrenB := []string{"GO:2", "GO:3", "GO:4"}
	assert.True(t, localizationAggreement(idsA, idsB, childrenA, childrenB), "Should return true when first slice of IDs is in the children of the second")

	// TEST2: one of the second slice's IDs is in the children of the first
	idsA = []string{"GO:3", "GO:4"}
	idsB = []string{"GO:1", "GO:2"}
	childrenA = []string{"GO:2", "GO:3", "GO:4"}
	childrenB = []string{"GO:1", "GO:2"}
	assert.True(t, localizationAggreement(idsA, idsB, childrenA, childrenB), "Should return true when second slice of IDs is in the children of the first")

	// TEST3: neither slice contains an ID in the others' children
	idsA = []string{"GO:1", "GO:2"}
	idsB = []string{"GO:3", "GO:4"}
	childrenA = []string{"GO:1", "GO:2"}
	childrenB = []string{"GO:3", "GO:4"}
	assert.False(t, localizationAggreement(idsA, idsB, childrenA, childrenB), "Should return false when neither slice contains an ID in the others' children")
}

func TestComparse(t *testing.T) {
	genes := map[string]*localizationInfo{
		"A": &localizationInfo{
			NMFids:  []string{"GO:1"},
			SAFEids: []string{"GO:3"},
		},
		"B": &localizationInfo{
			NMFids:  []string{"GO:1"},
			SAFEids: []string{"GO:6"},
		},
	}
	hierarchy := map[string]*geneontology.GOterm{
		"GO:1": &geneontology.GOterm{
			Children: []string{"GO:3", "GO:4"},
		},
		"GO:2": &geneontology.GOterm{
			Children: []string{"GO:3", "GO:5"},
		},
		"GO:3": &geneontology.GOterm{
			Children: []string{"GO:5"},
		},
		"GO:4": &geneontology.GOterm{
			Children: []string{"GO:5"},
		},
		"GO:5": &geneontology.GOterm{
			Children: []string{},
		},
		"GO:6": &geneontology.GOterm{
			Children: []string{"GO:7"},
		},
		"GO:7": &geneontology.GOterm{
			Children: []string{},
		},
	}

	wanted := map[string]*localizationInfo{
		"A": &localizationInfo{
			Concordant: true,
			NMFids:     []string{"GO:1"},
			NMFinSAFE:  false,
			SAFEids:    []string{"GO:3"},
			SAFEinNMF:  true,
		},
		"B": &localizationInfo{
			Concordant: false,
			NMFids:     []string{"GO:1"},
			NMFinSAFE:  false,
			SAFEids:    []string{"GO:6"},
			SAFEinNMF:  false,
		},
	}
	compare(genes, hierarchy)
	assert.Equal(t, wanted, genes, "Should add NMF and SAFE agreement data to gene map")
}
