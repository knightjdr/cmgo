package transmembrane

import (
	"strconv"

	"github.com/knightjdr/cmgo/pkg/slice"
)

type summaryOptions struct {
	basis                 [][]float64
	cytosolicCompartments []string
	cytosolicPreys        []string
	lumenalCompartments   []string
	lumenalPreys          []string
	organelleBaitsPerPrey map[string]map[string]int
	rows                  []string
	transmembranePreys    []string
}

type preySummary struct {
	cytosolicBaits int
	cytosolicScore float64
	localization   string
	lumenalBaits   int
	lumenalScore   float64
}

func summarize(options summaryOptions) map[string]preySummary {
	summary := make(map[string]preySummary, 0)

	for _, prey := range options.transmembranePreys {
		rowIndex := slice.IndexOfString(prey, options.rows)
		summary[prey] = preySummary{
			cytosolicBaits: options.organelleBaitsPerPrey[prey]["cytosolic"],
			cytosolicScore: findMaxScore(options.basis[rowIndex], options.cytosolicCompartments),
			localization:   getPreyLocalization(prey, options.cytosolicPreys),
			lumenalBaits:   options.organelleBaitsPerPrey[prey]["lumenal"],
			lumenalScore:   findMaxScore(options.basis[rowIndex], options.lumenalCompartments),
		}
	}

	return summary
}

func findMaxScore(row []float64, ranks []string) float64 {
	max := float64(0)

	for _, rank := range ranks {
		rankIndex, _ := strconv.Atoi(rank)
		value := row[rankIndex-1]
		if value > max {
			max = value
		}
	}

	return max
}

func getPreyLocalization(prey string, cytosolicPreys []string) string {
	if slice.Contains(prey, cytosolicPreys) {
		return "cytosolic"
	}
	return "lumenal"
}
