package correlation

import (
	"fmt"
	"sort"

	customSort "github.com/knightjdr/cmgo/pkg/sort"
)

type edgePair struct {
	Target string
	Weight float64
}

func filterPairs(corr [][]float64, genes []string, cutoff float64, maxEdges int) map[string][]edgePair {
	pairs := make(map[string][]edgePair, len(genes))
	edges := make(map[string]bool, 0)
	for i, gene := range genes {
		pairs[gene] = make([]edgePair, 0)
		geneEdges := make(map[string]float64, 0)
		for j := 0; j < len(corr); j++ {
			coefficient := corr[i][j]
			if i != j && coefficient >= cutoff {
				geneEdges[genes[j]] = coefficient
			}
		}

		topEdges := customSort.ByMapValueStringFloat(geneEdges, "descending")
		if maxEdges > 0 && len(topEdges) >= maxEdges {
			topEdges = topEdges[:maxEdges]
		}

		for _, target := range topEdges {
			sortedEdge := []string{gene, target.Key}
			sort.Strings(sortedEdge)
			edgeName := fmt.Sprintf("%s-%s", sortedEdge[0], sortedEdge[1])
			if _, ok := edges[edgeName]; !ok {
				edges[edgeName] = true
				edgeData := edgePair{
					Target: target.Key,
					Weight: target.Value,
				}
				pairs[gene] = append(pairs[gene], edgeData)
			}
		}
	}
	return pairs
}
