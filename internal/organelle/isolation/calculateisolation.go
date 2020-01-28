package isolation

import (
	"sort"
)

type isolationScores map[int]*isolationScore

type isolationScore struct {
	edgesOutside       int
	edgesWithin        int
	isolation          float64
	nodesOutside       []int
	sharedCompartments []int
}

func calculateIsolation(correlation [][]float64, cutoff float64, preysPerCompartment map[int]map[int]bool) *isolationScores {
	countEdges := getEdgeCounter(correlation, cutoff)
	noCompartments := len(preysPerCompartment)
	scores := &isolationScores{}

	for compartment, genes := range preysPerCompartment {
		score := initializeScore(noCompartments + 1)
		countEdges(score, genes)
		scoreIsolation(score)
		(*scores)[compartment] = score
	}

	return scores
}

func initializeScore(noCompartments int) *isolationScore {
	return &isolationScore{
		nodesOutside:       make([]int, 0),
		sharedCompartments: make([]int, noCompartments),
	}
}

func getEdgeCounter(correlation [][]float64, cutoff float64) func(*isolationScore, map[int]bool) {
	noRows := len(correlation)

	return func(score *isolationScore, genes map[int]bool) {
		for geneIndex := range genes {
			for j := 0; j < noRows; j++ {
				if j != geneIndex && correlation[geneIndex][j] >= cutoff {
					if _, ok := genes[j]; ok {
						score.edgesWithin++
					} else {
						score.edgesOutside++
						score.nodesOutside = append(score.nodesOutside, j)
					}
				}
			}
		}
	}
}

func scoreIsolation(score *isolationScore) {
	score.edgesWithin /= 2
	sort.Ints(score.nodesOutside)

	score.isolation = float64(score.edgesWithin) / float64(score.edgesWithin+score.edgesOutside)
}
