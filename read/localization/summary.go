package localization

import (
	"encoding/csv"
	"io"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/knightjdr/cmgo/fs"
	"github.com/knightjdr/cmgo/slice"
)

// Summary contains domain/rank level information.
type Summary map[int]CompartmentSummary

// CompartmentSummary contains information about each NMF/SAFE compartment.
type CompartmentSummary struct {
	DisplayTerms []string
	GOid         []string
	GOterms      []string
	IC           []float64
}

var reBracket = regexp.MustCompile(`^"*\[(.+)\]"*`)

func mapSummaryLine(line []string) (int, CompartmentSummary) {
	comparment, _ := strconv.Atoi(line[0])

	// Display terms.
	displayTerms := make([]string, 0)
	matches := reBracket.FindStringSubmatch(line[2])
	if len(matches) > 0 {
		displayTerms = strings.Split(matches[1], ", ")
	}

	// GO IDs.
	goIDs := make([]string, 0)
	matches = reBracket.FindStringSubmatch(line[3])
	if len(matches) > 0 {
		goIDs = strings.Split(matches[1], ", ")
	}

	// GO terms.
	goTerms := make([]string, 0)
	matches = reBracket.FindStringSubmatch(line[1])
	if len(matches) > 0 {
		goTerms = strings.Split(matches[1], ", ")
	}

	// GO terms.
	ic := make([]float64, 0)
	matches = reBracket.FindStringSubmatch(line[5])
	if len(matches) > 0 {
		ic = slice.ConvertStringToFloat(strings.Split(matches[1], ", "))
	}

	summary := CompartmentSummary{
		DisplayTerms: displayTerms,
		GOid:         goIDs,
		GOterms:      goTerms,
		IC:           ic,
	}
	return comparment, summary
}

// SummaryFile reads an NMF or SAFE summary file.
func SummaryFile(filename string) Summary {
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

	// Read compartment information.
	summary := make(Summary, 0)
	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}

		compartment, compartmentSummary := mapSummaryLine(line)
		summary[compartment] = compartmentSummary
	}

	return summary
}
