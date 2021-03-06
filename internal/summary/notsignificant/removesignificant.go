package notsignificant

import (
	"github.com/knightjdr/cmgo/internal/pkg/read/saint"
)

func removeSignificant(data *saint.SAINT, fdr float64) *saint.SAINT {
	significant := make(map[string]bool, 0)
	for _, row := range *data {
		if row.FDR <= fdr {
			significant[row.PreyGene] = true
		}
	}

	filtered := new(saint.SAINT)
	for _, row := range *data {
		if _, ok := significant[row.PreyGene]; !ok {
			*filtered = append(*filtered, row)
		}
	}
	return filtered
}
