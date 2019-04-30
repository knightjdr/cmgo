package notsignificant

import (
	"github.com/knightjdr/cmgo/read"
)

func removeSignificant(data []read.SaintRow, fdr float64) []read.SaintRow {
	significant := make(map[string]bool, 0)
	for _, row := range data {
		if row.FDR <= fdr {
			significant[row.PreyGene] = true
		}
	}

	filtered := make([]read.SaintRow, 0, len(data))
	for _, row := range data {
		if _, ok := significant[row.PreyGene]; !ok {
			filtered = append(filtered, row)
		}
	}
	return filtered
}
