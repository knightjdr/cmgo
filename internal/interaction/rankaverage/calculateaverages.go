package rankaverage

import (
	"sort"

	"github.com/knightjdr/cmgo/pkg/math"
	"github.com/knightjdr/cmgo/pkg/slice"
	"github.com/knightjdr/cmgo/pkg/stats"
)

type preySummary struct {
	mean  float64
	ranks []int
	sd    float64
}

func calculateAverages(preys []string, sortedPreysPerBait map[string][]string) (map[string]preySummary, float64, float64) {
	values := make([]int, 0)
	summary := make(map[string]preySummary, len(preys))

	for _, prey := range preys {
		ranks := make([]int, 0)
		for _, rankedInteractors := range sortedPreysPerBait {
			index := slice.IndexOfString(prey, rankedInteractors)
			if index > -1 {
				ranks = append(ranks, index+1)
			}
		}

		sort.Ints(ranks)
		summary[prey] = preySummary{
			mean:  math.Round(stats.MeanInt(ranks), 0.001),
			ranks: ranks,
			sd:    math.Round(stats.SDInt(ranks), 0.001),
		}
		values = append(values, ranks...)
	}

	mean := math.Round(stats.MeanInt(values), 0.001)
	sd := math.Round(stats.SDInt(values), 0.001)
	return summary, mean, sd
}
