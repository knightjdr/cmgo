package rankmetrics

import (
	"strings"

	"github.com/knightjdr/cmgo/internal/pkg/read/database"
)

func countLysines(fasta string) map[string]int {
	db := database.Read(fasta, true)

	lysines := make(map[string]int, len(db))
	for _, entry := range db {
		lysines[entry.Symbol] = strings.Count(entry.Sequence, "K")
	}

	return lysines
}
