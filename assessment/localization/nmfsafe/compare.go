package nmfsafe

import (
	"github.com/knightjdr/cmgo/read/geneontology"
	"github.com/knightjdr/cmgo/slice"
)

func getAllChildren(ids []string, hierarchy map[string]*geneontology.GOterm) []string {
	children := make([]string, 0)

	for _, id := range ids {
		children = append(children, id)
		children = append(children, hierarchy[id].Children...)
	}
	children = slice.UniqueStrings(children)

	return children
}

func localizationAggreement(nmfIDs, safeIDs, nmfChildren, safeChildren []string) bool {
	if slice.HasIntersect(nmfIDs, safeChildren) || slice.HasIntersect(safeIDs, nmfChildren) {
		return true
	}
	return false
}

func compare(genes map[string]*localizationInfo, hierarchy map[string]*geneontology.GOterm) {
	for gene, info := range genes {
		nmfChildren := getAllChildren(info.NMFids, hierarchy)
		safeChildren := getAllChildren(info.SAFEids, hierarchy)

		genes[gene].NMFinSAFE = slice.HasIntersect(info.NMFids, safeChildren)
		genes[gene].SAFEinNMF = slice.HasIntersect(info.SAFEids, nmfChildren)
		genes[gene].Concordant = localizationAggreement(info.NMFids, info.SAFEids, nmfChildren, safeChildren)
	}
}
