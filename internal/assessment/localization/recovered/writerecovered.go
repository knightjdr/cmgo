package recovered

import (
	"bytes"
	"fmt"
	"sort"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/spf13/afero"
)

func writeRecovered(recoveredGeneSummary map[string]bool, outFile string) {
	var buffer bytes.Buffer

	writeHeader(&buffer)
	writeBody(&buffer, recoveredGeneSummary)

	afero.WriteFile(fs.Instance, outFile, buffer.Bytes(), 0644)
}

func writeHeader(buffer *bytes.Buffer) {
	buffer.WriteString("gene\trecovered\n")
}

func writeBody(buffer *bytes.Buffer, summary map[string]bool) {
	outputOrder := orderGenes(summary)

	for _, gene := range outputOrder {
		buffer.WriteString(fmt.Sprintf("%s\t%t\n", gene, summary[gene]))
	}
}

func orderGenes(summary map[string]bool) []string {
	keys := make([]string, len(summary))

	i := 0
	for key := range summary {
		keys[i] = key
		i++
	}

	sort.Strings(keys)
	return keys
}
