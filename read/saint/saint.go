// Package saint reads and parses SAINT-related files.
package saint

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/knightjdr/cmgo/fs"
	"github.com/knightjdr/cmgo/slice"
)

// Row defines the headers in a SAINT file.
type Row struct {
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

func mapSaintLine(line []string) Row {
	avgspec, _ := strconv.ParseFloat(line[5], 64)
	avgp, _ := strconv.ParseFloat(line[8], 64)
	control := slice.ConvertStringToFloat(strings.Split(line[7], "|"))
	foldchange, _ := strconv.ParseFloat(line[14], 64)
	fdr, _ := strconv.ParseFloat(line[15], 64)
	spec := slice.ConvertStringToFloat(strings.Split(line[3], "|"))
	row := Row{
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

func filterBaits(rows []Row, minBaits int) []Row {
	// Filter by bait number.
	if minBaits <= 1 {
		return rows
	}

	// Count how many times a prey occurs
	preys := make(map[string]int, 0)
	for _, row := range rows {
		preys[row.PreyGene]++
	}

	filteredRows := make([]Row, 0)
	for _, row := range rows {
		if preys[row.PreyGene] >= minBaits {
			filteredRows = append(filteredRows, row)
		}
	}
	return filteredRows
}

// Read reads a SAINT file and filters by FDR.
func Read(filename string, fdr float64, minBaits int) []Row {
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
	rows := make([]Row, 0)
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
