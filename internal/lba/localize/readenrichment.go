package localize

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"

	"github.com/knightjdr/cmgo/pkg/fs"
)

// Enrichment values relevant to localizing genes.
type Enrichment struct {
	Entrez    string
	ID        string
	Precision float64
	Pvalue    float64
}

func mapEnrichmentLine(line []string) Enrichment {
	precision, _ := strconv.ParseFloat(line[7], 64)
	pvalue, _ := strconv.ParseFloat(line[5], 64)
	return Enrichment{
		Entrez:    line[0],
		ID:        line[3],
		Precision: precision,
		Pvalue:    pvalue,
	}
}

func readEnrichment(filename string) map[string][]Enrichment {
	file, err := fs.Instance.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	reader := csv.NewReader(file)
	reader.Comma = '\t'
	reader.FieldsPerRecord = -1
	reader.LazyQuotes = true

	// Skip header.
	_, err = reader.Read()
	if err != nil {
		log.Fatalln(err)
	}

	// Read file and filter by FDR.
	enrichments := make(map[string][]Enrichment, 0)
	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}

		gene := line[0]
		enrichments[gene] = append(enrichments[gene], mapEnrichmentLine(line[1:]))
	}

	return enrichments

}
