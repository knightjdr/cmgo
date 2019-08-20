package genes

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/knightjdr/cmgo/pkg/gprofiler"
	"github.com/spf13/afero"
)

func write(enrichment []gprofiler.EnrichedTerm, outfile string) {
	var buffer bytes.Buffer
	buffer.WriteString("GO ID\tGO term\tp-value\trecall\tprecision\tquery size\tterm size\tno. genes\tgenes\n")
	for _, term := range enrichment {
		buffer.WriteString(fmt.Sprintf(
			"%s\t%s\t%0.2e\t%0.4f\t%0.4f\t%d\t%d\t%d\t%s\n",
			term.ID,
			term.Name,
			term.Pvalue,
			term.Recall,
			term.Precision,
			term.QuerySize,
			term.TermSize,
			len(term.Genes),
			strings.Join(term.Genes, ", "),
		))
	}

	afero.WriteFile(fs.Instance, outfile, buffer.Bytes(), 0644)
}
