// Package localization contains functions for assessing localizations
package localization

import (
	"github.com/knightjdr/cmgo/read/geneontology"
	"github.com/knightjdr/cmgo/slice"
)

func addImputed(id string, hierarchy map[string]*geneontology.GOterm) []string {
	withParents := make([]string, 0)
	withParents = append(withParents, id)
	return append(withParents, hierarchy[id].Parents...)
}

// IsKnown determines if an assigned localization(s) is previously known.
func IsKnown(gene string, assignedIDs []string, annotations map[string]map[string]*geneontology.GOannotation, hierarchy map[string]*geneontology.GOterm) bool {
	// Get known IDs.
	knownIDs := make([]string, 0)
	for id := range annotations[gene] {
		knownIDs = append(knownIDs, addImputed(id, hierarchy)...)
	}

	for _, id := range assignedIDs {
		isKnown := slice.Contains(id, knownIDs)
		if isKnown {
			return true
		}
	}

	return false
}
