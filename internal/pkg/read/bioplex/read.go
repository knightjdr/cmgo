// Package bioplex reads a Bioplex interaction file
package bioplex

import (
	"encoding/csv"
	"io"
	"log"

	"github.com/knightjdr/cmgo/pkg/fs"
)

// Interaction details for an entry in BioPlex
type Interaction struct {
	Source string
	Target string
}

// Entry contains Entrez, gene symbol and UniProt IDs for each interaction.
type Entry struct {
	Entrez  Interaction
	Symbol  Interaction
	UniProt Interaction
}

func mapBioplexLine(line []string) Entry {
	entrezSource := line[0]
	entrezTarget := line[1]
	uniprotSource := line[2]
	uniprotTarget := line[3]
	symbolSource := line[4]
	symbolTarget := line[5]
	entry := Entry{
		Entrez: Interaction{
			Source: entrezSource,
			Target: entrezTarget,
		},
		Symbol: Interaction{
			Source: symbolSource,
			Target: symbolTarget,
		},
		UniProt: Interaction{
			Source: uniprotSource,
			Target: uniprotTarget,
		},
	}

	return entry
}

// Read a BioPlex interation file.
func Read(filename string) []Entry {
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
	entries := make([]Entry, 0)
	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}

		entry := mapBioplexLine(line)
		entries = append(entries, entry)
	}

	return entries
}
