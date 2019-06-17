package robustness

import "github.com/knightjdr/cmgo/pkg/stats"

type meanSD struct {
	Mean float64
	SD   float64
}

func summaryStats(data [][][]float64, percentiles []float64) [][]meanSD {
	summary := make([][]meanSD, len(data))
	for rankIndex, rankData := range data {
		summary[rankIndex] = make([]meanSD, len(percentiles))
		for percentileIndex, percentileData := range rankData {
			summary[rankIndex][percentileIndex] = meanSD{
				Mean: stats.MeanFloat(percentileData),
				SD:   stats.SDFloat(percentileData),
			}
		}
	}
	return summary
}
