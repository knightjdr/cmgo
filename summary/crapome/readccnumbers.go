package crapome

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/knightjdr/cmgo/fs"
)

func readCC(filename string) map[int]string {
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

	// Read file and filter by FDR.
	idMap := make(map[int]string, 0)
	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}

		id, _ := strconv.Atoi(strings.Split(line[0], "_")[0])
		idMap[id] = line[1]
	}

	return idMap
}
