package preys

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/knightjdr/cmgo/pkg/gprofiler"
	"github.com/spf13/afero"
)

func goEnrichment(summary map[string]*preyInteraction, options parameters) {
	preys := orderPreys(summary)

	end := options.enrichmentLimit
	if len(preys) < end {
		end = len(preys)
	}

	body := gprofiler.RequestBody{
		Organism:      "hsapiens",
		Query:         preys[:end],
		Sources:       []string{"GO"},
		UserThreshold: 0.01,
	}

	service := gprofiler.Service{
		Body: body,
	}
	gprofiler.Fetch(&service)

	writeEnrichment(service.Result, options.outFileEnrichment)
}

func writeEnrichment(results []gprofiler.EnrichedTerm, outFile string) {
	var buffer bytes.Buffer

	writeEnrichmentHeader(&buffer)
	writeEnrichmentBody(&buffer, results)

	afero.WriteFile(fs.Instance, outFile, buffer.Bytes(), 0644)
}

func writeEnrichmentHeader(buffer *bytes.Buffer) {
	header := "term\tid\tsource\tquery size\tterm size\tintersection size\tprecision\trecall\tp-value\tgenes\n"
	buffer.WriteString(header)
}

func writeEnrichmentBody(buffer *bytes.Buffer, results []gprofiler.EnrichedTerm) {
	for _, term := range results {
		buffer.WriteString(
			fmt.Sprintf(
				"%s\t%s\t%s\t%d\t%d\t%d\t%0.3f\t%0.3f\t%e\t%s\n",
				term.Name,
				term.ID,
				term.Source,
				term.QuerySize,
				term.TermSize,
				term.IntersectionSize,
				term.Precision,
				term.Recall,
				term.Pvalue,
				strings.Join(term.Genes, ", "),
			),
		)
	}
}
