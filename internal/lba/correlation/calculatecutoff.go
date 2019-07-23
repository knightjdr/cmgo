package correlation

import (
	"math"

	"github.com/knightjdr/cmgo/pkg/stats"
)

func calculateCutoff(matrix [][]float64, idealEdgeNumber int) float64 {
	var bestEdgeNumber float64
	var cutoff float64
	for i := 0; i <= 100; i++ {
		edges := make([]int, len(matrix))
		testCutoff := float64(i) / 100
		for j, row := range matrix {
			for k, coefficient := range row {
				if j != k && coefficient >= testCutoff {
					edges[j]++
				}
			}
		}
		avg := stats.MedianInt(edges)
		if math.Abs(avg-float64(idealEdgeNumber)) <= math.Abs(bestEdgeNumber-float64(idealEdgeNumber)) {
			bestEdgeNumber = avg
			cutoff = testCutoff
		}
	}
	return cutoff
}
