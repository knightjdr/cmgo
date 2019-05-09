package saint

import (
	"encoding/csv"
	"io"
	"log"

	"github.com/knightjdr/cmgo/fs"
)

// BaitDatRow defines the headers in a bait.dat file.
type BaitDatRow struct {
	ID      string
	Name    string
	Control bool
}

func mapBaitDatLine(line []string) BaitDatRow {
	id := line[0]
	name := line[1]
	var control bool
	if line[2] == "C" {
		control = true
	}
	row := BaitDatRow{
		Control: control,
		ID:      id,
		Name:    name,
	}
	return row
}

// BaitDat reads a bait.dat file.
func BaitDat(filename string) []BaitDatRow {
	file, err := fs.Instance.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	reader := csv.NewReader(file)
	reader.Comma = '\t'
	reader.FieldsPerRecord = -1
	reader.LazyQuotes = true

	// Read file and filter by FDR.
	rows := make([]BaitDatRow, 0)
	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}

		row := mapBaitDatLine(line)
		rows = append(rows, row)
	}

	return rows
}
