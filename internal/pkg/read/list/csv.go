package list

import (
	"encoding/csv"
	"io"
	"log"

	"github.com/knightjdr/cmgo/pkg/fs"
)

// CSV reads a csv file to a slice of maps with header as keys
func CSV(filename string, sep rune) []map[string]string {
	file, err := fs.Instance.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	reader := csv.NewReader(file)
	reader.Comma = sep
	reader.LazyQuotes = true

	// Create header.
	header, err := reader.Read()
	if err != nil {
		log.Fatalln(err)
	}

	// Read compartment information.
	data := make([]map[string]string, 0)
	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}

		parsedLine := make(map[string]string, len(header))
		for i, field := range header {
			parsedLine[field] = line[i]
		}

		data = append(data, parsedLine)
	}

	return data
}
