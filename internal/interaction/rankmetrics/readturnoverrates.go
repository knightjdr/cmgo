package rankmetrics

import (
	"strconv"
	"strings"

	"github.com/knightjdr/cmgo/internal/pkg/read/csv"
)

func readTurnoverRates(data analysis) map[string]float64 {
	reader := csv.Read(data.parameters.turnoverFile, false)

	turnoverRatePerGene := make(map[string]float64, 0)
	for {
		eof, line := csv.Readline(reader)
		if eof {
			break
		}

		uniprotIDs, halflife := parseLine(line)
		for _, uniprotID := range uniprotIDs {
			if _, ok := data.uniprotMapping[uniprotID]; ok {
				gene := data.uniprotMapping[uniprotID]
				turnoverRatePerGene[gene] = halflife
			}
		}
	}

	return turnoverRatePerGene
}

func parseLine(line []string) ([]string, float64) {
	uniprotIDs := strings.Split(line[1], ";")
	halflife, _ := strconv.ParseFloat(line[15], 64)
	return uniprotIDs, halflife
}
