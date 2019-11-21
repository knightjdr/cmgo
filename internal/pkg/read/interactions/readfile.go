package interactions

import (
	"sort"

	"github.com/knightjdr/cmgo/internal/pkg/read/csv"
	"github.com/knightjdr/cmgo/pkg/slice"
)

type parse func([]string, string) (string, string)

func readFile(filename, species string, parseFunc parse) map[string][]string {
	interactions := make(map[string][]string)

	if filename == "" {
		return interactions
	}

	reader := csv.Read(filename, false)

	for {
		eof, line := csv.Readline(reader)
		if eof {
			break
		}

		source, target := parseFunc(line, species)
		if source != "" && target != "" {
			allocateMemory(&interactions, source, target)
			interactions[source] = append(interactions[source], target)
			interactions[target] = append(interactions[target], source)
		}
	}

	for gene := range interactions {
		interactions[gene] = slice.UniqueStrings(interactions[gene])
		sort.Strings(interactions[gene])
	}

	return interactions
}

func allocateMemory(interactions *map[string][]string, source, target string) {
	if _, ok := (*interactions)[source]; !ok {
		(*interactions)[source] = make([]string, 0)
	}
	if _, ok := (*interactions)[target]; !ok {
		(*interactions)[target] = make([]string, 0)
	}
}
