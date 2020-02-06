package preys

import (
	"sort"

	"github.com/knightjdr/cmgo/pkg/math"
	"github.com/knightjdr/cmgo/pkg/stats"
)

type preyInteraction struct {
	Average  average
	BirAFlag []int
	BirAGFP  []int
	Empty    []int
	Max      maxMin
	Min      maxMin
}

type average struct {
	BirAFlag float64
	BirAGFP  float64
	Empty    float64
	Overall  float64
}

type maxMin struct {
	BirAFlag int
	BirAGFP  int
	Empty    int
	Overall  int
}

type controls struct {
	BirAFlag int
	BirAGFP  int
	Empty    int
	Overall  int
}

func summarize(baits map[string]string, interactions map[string]map[string]int) map[string]*preyInteraction {
	summary := make(map[string]*preyInteraction, 0)

	for bait, preys := range interactions {
		controlType := baits[bait]
		for prey, spec := range preys {
			allocatePreySummary(&summary, prey)
			if controlType == "bira-flag" {
				summary[prey].BirAFlag = append(summary[prey].BirAFlag, spec)
			}
			if controlType == "bira-gfp" {
				summary[prey].BirAGFP = append(summary[prey].BirAGFP, spec)
			}
			if controlType == "empty" {
				summary[prey].Empty = append(summary[prey].Empty, spec)
			}
		}
	}

	totalControls := calculateTotalControls(baits)
	calculateMetrics(&summary, totalControls)

	return summary
}

func allocatePreySummary(summary *map[string]*preyInteraction, prey string) {
	if _, ok := (*summary)[prey]; !ok {
		(*summary)[prey] = &preyInteraction{
			BirAFlag: []int{},
			BirAGFP:  []int{},
			Empty:    []int{},
		}
	}
}

func calculateTotalControls(baits map[string]string) controls {
	totalControls := controls{
		Overall: len(baits),
	}

	for _, controlType := range baits {
		if controlType == "bira-flag" {
			totalControls.BirAFlag++
		}
		if controlType == "bira-gfp" {
			totalControls.BirAGFP++
		}
		if controlType == "empty" {
			totalControls.Empty++
		}
	}

	return totalControls
}

func calculateMetrics(summary *map[string]*preyInteraction, totalControls controls) {
	for prey, preySummary := range *summary {
		(*summary)[prey].Average = calculateAverage(preySummary, totalControls.Overall)
		(*summary)[prey].Max = calculateMax(preySummary)
		(*summary)[prey].Min = calculateMin(preySummary, totalControls)
		sort.Ints((*summary)[prey].BirAFlag)
		sort.Ints((*summary)[prey].BirAGFP)
		sort.Ints((*summary)[prey].Empty)
	}
}

func calculateAverage(summary *preyInteraction, totalControls int) average {
	overallSum := float64(0)
	for _, value := range appendAllControlValues(summary) {
		overallSum += float64(value)
	}
	overallAverage := overallSum / float64(totalControls)

	return average{
		BirAFlag: stats.MeanInt(summary.BirAFlag),
		BirAGFP:  stats.MeanInt(summary.BirAGFP),
		Empty:    stats.MeanInt(summary.Empty),
		Overall:  overallAverage,
	}
}

func calculateMax(summary *preyInteraction) maxMin {
	maxBirAFlag := math.MaxSliceInt(summary.BirAFlag)
	maxBirAGFP := math.MaxSliceInt(summary.BirAGFP)
	maxEmpty := math.MaxSliceInt(summary.Empty)
	maxOverall := math.MaxSliceInt(appendAllControlValues(summary))
	return maxMin{
		BirAFlag: maxBirAFlag,
		BirAGFP:  maxBirAGFP,
		Empty:    maxEmpty,
		Overall:  maxOverall,
	}
}

func calculateMin(summary *preyInteraction, totalControls controls) maxMin {
	return maxMin{
		BirAFlag: minCount(summary.BirAFlag, totalControls.BirAFlag),
		BirAGFP:  minCount(summary.BirAGFP, totalControls.BirAGFP),
		Empty:    minCount(summary.Empty, totalControls.Empty),
		Overall:  minCount(appendAllControlValues(summary), totalControls.Overall),
	}
}

func minCount(ints []int, total int) int {
	if len(ints) < total {
		return 0
	}
	min, _ := math.MinSliceInt(ints)
	return min
}

func appendAllControlValues(summary *preyInteraction) []int {
	all := append(summary.BirAFlag, summary.BirAGFP...)
	return append(all, summary.Empty...)
}
