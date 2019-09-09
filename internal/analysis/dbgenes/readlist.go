package dbgenes

import (
	"bufio"
	"log"
	"strings"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/knightjdr/cmgo/pkg/slice"
)

func readList(filename string) []string {
	file, err := fs.Instance.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	genes := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "9606") {
			cells := strings.Split(line, "\t")
			if cells[2] != "" {
				genes = append(genes, cells[2])
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	genes = slice.UniqueStrings(genes)

	return genes
}
