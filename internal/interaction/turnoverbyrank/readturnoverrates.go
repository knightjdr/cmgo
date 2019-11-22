package turnoverbyrank

import (
	"github.com/knightjdr/cmgo/internal/pkg/read/csv"
	"strconv"
	"strings"
)

func readTurnoverRates(turnoverFile string) map[string]float64 {
	reader := csv.Read(turnoverFile, false)

	turnoverRatePerGene := make(map[string]float64, 0)
	for {
		eof, line := csv.Readline(reader)
		if eof {
			break
		}

		genes, halflife := parseLine(line)
		for _, gene := range genes {
			turnoverRatePerGene[gene] = halflife
		}
	}

	return turnoverRatePerGene
}

func parseLine(line []string) ([]string, float64) {
	genes := strings.Split(line[2], ";")
	halflife, _ := strconv.ParseFloat(line[15], 64)
	return genes, halflife
}
