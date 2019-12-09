package localization

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/knightjdr/cmgo/pkg/fs"
)

// ExpectedLocalization contains the ID and localization terms for a bait
type ExpectedLocalization struct {
	GoID  []string
	ID    int
	Terms []string
}

// ExpectedLocalizations is a map of bait name to ID, GO ID and localization terms.
type ExpectedLocalizations map[string]ExpectedLocalization

func mapExpectedLine(line []string) (string, ExpectedLocalization) {
	bait := line[1]
	goids := line[3]
	id, _ := strconv.Atoi(line[0])
	termString := strings.Replace(line[2], "\"", "", -1)

	localization := ExpectedLocalization{
		GoID:  strings.Split(goids, ";"),
		ID:    id,
		Terms: strings.Split(termString, ";"),
	}
	return bait, localization
}

// Expected reads a list of expected localizations
func Expected(filename string) ExpectedLocalizations {
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

	// Read localizations.
	expected := make(ExpectedLocalizations, 0)
	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}

		bait, info := mapExpectedLine(line)
		expected[bait] = info
	}

	return expected
}
