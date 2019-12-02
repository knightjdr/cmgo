package countgo

import "github.com/knightjdr/cmgo/internal/pkg/read/csv"

func readGenes(filename string) []string {
	genes := make([]string, 0)

	reader := csv.Read(filename, false)
	for {
		eof, line := csv.Readline(reader)
		if eof {
			break
		}

		genes = append(genes, line[0])
	}

	return genes
}
