package svg

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"

	"github.com/knightjdr/cmgo/pkg/fs"
)

type coordinate struct {
	X float64
	Y float64
}

func readCoordinates(filename string) map[string]coordinate {
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

	// Read compartment information.
	coordinates := make(map[string]coordinate, 0)
	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}

		gene := line[0]
		x, _ := strconv.ParseFloat(line[1], 64)
		y, _ := strconv.ParseFloat(line[2], 64)

		coordinates[gene] = coordinate{
			X: x,
			Y: y,
		}
	}

	return coordinates
}
