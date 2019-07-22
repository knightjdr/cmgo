package localize

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/spf13/afero"
)

func writeProfile(enrichment map[string][]Enrichment, orderIDs []string, outfile string) {
	// Determine write order for genes.
	geneOrder := make([]string, 0)
	for gene := range enrichment {
		geneOrder = append(geneOrder, gene)
	}
	sort.Strings(geneOrder)

	var buffer bytes.Buffer
	buffer.WriteString("gene")
	// Write column headers.
	for _, id := range orderIDs {
		buffer.WriteString(fmt.Sprintf("\t%s", id))
	}
	buffer.WriteString("\n")

	for _, gene := range geneOrder {
		profile := make([]string, len(orderIDs))
		for i, id := range orderIDs {
			for _, term := range enrichment[gene] {
				if term.ID == id {
					profile[i] = fmt.Sprintf("%0.4f", term.Precision)
					break
				}
			}
			if profile[i] == "" {
				profile[i] = "0.0000"
			}
		}
		buffer.WriteString(fmt.Sprintf("%s\t%s\n", gene, strings.Join(profile, string('\t'))))
	}

	afero.WriteFile(fs.Instance, outfile, buffer.Bytes(), 0644)
}
