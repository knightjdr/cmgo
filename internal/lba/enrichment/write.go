package enrichment

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/knightjdr/cmgo/pkg/gprofiler"
	"github.com/knightjdr/cmgo/pkg/mapfunc"
	"github.com/spf13/afero"
)

func findRefseq(entrez string, refseqMap map[string]map[string]string) string {
	for refseq, terms := range refseqMap {
		if terms["Entrez"] == entrez {
			return refseq
		}
	}
	return ""
}

func write(enrichment map[string][]gprofiler.EnrichedTerm, refseqMap map[string]map[string]string, outfile string) {
	// Determine write order for genes.
	symbolToRefseq := make(map[string]string, 0)
	for refseq := range enrichment {
		if refseqMap[refseq]["Symbol"] != "" {
			symbolToRefseq[refseqMap[refseq]["Symbol"]] = refseq
		}
	}
	geneOrder := mapfunc.KeysStringString(symbolToRefseq)
	sort.Strings(geneOrder)

	// Write file
	var buffer bytes.Buffer
	buffer.WriteString("symbol\tEntrez\tRefseq\tUniProt\tGO ID\tGO term\tp-value\trecall\tprecision\tquery size\tterm size\tgenes\n")
	for _, symbol := range geneOrder {
		refseq := symbolToRefseq[symbol]
		for _, term := range enrichment[refseq] {
			// Convert query Entrez gene IDs to symbol
			queryGenes := make([]string, 0)
			for _, queryEntrezID := range term.Genes {
				refseqID := findRefseq(queryEntrezID, refseqMap)
				if refseqID != "" && refseqMap[refseqID]["Symbol"] != "" {
					queryGenes = append(queryGenes, refseqMap[refseqID]["Symbol"])
				}
			}

			buffer.WriteString(fmt.Sprintf(
				"%s\t%s\t%s\t%s\t%s\t%s\t%0.2e\t%0.4f\t%0.4f\t%d\t%d\t%s\n",
				refseqMap[refseq]["Symbol"],
				refseqMap[refseq]["Entrez"],
				refseq,
				refseqMap[refseq]["UniProt"],
				term.ID,
				term.Name,
				term.Pvalue,
				term.Recall,
				term.Precision,
				term.QuerySize,
				term.TermSize,
				strings.Join(queryGenes, ", "),
			))
		}
	}

	afero.WriteFile(fs.Instance, outfile, buffer.Bytes(), 0644)
}
