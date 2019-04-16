package shared

import (
	"encoding/csv"
	"io"
	"log"

	"github.com/knightjdr/cmgo/fs"
)

func readRegions(filename string) map[string]map[string]bool {
	file, err := fs.Instance.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	reader := csv.NewReader(file)
	reader.Comma = '\t'
	reader.LazyQuotes = true

	regions := make(map[string]map[string]bool, 0)
	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}

		protein := line[0]
		region := line[1]
		if _, ok := regions[protein]; !ok {
			regions[protein] = make(map[string]bool, 0)
		}

		regions[protein][region] = true
	}
	return regions
}
