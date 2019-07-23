// Package tsne reads t-SNE coordinate file.
package tsne

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"

	"github.com/knightjdr/cmgo/pkg/fs"
)

// Coordinate are X and Y coordinates for a 2D space.
type Coordinate struct {
	X float64
	Y float64
}

// Coordinates reads t-SNE coordinates.
func Coordinates(filename string) map[string]Coordinate {
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
	coordinates := make(map[string]Coordinate, 0)
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

		coordinates[gene] = Coordinate{
			X: x,
			Y: y,
		}
	}

	return coordinates
}
