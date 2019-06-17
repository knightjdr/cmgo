package robustness

import (
	"bytes"
	"fmt"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/knightjdr/cmgo/pkg/stats"
	"github.com/spf13/afero"
)

func writeData(data [][][]float64, rankNames []string, percentiles []float64, replicates int, outfile string) {
	var buffer bytes.Buffer
	buffer.WriteString("rank\tpercentile\treplicate\tRBD\tmean\tSD\n")
	for rankIndex, rankData := range data {
		rankName := rankNames[rankIndex]
		for percentileIndex, percentileData := range rankData {
			percentileValue := percentiles[percentileIndex]
			mean := stats.MeanFloat(percentileData)
			sd := stats.SDFloat(percentileData)
			for replicate := 0; replicate < replicates; replicate++ {
				buffer.WriteString(fmt.Sprintf("%s\t%0.2f\t%d\t%0.5f", rankName, percentileValue, replicate+1, percentileData[replicate]))
				if replicate == 0 {
					buffer.WriteString(fmt.Sprintf("\t%0.4f\t%0.4f", mean, sd))
				}
				buffer.WriteString("\n")
			}
		}
	}
	afero.WriteFile(fs.Instance, outfile, buffer.Bytes(), 0644)
}
