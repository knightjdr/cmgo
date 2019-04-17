package heatmap

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"

	"github.com/knightjdr/cmgo/fs"
)

func readEnrichment(filename string, pValue float64) map[string]map[string]float64 {
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

	enrichment := make(map[string]map[string]float64, 0)
	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}

		currPValue, _ := strconv.ParseFloat(line[6], 64)
		if currPValue <= pValue {
			compartment := line[0]
			foldEnrichment, _ := strconv.ParseFloat(line[4], 64)
			term := line[1]
			if _, ok := enrichment[compartment]; !ok {
				enrichment[compartment] = make(map[string]float64, 0)
			}
			enrichment[compartment][term] = foldEnrichment
		}
	}
	return enrichment
}
