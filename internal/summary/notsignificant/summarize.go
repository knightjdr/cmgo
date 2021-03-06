package notsignificant

import (
	"github.com/knightjdr/cmgo/internal/pkg/read/saint"
	customMath "github.com/knightjdr/cmgo/pkg/math"
	"github.com/knightjdr/cmgo/pkg/stats"
)

type preySummary struct {
	baits   []string
	bestFDR float64
	ctrlAvg float64
	maxSpec float64
}

type summary map[string]*preySummary

func summarize(data *saint.SAINT) map[string]*preySummary {
	preys := make(map[string]*preySummary, 0)

	for _, preyData := range *data {
		bait := preyData.Bait
		ctrl := preyData.Control
		fdr := preyData.FDR
		preyName := preyData.PreyGene
		spec := customMath.MaxSliceFloat(preyData.Spec)
		if _, ok := preys[preyName]; !ok {
			preys[preyName] = &preySummary{bestFDR: 1}
		}

		preys[preyName].baits = append(preys[preyName].baits, bait)
		if spec > preys[preyName].maxSpec {
			preys[preyName].ctrlAvg = stats.MeanFloat(ctrl)
			preys[preyName].maxSpec = spec
		}
		if fdr < preys[preyName].bestFDR {
			preys[preyName].bestFDR = fdr
		}
	}

	return preys
}
