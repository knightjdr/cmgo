package transmembrane

import (
	"strconv"

	"github.com/knightjdr/cmgo/pkg/math"
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
	transmembranePreyData map[string]orientationData
}

type preySummary struct {
	cytosolicBaits    int
	cytosolicFraction float64
	cytosolicScore    float64
	length            int
	localization      string
	lumenalBaits      int
	lumenalFraction   float64
	lumenalScore      float64
	maxCytosolicScore float64
	maxLumenalScore   float64
	uniprotID         string
}

func summarize(options summaryOptions) map[string]preySummary {
	summary := make(map[string]preySummary, 0)

	rankMaximums := findRankMaximums(options.basis, append(options.cytosolicCompartments, options.lumenalCompartments...))

	for _, prey := range options.transmembranePreys {
		sequenceLength := options.transmembranePreyData[prey].Length
		cytosolicFraction := float64(options.transmembranePreyData[prey].Cytosolic) / float64(sequenceLength)
		lumenalFraction := float64(options.transmembranePreyData[prey].Lumenal) / float64(sequenceLength)

		rowIndex := slice.IndexOfString(prey, options.rows)

		cytosolicScore, cytosolicRank := findMaxScore(options.basis[rowIndex], options.cytosolicCompartments)
		lumenalScore, lumenalRank := findMaxScore(options.basis[rowIndex], options.lumenalCompartments)

		summary[prey] = preySummary{
			cytosolicBaits:    options.organelleBaitsPerPrey[prey]["cytosolic"],
			cytosolicFraction: math.Round(cytosolicFraction, 0.0001),
			cytosolicScore:    cytosolicScore,
			length:            sequenceLength,
			localization:      getPreyLocalization(prey, options.cytosolicPreys),
			lumenalBaits:      options.organelleBaitsPerPrey[prey]["lumenal"],
			lumenalFraction:   math.Round(lumenalFraction, 0.0001),
			lumenalScore:      lumenalScore,
			maxCytosolicScore: rankMaximums[cytosolicRank],
			maxLumenalScore:   rankMaximums[lumenalRank],
			uniprotID:         options.transmembranePreyData[prey].UniProt,
		}
	}

	return summary
}

func findRankMaximums(basis [][]float64, ranks []string) map[int]float64 {
	rankMaximums := make(map[int]float64, 0)

	for _, rank := range ranks {
		rankIndex, _ := strconv.Atoi(rank)
		rankMaximums[rankIndex] = 0
	}

	for _, row := range basis {
		for _, rank := range ranks {
			rankIndex, _ := strconv.Atoi(rank)
			score := row[rankIndex-1]
			if score > rankMaximums[rankIndex] {
				rankMaximums[rankIndex] = score
			}
		}
	}

	return rankMaximums
}

func findMaxScore(row []float64, ranks []string) (float64, int) {
	max := float64(-1)
	bestRank := 0

	for _, rank := range ranks {
		rankIndex, _ := strconv.Atoi(rank)
		value := row[rankIndex-1]
		if value > max {
			max = value
			bestRank = rankIndex
		}
	}

	return max, bestRank
}

func getPreyLocalization(prey string, cytosolicPreys []string) string {
	if slice.ContainsString(prey, cytosolicPreys) {
		return "cytosolic"
	}
	return "lumenal"
}
