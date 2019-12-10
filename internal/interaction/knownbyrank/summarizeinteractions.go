package knownbyrank

import (
	"fmt"
	"github.com/knightjdr/cmgo/pkg/slice"
	"sort"
)

type rankSummary struct {
	BaitNumber int
	Known      int
	Pairs      []string
}

func summarizeInteractions(sortedPreysPerBait, knownInteractions map[string][]string) map[int]*rankSummary {
	knownInteractionsByRank := make(map[int]*rankSummary, 0)

	for bait, preys := range sortedPreysPerBait {
		for index, prey := range preys {
			preyRank := index + 1
			allocateMap(preyRank, &knownInteractionsByRank)
			knownInteractionsByRank[preyRank].BaitNumber++
			if slice.ContainsString(prey, knownInteractions[bait]) {
				knownInteractionsByRank[preyRank].Known++
				pair := fmt.Sprintf("%s-%s", bait, prey)
				knownInteractionsByRank[preyRank].Pairs = append(knownInteractionsByRank[preyRank].Pairs, pair)
			}
		}
	}

	sortPairs(&knownInteractionsByRank)

	return knownInteractionsByRank
}

func allocateMap(key int, knownInteractions *map[int]*rankSummary) {
	if _, ok := (*knownInteractions)[key]; !ok {
		(*knownInteractions)[key] = &rankSummary{
			BaitNumber: 0,
			Known:      0,
			Pairs:      make([]string, 0),
		}
	}
}

func sortPairs(knownInteractions *map[int]*rankSummary) {
	for rank, rankSummary := range *knownInteractions {
		sort.Strings(rankSummary.Pairs)
		(*knownInteractions)[rank].Pairs = rankSummary.Pairs
	}
}
