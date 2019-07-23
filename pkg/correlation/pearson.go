package correlation

import (
	"math"

	"github.com/knightjdr/cmgo/pkg/stats"
)

// Pearson calculates the correlation coefficiant using Pearson's equation.
func Pearson(x, y []float64) float64 {
	xMean := stats.MeanFloat(x)
	yMean := stats.MeanFloat(y)
	numerator := float64(0)
	demoninatorLeft := float64(0)
	demoninatorRight := float64(0)
	for i := range x {
		numerator += (x[i] - xMean) * (y[i] - yMean)
		demoninatorLeft += math.Pow(x[i]-xMean, 2)
		demoninatorRight += math.Pow(y[i]-yMean, 2)
	}
	return numerator / (math.Sqrt(demoninatorLeft) * math.Sqrt(demoninatorRight))
}
