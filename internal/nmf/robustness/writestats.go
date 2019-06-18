package robustness

import (
	"bytes"
	"fmt"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/spf13/afero"
)

func writeStats(stats [][]meanSD, rankNames []string, percentiles []float64, outfile string) {
	var buffer bytes.Buffer
	buffer.WriteString("rank\tpercentile\tmean\tSD\n")
	for rankIndex, rankData := range stats {
		rankName := rankNames[rankIndex]
		for percentileIndex, percentileData := range rankData {
			percentileValue := percentiles[percentileIndex]
			buffer.WriteString(fmt.Sprintf("%s\t%0.2f\t%0.5f\t%0.5f\n", rankName, percentileValue, percentileData.Mean, percentileData.SD))
		}
	}
	afero.WriteFile(fs.Instance, outfile, buffer.Bytes(), 0644)
}
