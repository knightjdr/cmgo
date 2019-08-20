package subset

import (
	"encoding/csv"
	"io"
	"log"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/knightjdr/cmgo/pkg/slice"
)

func parseHeatmap(filename string, genes []string) ([][]string, string) {
	geneDict := slice.Dict(genes)

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

	lines := make([][]string, 0)
	var params string
	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}

		if params == "" {
			params = line[3]
		}

		if _, ok := geneDict[line[0]]; ok {
			if _, ok := geneDict[line[1]]; ok {
				lines = append(lines, line[:3])
			}
		}

	}

	return lines, params
}
