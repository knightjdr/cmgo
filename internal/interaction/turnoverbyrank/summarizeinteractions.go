package turnoverbyrank

import (
	"sort"
)

type rankSummary struct {
	TurnoverPreys map[string]bool
	TurnoverRates []float64
}

func summarizeInteractions(sortedPreysPerBait map[string][]string, turnoverRates map[string]float64) map[int]*rankSummary {
	knownInteractionsByRank := make(map[int]*rankSummary, 0)

	for _, preys := range sortedPreysPerBait {
		for index, prey := range preys {
			preyRank := index + 1
			allocateMap(preyRank, &knownInteractionsByRank)
			addTurnoverRate(prey, turnoverRates, knownInteractionsByRank[preyRank])
		}
	}

	sortTurnoverRates(&knownInteractionsByRank)

	return knownInteractionsByRank
}

func allocateMap(key int, knownInteractions *map[int]*rankSummary) {
	if _, ok := (*knownInteractions)[key]; !ok {
		(*knownInteractions)[key] = &rankSummary{
			TurnoverPreys: make(map[string]bool, 0),
			TurnoverRates: make([]float64, 0),
		}
	}
}

func addTurnoverRate(prey string, turnoverRates map[string]float64, rankSummary *rankSummary) {
	if _, hasTurnoverRate := turnoverRates[prey]; hasTurnoverRate {
		if _, hasBeenUsedAlready := rankSummary.TurnoverPreys[prey]; !hasBeenUsedAlready {
			rankSummary.TurnoverPreys[prey] = true
			rankSummary.TurnoverRates = append(rankSummary.TurnoverRates, turnoverRates[prey])
		}
	}
}

func sortTurnoverRates(knownInteractions *map[int]*rankSummary) {
	for rank, rankSummary := range *knownInteractions {
		sort.Float64s(rankSummary.TurnoverRates)
		(*knownInteractions)[rank].TurnoverRates = rankSummary.TurnoverRates
	}
}
