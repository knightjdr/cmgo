package overlap

import (
	"bytes"
	"fmt"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/knightjdr/cmgo/internal/organelle"
	"github.com/knightjdr/cmgo/pkg/slice"
	"github.com/knightjdr/cmgo/pkg/stats"
	"github.com/spf13/afero"
)

func rangeIndex(source, target string, dict1, dict2 map[string]bool) int {
	if _, okSource := dict1[source]; okSource {
		if _, okTarget := dict1[target]; okTarget {
			return 0
		} else if _, okTarget := dict2[target]; okTarget {
			return 2
		}
	} else if _, okSource := dict2[source]; okSource {
		if _, okTarget := dict2[target]; okTarget {
			return 1
		} else if _, okTarget := dict1[target]; okTarget {
			return 2
		}
	}
	return -1
}

func compare(compartments organelle.Compartments, similarity map[string]map[string]float64, outfile string) {
	compartmentDict1 := slice.Dict(compartments[0].Proteins)
	compartmentDict2 := slice.Dict(compartments[1].Proteins)

	ranges := make([][]float64, 3)
	ranges[0] = make([]float64, 0)
	ranges[1] = make([]float64, 0)
	ranges[2] = make([]float64, 0)

	for source, targets := range similarity {
		for target, distance := range targets {
			if source != target {
				index := rangeIndex(source, target, compartmentDict1, compartmentDict2)
				if index >= 0 {
					ranges[index] = append(ranges[index], distance)
				}
			}
		}
	}

	var buffer bytes.Buffer
	buffer.WriteString("\tmedian\tmean\tmin\tmax\n")

	for i := 0; i < 3; i++ {
		var name string
		if i < 2 {
			name = compartments[i].Name
		} else {
			name = "between"
		}
		max := stats.Round(stats.MaxFloat(ranges[i]), 0.001)
		mean := stats.Round(stats.MeanFloat(ranges[i]), 0.001)
		median := stats.Round(stats.MedianFloat(ranges[i]), 0.001)
		min := stats.Round(stats.MinFloat(ranges[i]), 0.001)
		buffer.WriteString(fmt.Sprintf("%s\t%.3f\t%.3f\t%.3f\t%.3f\n", name, median, mean, min, max))
	}
	afero.WriteFile(fs.Instance, outfile, buffer.Bytes(), 0644)
}
