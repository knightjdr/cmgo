package overlap

import (
	"encoding/csv"
	"io"
	"log"
	"sort"
	"strconv"

	"github.com/knightjdr/cmgo/pkg/fs"
)

func readSimilarity(filename string) map[string]map[string]float64 {
	file, err := fs.Instance.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	reader := csv.NewReader(file)
	reader.Comma = '\t'
	reader.LazyQuotes = true

	// Skip header.
	_, err = reader.Read()
	if err != nil {
		log.Fatalln(err)
	}

	similarity := make(map[string]map[string]float64, 0)
	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}

		names := []string{line[0], line[1]}
		if names[0] != names[1] {
			sort.Strings(names)
			source := names[0]
			target := names[1]

			if _, ok := similarity[source]; !ok {
				similarity[source] = make(map[string]float64, 0)
			}

			similarity[source][target], _ = strconv.ParseFloat(line[2], 64)
		}
	}

	return similarity
}
