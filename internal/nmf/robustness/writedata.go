package robustness

import (
	"bytes"
	"fmt"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/spf13/afero"
)

func writeData(data [][][]float64, rankNames []string, percentiles []float64, replicates int, outfile string) {
	var buffer bytes.Buffer
	buffer.WriteString("rank\tpercentile\treplicate\tRBD\n")
	for rankIndex, rankData := range data {
		rankName := rankNames[rankIndex]
		for percentileIndex, percentileData := range rankData {
			percentileValue := percentiles[percentileIndex]
			for replicate := 0; replicate < replicates; replicate++ {
				buffer.WriteString(fmt.Sprintf("%s\t%0.2f\t%d\t%0.5f\n", rankName, percentileValue, replicate+1, percentileData[replicate]))
			}
		}
	}
	afero.WriteFile(fs.Instance, outfile, buffer.Bytes(), 0644)
}
