// Package read reads and parses common files
package read

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"

	"github.com/knightjdr/cmgo/fs"
)

// SaintRow defines the headers in a SAINT file
type SaintRow struct {
	Bait               string
	PreyGene           string
	AvgSpec            float64
	Control            string
	AvgP               float64
	FoldChange         float64
	FDR                float64
	PreySequenceLength int64
}

func mapLine(line []string) SaintRow {
	avgspec, _ := strconv.ParseFloat(line[5], 64)
	avgp, _ := strconv.ParseFloat(line[8], 64)
	foldchange, _ := strconv.ParseFloat(line[14], 64)
	fdr, _ := strconv.ParseFloat(line[15], 64)
	row := SaintRow{
		Bait:       line[0],
		PreyGene:   line[2],
		AvgSpec:    avgspec,
		Control:    line[7],
		AvgP:       avgp,
		FoldChange: foldchange,
		FDR:        fdr,
	}

	if len(line) >= 21 && line[20] != "" {
		preylength, _ := strconv.ParseInt(line[20], 10, 64)
		row.PreySequenceLength = preylength
	}
	return row
}

// Saint reads a SAINT file and filters by FDR
func Saint(filename string, filter float64) []SaintRow {
	file, err := fs.Instance.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	reader := csv.NewReader(file)
	reader.Comma = '\t'
	reader.LazyQuotes = true

	// Skip header.
	_, err = reader.Read()
	if err != nil {
		log.Fatalln(err)
	}

	rows := make([]SaintRow, 0)
	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}

		row := mapLine(line)

		if row.FDR <= filter {
			rows = append(rows, row)
		}
	}
	return rows
}
