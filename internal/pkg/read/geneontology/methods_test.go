package geneontology

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoGetChildren(t *testing.T) {
	hierarchy := &GOhierarchy{
		"CC": map[string]*GOterm{
			"GO:0034396": &GOterm{
				DirectParents: []string{},
				Name:          "parent a",
			},
			"GO:0034399": &GOterm{
				DirectParents: []string{"GO:0034396"},
				Name:          "parent a",
			},
			"GO:0044428": &GOterm{
				DirectParents: []string{"GO:0034396"},
				Name:          "parent b",
			},
			"GO:0005652": &GOterm{
				DirectParents: []string{"GO:0044428", "GO:0034399"},
				Name:          "nuclear lamina",
			},
		},
	}

	// TEST
	wanted := &GOhierarchy{
		"CC": map[string]*GOterm{
			"GO:0034396": &GOterm{
				Children:      []string{"GO:0005652", "GO:0034399", "GO:0044428"},
				DirectParents: []string{},
				Name:          "parent a",
			},
			"GO:0034399": &GOterm{
				Children:      []string{"GO:0005652"},
				DirectParents: []string{"GO:0034396"},
				Name:          "parent a",
			},
			"GO:0044428": &GOterm{
				Children:      []string{"GO:0005652"},
				DirectParents: []string{"GO:0034396"},
				Name:          "parent b",
			},
			"GO:0005652": &GOterm{
				DirectParents: []string{"GO:0044428", "GO:0034399"},
				Name:          "nuclear lamina",
			},
		},
	}
	hierarchy.GetChildren("CC")
	assert.Equal(t, wanted, hierarchy, "Should define children for each GO term")
}

func TestGoGetParents(t *testing.T) {
	hierarchy := &GOhierarchy{
		"CC": map[string]*GOterm{
			"GO:0034396": &GOterm{
				DirectParents: []string{},
				Name:          "parent a",
			},
			"GO:0034399": &GOterm{
				DirectParents: []string{"GO:0034396"},
				Name:          "parent a",
			},
			"GO:0044428": &GOterm{
				DirectParents: []string{"GO:0034396"},
				Name:          "parent b",
			},
			"GO:0005652": &GOterm{
				DirectParents: []string{"GO:0044428", "GO:0034399"},
				Name:          "nuclear lamina",
			},
		},
	}

	// TEST
	wanted := &GOhierarchy{
		"CC": map[string]*GOterm{
			"GO:0034396": &GOterm{
				DirectParents: []string{},
				Name:          "parent a",
				Parents:       []string{},
			},
			"GO:0034399": &GOterm{
				DirectParents: []string{"GO:0034396"},
				Name:          "parent a",
				Parents:       []string{"GO:0034396"},
			},
			"GO:0044428": &GOterm{
				DirectParents: []string{"GO:0034396"},
				Name:          "parent b",
				Parents:       []string{"GO:0034396"},
			},
			"GO:0005652": &GOterm{
				DirectParents: []string{"GO:0044428", "GO:0034399"},
				Name:          "nuclear lamina",
				Parents:       []string{"GO:0034396", "GO:0034399", "GO:0044428"},
			},
		},
	}
	hierarchy.GetParents("CC")
	assert.Equal(t, wanted, hierarchy, "Should define parents for each GO term")
}

func TestAreConsitent(t *testing.T) {
	hierarchy := &GOhierarchy{
		"CC": map[string]*GOterm{
			"GO:0034396": &GOterm{
				Children:      []string{"GO:0005652", "GO:0034399", "GO:0044428"},
				DirectParents: []string{},
				Name:          "parent a",
			},
			"GO:0034399": &GOterm{
				Children:      []string{"GO:0005652"},
				DirectParents: []string{"GO:0034396"},
				Name:          "parent a",
			},
			"GO:0044428": &GOterm{
				Children:      []string{"GO:0005652"},
				DirectParents: []string{"GO:0034396"},
				Name:          "parent b",
			},
			"GO:0005652": &GOterm{
				DirectParents: []string{"GO:0044428", "GO:0034399"},
				Name:          "nuclear lamina",
			},
			"GO:0005653": &GOterm{
				DirectParents: []string{"GO:0044427", "GO:0034392"},
				Name:          "made up term",
			},
		},
	}

	// TEST: equal IDs
	assert.True(t, hierarchy.AreConsistent("CC", "GO:0034396", "GO:0034396"), "should consider same ID consistent")

	// TEST: id1 is child of id2
	assert.True(t, hierarchy.AreConsistent("CC", "GO:0005652", "GO:0034396"), "should consider children consistent")

	// TEST: id2 is child of id1
	assert.True(t, hierarchy.AreConsistent("CC", "GO:0034396", "GO:0034399"), "should consider children consistent")

	// TEST: neither id is child of the other
	assert.False(t, hierarchy.AreConsistent("CC", "GO:0005652", "GO:0005653"), "should not consider unrelated IDs consistent")
}
