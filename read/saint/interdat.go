package saint

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"

	"github.com/knightjdr/cmgo/fs"
)

// InterDatRow defines the headers in a bait.dat file.
type InterDatRow struct {
	ID   string
	Bait string
	Prey string
	Spec int
}

func mapInterDatLine(line []string) InterDatRow {
	id := line[0]
	bait := line[1]
	prey := line[2]
	spec, _ := strconv.Atoi(line[3])
	row := InterDatRow{
		ID:   id,
		Bait: bait,
		Prey: prey,
		Spec: spec,
	}
	return row
}

// InterDat reads a prey.dat file.
func InterDat(filename string) []InterDatRow {
	file, err := fs.Instance.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	reader := csv.NewReader(file)
	reader.Comma = '\t'
	reader.FieldsPerRecord = -1
	reader.LazyQuotes = true

	// Read file and filter by FDR.
	rows := make([]InterDatRow, 0)
	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}

		row := mapInterDatLine(line)
		rows = append(rows, row)
	}

	return rows
}
