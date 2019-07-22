package localize

import (
	"encoding/csv"
	"io"
	"log"

	"github.com/knightjdr/cmgo/pkg/fs"
)

func readLocalizations(filename string) (map[string]string, []string) {
	file, err := fs.Instance.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	reader := csv.NewReader(file)
	reader.Comma = '\t'
	reader.LazyQuotes = true

	localizationOrder := make([]string, 0)
	localizations := make(map[string]string, 0)
	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}

		localizationOrder = append(localizationOrder, line[0])
		localizations[line[0]] = line[1]
	}

	return localizations, localizationOrder
}
