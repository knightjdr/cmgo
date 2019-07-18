package localization

import (
	"encoding/csv"
	"io"
	"log"
	"strings"

	"github.com/knightjdr/cmgo/pkg/fs"
)

func mapPreyLine(line []string) (string, map[string]string) {
	prey := line[0]
	ids := make([]string, 0)
	terms := make([]string, 0)
	if len(line) >= 3 {
		ids = strings.Split(line[2], ";")
		terms = strings.Split(line[1], ";")
	}

	localization := make(map[string]string, 0)
	for i, id := range ids {
		localization[strings.TrimSpace(id)] = strings.TrimSpace(terms[i])
	}
	return prey, localization
}

// Prey reads a list of assigned localizations for preys.
func Prey(filename string) map[string]map[string]string {
	file, err := fs.Instance.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	reader := csv.NewReader(file)
	reader.Comma = '\t'
	reader.FieldsPerRecord = -1
	reader.LazyQuotes = true

	// Skip header.
	_, err = reader.Read()
	if err != nil {
		log.Fatalln(err)
	}

	localizations := make(map[string]map[string]string, 0)
	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}

		prey, terms := mapPreyLine(line)
		localizations[prey] = terms
	}

	return localizations
}
