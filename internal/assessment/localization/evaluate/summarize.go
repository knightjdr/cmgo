package evaluate

import (
	"bytes"
	"fmt"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/spf13/afero"
)

func summarize(genes map[string]bool, outfile string) {
	totalGenes := len(genes)

	knownCount := 0
	for _, known := range genes {
		if known {
			knownCount++
		}
	}

	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("total genes: %d\n", totalGenes))
	buffer.WriteString(fmt.Sprintf("known genes: %d\n", knownCount))
	buffer.WriteString(fmt.Sprintf("fraction known: %0.4f\n", float64(knownCount)/float64(totalGenes)))
	afero.WriteFile(fs.Instance, outfile, buffer.Bytes(), 0644)
}
