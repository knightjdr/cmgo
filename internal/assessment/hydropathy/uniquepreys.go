package hydropathy

import (
	"github.com/knightjdr/cmgo/internal/pkg/read/saint"
	"github.com/knightjdr/cmgo/pkg/mapfunc"
)

func uniquePreys(saintData *saint.SAINT) []string {
	preys := make(map[string]bool, 0)
	for _, row := range *saintData {
		preys[row.Prey] = true
	}

	return mapfunc.KeysStringBool(preys)
}
