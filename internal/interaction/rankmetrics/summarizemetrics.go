package rankmetrics

import (
	"sort"
)

type rankSummary struct {
	LysinePreys   map[string]bool
	Lysines       []int
	TurnoverPreys map[string]bool
	TurnoverRates []float64
}

func summarizeMetrics(sortedPreysPerBait map[string][]string, lysines map[string]int, turnoverRates map[string]float64) map[int]*rankSummary {
	metricsByRank := make(map[int]*rankSummary, 0)

	for _, preys := range sortedPreysPerBait {
		for index, prey := range preys {
			preyRank := index + 1
			allocateMap(preyRank, &metricsByRank)
			addLysines(prey, lysines, metricsByRank[preyRank])
			addTurnoverRate(prey, turnoverRates, metricsByRank[preyRank])
		}
	}

	sortMetrics(&metricsByRank)

	return metricsByRank
}

func allocateMap(key int, knownInteractions *map[int]*rankSummary) {
	if _, ok := (*knownInteractions)[key]; !ok {
		(*knownInteractions)[key] = &rankSummary{
			LysinePreys:   make(map[string]bool, 0),
			Lysines:       make([]int, 0),
			TurnoverPreys: make(map[string]bool, 0),
			TurnoverRates: make([]float64, 0),
		}
	}
}

func addLysines(prey string, lysines map[string]int, rankSummary *rankSummary) {
	if _, hasLysine := lysines[prey]; hasLysine {
		if _, hasBeenUsedAlready := rankSummary.LysinePreys[prey]; !hasBeenUsedAlready {
			rankSummary.LysinePreys[prey] = true
			rankSummary.Lysines = append(rankSummary.Lysines, lysines[prey])
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

func sortMetrics(knownInteractions *map[int]*rankSummary) {
	for rank, rankSummary := range *knownInteractions {
		sort.Ints(rankSummary.Lysines)
		sort.Float64s(rankSummary.TurnoverRates)
		(*knownInteractions)[rank].Lysines = rankSummary.Lysines
		(*knownInteractions)[rank].TurnoverRates = rankSummary.TurnoverRates
	}
}
