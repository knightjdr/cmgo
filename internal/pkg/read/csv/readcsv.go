// Package csv reads a csv file
package csv

import (
	"encoding/csv"
	"github.com/knightjdr/cmgo/pkg/fs"
	"log"
)

// Read a csv file.
func Read(filename string, header bool) *csv.Reader {
	file, err := fs.Instance.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	reader := csv.NewReader(file)
	reader.Comma = '\t'
	reader.FieldsPerRecord = -1
	reader.LazyQuotes = true

	if !header {
		_, err = reader.Read()
		if err != nil {
			log.Fatalln(err)
		}
	}

	return reader
}
