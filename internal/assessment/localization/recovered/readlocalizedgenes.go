package recovered

import "github.com/knightjdr/cmgo/internal/pkg/read/csv"

func readLocalizedGenes(filename, localizationID string) []string {
	genes := make([]string, 0)

	reader := csv.Read(filename, false)
	for {
		eof, line := csv.Readline(reader)
		if eof {
			break
		}

		if line[1] == localizationID {
			genes = append(genes, line[0])
		}
	}

	return genes
}
