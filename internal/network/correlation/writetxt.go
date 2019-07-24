package correlation

import (
	"bytes"
	"fmt"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/spf13/afero"
)

func writeTXT(genes []string, pairs map[string][]edgePair, outfile string) {
	var buffer bytes.Buffer
	buffer.WriteString("source\ttarget\tweight\n")
	for _, gene := range genes {
		for _, edge := range pairs[gene] {
			buffer.WriteString(fmt.Sprintf("%s\t%s\t%0.4f\n", gene, edge.Target, edge.Weight))
		}
	}
	afero.WriteFile(fs.Instance, outfile, buffer.Bytes(), 0644)
}
