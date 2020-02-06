package preys

import "github.com/knightjdr/cmgo/internal/pkg/read/saint"

func readSaintFiles(options parameters) (map[string]string, map[string]map[string]int) {
	baitDat := saint.BaitDat(options.bait)
	interDat := saint.InterDat(options.inter)
	preyDat := saint.PreyDat(options.prey)

	baits := parseBaits(baitDat)
	preys := parsePreys(preyDat)

	interactions := parseInteractions(interDat, baits, preys)

	return baits, interactions
}

func parseBaits(baitDat []saint.BaitDatRow) map[string]string {
	baits := make(map[string]string, 0)

	for _, row := range baitDat {
		baits[row.ID] = row.Type
	}

	return baits
}

func parsePreys(preyDat []saint.PreyDatRow) map[string]string {
	preys := make(map[string]string, 0)

	for _, row := range preyDat {
		preys[row.Accession] = row.Name
	}

	return preys
}

func parseInteractions(interDat []saint.InterDatRow, baits map[string]string, preys map[string]string) map[string]map[string]int {
	interactions := make(map[string]map[string]int, 0)

	for _, row := range interDat {
		if _, ok := baits[row.ID]; ok {
			allocateBaitInteractionsMap(&interactions, row.ID)
			preyGene := preys[row.Prey]
			interactions[row.ID][preyGene] = row.Spec
		}
	}

	return interactions
}

func allocateBaitInteractionsMap(interactions *map[string]map[string]int, id string) {
	if _, ok := (*interactions)[id]; !ok {
		(*interactions)[id] = make(map[string]int, 0)
	}
}
