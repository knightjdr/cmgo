package hydropathy

import (
	"github.com/knightjdr/cmgo/internal/pkg/read/database"
)

func sequenceMapping(data []database.Fasta) (map[string]string, map[string]string, map[string]string) {
	entrezSequenceMap := make(map[string]string, 0)
	refseqSequenceMap := make(map[string]string, 0)
	refseqEntrez := make(map[string]string, 0)

	for _, entry := range data {
		entrezSequenceMap[entry.Entrez] = entry.Sequence
		refseqSequenceMap[entry.Refseq] = entry.Sequence
		refseqEntrez[entry.Refseq] = entry.Entrez
	}

	return entrezSequenceMap, refseqSequenceMap, refseqEntrez
}
