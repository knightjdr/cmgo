// Package geneontology contains functions for reading and parsing Gene Ontology (GO) files.
package geneontology

import (
	"sort"

	"github.com/knightjdr/cmgo/pkg/slice"
)

func getAllChildren(directChildren []string, children map[string][]string) []string {
	allChildren := directChildren
	for _, directChild := range directChildren {
		if _, ok := children[directChild]; ok {
			allChildren = append(allChildren, getAllChildren(children[directChild], children)...)
		}
	}
	return allChildren
}

// GetChildren defines all child terms for each GO term in namespace.
func (g *GOhierarchy) GetChildren(namespace string) {
	children := make(map[string][]string, 0)

	// Find direct children for parents.
	for id, term := range (*g)[namespace] {
		for _, directParent := range term.DirectParents {
			if _, ok := children[directParent]; !ok {
				children[directParent] = make([]string, 0)
			}
			children[directParent] = append(children[directParent], id)
		}
	}

	// Find all children GO IDs.
	for parentID, directChildren := range children {
		allChildren := getAllChildren(directChildren, children)
		allChildren = slice.UniqueStrings(allChildren)
		sort.Strings(allChildren)
		(*g)[namespace][parentID].Children = allChildren
	}
}

func getAllParents(directParents []string, hierarchy map[string]*GOterm) []string {
	allParents := directParents
	for _, directParent := range directParents {
		if _, ok := hierarchy[directParent]; ok {
			allParents = append(allParents, getAllParents(hierarchy[directParent].DirectParents, hierarchy)...)
		}
	}
	return allParents
}

// GetParents defines all parent terms for each GO term in namespace.
func (g *GOhierarchy) GetParents(namespace string) {
	// Find all parents for GO IDs.
	for id, term := range (*g)[namespace] {
		allParents := getAllParents(term.DirectParents, (*g)[namespace])
		allParents = slice.UniqueStrings(allParents)
		sort.Strings(allParents)
		(*g)[namespace][id].Parents = allParents
	}
}

// AreConsistent determines if two GO ids are consistent (equal or one is a child of the other).
func (g *GOhierarchy) AreConsistent(namespace, id1, id2 string) bool {
	if id1 == id2 ||
		slice.ContainsString(id1, (*g)[namespace][id2].Children) ||
		slice.ContainsString(id2, (*g)[namespace][id1].Children) {
		return true
	}
	return false
}
