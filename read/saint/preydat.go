package saint

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"

	"github.com/knightjdr/cmgo/fs"
)

// PreyDatRow defines the headers in a bait.dat file.
type PreyDatRow struct {
	Accession string
	Length    int
	Name      string
}

func mapPreyDatLine(line []string) PreyDatRow {
	accession := line[0]
	length, _ := strconv.Atoi(line[1])
	name := line[2]
	row := PreyDatRow{
		Accession: accession,
		Length:    length,
		Name:      name,
	}
	return row
}

// PreyDat reads a prey.dat file.
func PreyDat(filename string) []PreyDatRow {
	file, err := fs.Instance.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	reader := csv.NewReader(file)
	reader.Comma = '\t'
	reader.FieldsPerRecord = -1
	reader.LazyQuotes = true

	// Read file and filter by FDR.
	rows := make([]PreyDatRow, 0)
	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}

		row := mapPreyDatLine(line)
		rows = append(rows, row)
	}

	return rows
}
