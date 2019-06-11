// Package safe read SAFE output files
package safe

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/knightjdr/cmgo/pkg/fs"
)

// SAFElocalization contains localization information for all genes.
type SAFElocalization map[string]GeneLocalization

// GeneLocalization contains the SAFE domain and score for a gene.
type GeneLocalization struct {
	Compartment int
	Score       float64
}

func mapLocalizationLine(line []string) (string, GeneLocalization) {
	gene := line[1]
	domain, _ := strconv.Atoi(line[2])
	score, _ := strconv.ParseFloat(line[3], 64)

	localization := GeneLocalization{
		Compartment: domain,
		Score:       score,
	}

	return gene, localization
}

// Localization reads SAFE domain localizations and scores.
func Localization(filename string) SAFElocalization {
	file, err := fs.Instance.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	reader := csv.NewReader(file)
	reader.Comma = '\t'
	reader.FieldsPerRecord = -1
	reader.LazyQuotes = true

	// Read compartment information.
	localization := make(SAFElocalization, 0)
	start := false
	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}

		if start {
			gene, geneLocalization := mapLocalizationLine(line)
			localization[gene] = geneLocalization
		}

		if strings.HasPrefix(line[0], "Node label") {
			start = true
		}
	}

	return localization
}
