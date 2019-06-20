package hydropathy

import (
	"github.com/knightjdr/cmgo/internal/pkg/read/bioplex"
	"github.com/knightjdr/cmgo/pkg/mapfunc"
)

func uniqueBioplex(bioplexData []bioplex.Entry) []string {
	interactors := make(map[string]bool, 0)
	for _, entry := range bioplexData {
		interactors[entry.Entrez.Source] = true
		interactors[entry.Entrez.Target] = true
	}

	return mapfunc.KeysStringBool(interactors)
}
