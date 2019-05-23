// Package nmf reads NMF output files and assess localizations
package nmf

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"

	"github.com/knightjdr/cmgo/fs"
)

// NMFlocalization contains localization information for all genes.
type NMFlocalization map[string]GeneLocalization

// GeneLocalization contains the NMF rank and score for a gene.
type GeneLocalization struct {
	Compartment int
	Score       float64
}

func mapLocalizationLine(line []string) (string, GeneLocalization) {
	gene := line[0]
	rank, _ := strconv.Atoi(line[1])
	score, _ := strconv.ParseFloat(line[2], 64)

	localization := GeneLocalization{
		Compartment: rank,
		Score:       score,
	}

	return gene, localization
}

// Localization reads NMF rank localizations and scores.
func Localization(filename string) NMFlocalization {
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

	// Read compartment information.
	localization := make(NMFlocalization, 0)
	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}

		gene, geneLocalization := mapLocalizationLine(line)
		localization[gene] = geneLocalization
	}

	return localization
}
