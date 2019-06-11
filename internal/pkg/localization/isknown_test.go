package localization

import (
	"testing"

	"github.com/knightjdr/cmgo/internal/pkg/read/geneontology"
	"github.com/stretchr/testify/assert"
)

func TestIsKnown(t *testing.T) {
	annotations := map[string]map[string]*geneontology.GOannotation{
		"a": map[string]*geneontology.GOannotation{
			"GO:1": &geneontology.GOannotation{},
		},
		"b": map[string]*geneontology.GOannotation{
			"GO:1": &geneontology.GOannotation{},
			"GO:7": &geneontology.GOannotation{},
		},
	}
	hierarchy := map[string]*geneontology.GOterm{
		"GO:1": &geneontology.GOterm{
			Parents: []string{"GO:9"},
		},
		"GO:2": &geneontology.GOterm{
			Parents: []string{"GO:8"},
		},
		"GO:7": &geneontology.GOterm{
			Parents: []string{"GO:9"},
		},
		"GO:8": &geneontology.GOterm{
			Parents: []string{},
		},
		"GO:9": &geneontology.GOterm{
			Parents: []string{},
		},
	}

	// TEST1: known localization.
	gene := "a"
	assignedIDs := []string{"GO:1", "GO:2"}
	assert.True(t, IsKnown(gene, assignedIDs, annotations, hierarchy), "Should return true when a localization is known")

	// TEST2: not previously known localization.
	gene = "b"
	assignedIDs = []string{"GO:2"}
	assert.False(t, IsKnown(gene, assignedIDs, annotations, hierarchy), "Should return false when a localization is not known")
}
