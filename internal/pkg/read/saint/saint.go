// Package saint reads and parses SAINT-related files.
package saint

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/knightjdr/cmgo/pkg/slice"
)

// Row defines the headers in a SAINT file.
type Row struct {
	Bait               string
	Prey               string
	PreyGene           string
	Spec               []float64
	AvgSpec            float64
	Control            []float64
	AvgP               float64
	FoldChange         float64
	FDR                float64
	NormalizedSpec     float64
	PreySequenceLength int
}

// SAINT is an array of Rows from a SAINT file.
type SAINT []Row

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
		Prey:       line[1],
		Spec:       spec,
	}

	if len(line) >= 21 && line[20] != "" {
		preylength, _ := strconv.ParseInt(line[20], 10, 64)
		row.PreySequenceLength = int(preylength)
	}
	return row
}

// Read reads a SAINT file and filters by FDR.
func Read(filename string, fdr float64, minBaits int) *SAINT {
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
	saint := new(SAINT)
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
			*saint = append(*saint, row)
		}
	}

	// Filter by bait number.
	saint.FilterByBaitNumber(minBaits)

	return saint
}
