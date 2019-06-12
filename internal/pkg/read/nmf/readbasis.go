package nmf

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"

	"github.com/knightjdr/cmgo/pkg/fs"
)

// Basis reads an NMF basis matrix.
func Basis(filename string) ([][]float64, []string, []string) {
	file, err := fs.Instance.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	reader := csv.NewReader(file)
	reader.Comma = ','
	reader.LazyQuotes = true

	// Read header.
	header, err := reader.Read()
	if err != nil {
		log.Fatalln(err)
	}
	columns := header[1:]

	matrix := make([][]float64, 0)
	rows := make([]string, 0)
	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}

		rows = append(rows, line[0])
		rowValues := make([]float64, len(columns))
		for i, value := range line[1:] {
			rowValues[i], _ = strconv.ParseFloat(value, 64)
		}
		matrix = append(matrix, rowValues)
	}

	return matrix, columns, rows
}
