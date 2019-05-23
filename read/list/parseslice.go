package list

import (
	"encoding/csv"
	"io"
	"log"

	"github.com/knightjdr/cmgo/fs"
)

// ParseSlice parses a txt list to a slice of strings
func ParseSlice(filename string) []string {
	file, err := fs.Instance.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	reader := csv.NewReader(file)
	reader.Comma = '\t'
	reader.FieldsPerRecord = -1
	reader.LazyQuotes = true

	// Read list.
	list := make([]string, 0)
	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}

		list = append(list, line[0])
	}

	return list
}
