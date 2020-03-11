package goenrich

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/knightjdr/cmgo/pkg/gprofiler"
	"github.com/spf13/afero"
)

func writeEnrichment(enrichment map[string][]gprofiler.EnrichedTerm, options parameters) {
	var buffer bytes.Buffer
	buffer.WriteString("bait\tp-value\tterm size\tquery size\toverlap size\trecall\tprecision\tterm id\tdomain\tterm name\tintersection\n")
	for bait, terms := range enrichment {
		for _, term := range terms {
			buffer.WriteString(fmt.Sprintf(
				"%s\t%0.2e\t%d\t%d\t%d\t%0.4f\t%0.4f\t%s\t%s\t%s\t%s\n",
				bait,
				term.Pvalue,
				term.TermSize,
				term.QuerySize,
				len(term.Genes),
				term.Recall,
				term.Precision,
				term.ID,
				options.namespace,
				term.Name,
				strings.Join(term.Genes, ", "),
			))
		}
	}

	afero.WriteFile(fs.Instance, options.outFile, buffer.Bytes(), 0644)
}
