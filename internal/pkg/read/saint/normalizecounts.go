package saint

import (
	"math"

	customMath "github.com/knightjdr/cmgo/pkg/math"
	"github.com/knightjdr/cmgo/pkg/stats"
)

// LengthNormalizeSpectralCounts substracts control counts from the average spectral count
// and then multiplies by the median prey length for the bait / current prey length.
func (s *SAINT) LengthNormalizeSpectralCounts() {
	medianPreyLengths := calculateMedianPreyLengths(s)
	normalizePreySpectralCounts(s, medianPreyLengths)
}

func calculateMedianPreyLengths(saint *SAINT) map[string]float64 {
	preyLengthsPerBait := getPreyLengthsPerBait(saint)
	return calculateMedianPreyLengthsPerBait(preyLengthsPerBait)
}

func getPreyLengthsPerBait(saint *SAINT) map[string][]int {
	preyLengthsPerBait := make(map[string][]int, 0)
	for _, row := range *saint {
		if row.PreySequenceLength != 0 {
			preyLengthsPerBait[row.Bait] = append(preyLengthsPerBait[row.Bait], row.PreySequenceLength)
		}
	}
	return preyLengthsPerBait
}

func calculateMedianPreyLengthsPerBait(preyLengthsPerBait map[string][]int) map[string]float64 {
	medianPreyLengthPerBait := make(map[string]float64, 0)
	for bait, preyLengths := range preyLengthsPerBait {
		medianPreyLengthPerBait[bait] = stats.MedianInt(preyLengths)
	}
	return medianPreyLengthPerBait
}

func normalizePreySpectralCounts(saint *SAINT, medianPreyLengths map[string]float64) {
	for rowIndex, row := range *saint {
		var sequenceLength float64
		if row.PreySequenceLength != 0 {
			sequenceLength = float64(row.PreySequenceLength)
		} else {
			sequenceLength = medianPreyLengths[row.Bait]
		}

		controlAverage := stats.MeanFloat(row.Control)
		controlSubtracted := math.Max(row.AvgSpec-controlAverage, 0)

		row.NormalizedSpec = customMath.Round(controlSubtracted*(medianPreyLengths[row.Bait]/sequenceLength), 0.01)
		(*saint)[rowIndex] = row
	}
}
