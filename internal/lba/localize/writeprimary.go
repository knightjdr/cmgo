package localize

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/spf13/afero"
)

func writePrimary(enrichment map[string][]Enrichment, goIDs map[string]string, outfile string) {
	// Determine write order for genes.
	geneOrder := make([]string, 0)
	for gene := range enrichment {
		geneOrder = append(geneOrder, gene)
	}
	sort.Strings(geneOrder)

	var buffer bytes.Buffer
	buffer.WriteString("gene\tterm(s)\tID(s)\tprecision\n")
	for _, gene := range geneOrder {
		// Get top GO IDs by p-value
		bestPvalue := float64(1.0)
		precision := make([]string, 0)
		primary := make([]string, 0)
		for _, term := range enrichment[gene] {
			if term.Pvalue < bestPvalue {
				bestPvalue = term.Pvalue
				precision = []string{fmt.Sprintf("%0.2f", term.Precision)}
				primary = []string{term.ID}
			} else if term.Pvalue == bestPvalue {
				precision = append(precision, fmt.Sprintf("%0.2f", term.Precision))
				primary = append(primary, term.ID)
			}
		}

		// Get term names for GO IDs.
		terms := make([]string, len(primary))
		for i, ID := range primary {
			terms[i] = goIDs[ID]
		}

		buffer.WriteString(fmt.Sprintf("%s\t%s\t%s\t%s\n", gene, strings.Join(terms, ";"), strings.Join(primary, ";"), strings.Join(precision, ";")))
	}

	afero.WriteFile(fs.Instance, outfile, buffer.Bytes(), 0644)
}
