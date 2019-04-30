// Package read reads and parses common files
package read

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/knightjdr/cmgo/fs"
	"github.com/knightjdr/cmgo/slice"
)

// SaintRow defines the headers in a SAINT file.
type SaintRow struct {
	Bait               string
	PreyGene           string
	Spec               []float64
	AvgSpec            float64
	Control            []float64
	AvgP               float64
	FoldChange         float64
	FDR                float64
	PreySequenceLength int64
}

func mapSaintLine(line []string) SaintRow {
	avgspec, _ := strconv.ParseFloat(line[5], 64)
	avgp, _ := strconv.ParseFloat(line[8], 64)
	control := slice.ConvertStringToFloat(strings.Split(line[7], "|"))
	foldchange, _ := strconv.ParseFloat(line[14], 64)
	fdr, _ := strconv.ParseFloat(line[15], 64)
	spec := slice.ConvertStringToFloat(strings.Split(line[3], "|"))
	row := SaintRow{
		AvgP:       avgp,
		AvgSpec:    avgspec,
		Bait:       line[0],
		Control:    control,
		FDR:        fdr,
		FoldChange: foldchange,
		PreyGene:   line[2],
		Spec:       spec,
	}

	if len(line) >= 21 && line[20] != "" {
		preylength, _ := strconv.ParseInt(line[20], 10, 64)
		row.PreySequenceLength = preylength
	}
	return row
}

func filterBaits(rows []SaintRow, minBaits int) []SaintRow {
	// Filter by bait number.
	if minBaits <= 1 {
		return rows
	}

	// Count how many times a prey occurs
	preys := make(map[string]int, 0)
	for _, row := range rows {
		preys[row.PreyGene]++
	}

	filteredRows := make([]SaintRow, 0)
	for _, row := range rows {
		if preys[row.PreyGene] >= minBaits {
			filteredRows = append(filteredRows, row)
		}
	}
	return filteredRows
}

// Saint reads a SAINT file and filters by FDR.
func Saint(filename string, fdr float64, minBaits int) []SaintRow {
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
	rows := make([]SaintRow, 0)
	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}

		row := mapSaintLine(line)
		if row.FDR <= fdr {
			rows = append(rows, row)
		}
	}

	// Filter by bait number.
	rows = filterBaits(rows, minBaits)

	return rows
}
