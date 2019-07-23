// Package matrix reads matrix from txt file.
package matrix

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"

	"github.com/knightjdr/cmgo/pkg/fs"
)

func parseString(line []string) []float64 {
	values := make([]float64, len(line))
	for i, str := range line {
		values[i], _ = strconv.ParseFloat(str, 64)
	}
	return values
}

// Read a matrix from a tab-delimited file. The first row should
// contain column names and the first column row names.
func Read(filename string) ([]string, []string, [][]float64) {
	file, err := fs.Instance.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	reader := csv.NewReader(file)
	reader.Comma = '\t'
	reader.FieldsPerRecord = -1
	reader.LazyQuotes = true

	columns := make([]string, 0)
	data := make([][]float64, 0)
	header := false
	rows := make([]string, 0)

	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}
		if !header {
			columns = line[1:]
		} else {
			rows = append(rows, line[0])
			data = append(data, parseString(line[1:]))
		}
		header = true
	}

	return rows, columns, data
}
