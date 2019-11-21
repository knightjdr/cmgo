package csv

import (
	"encoding/csv"
	"io"
	"log"
)

// Readline reads a line from a csv file
func Readline(reader *csv.Reader) (bool, []string) {
	line, err := reader.Read()
	if err != nil {
		if err == io.EOF {
			return true, []string{}
		}
		log.Fatalln(err)
	}
	return false, line
}
