package nmfsafe

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	"github.com/knightjdr/cmgo/fs"
	"github.com/spf13/afero"
)

func outputTable(genes map[string]*localizationInfo, outfile string) {
	sortedGenes := make([]string, len(genes))
	i := 0
	for gene := range genes {
		sortedGenes[i] = gene
		i++
	}
	sort.Strings(sortedGenes)

	var buffer bytes.Buffer
	buffer.WriteString("gene\trank\tNMF term(s)\tNMF known?\tdomain\tSAFE term(s)\tSAFE known?\tNMF in SAFE?\tSAFE in NMF?\tconcordant\n")
	for _, gene := range sortedGenes {
		info := genes[gene]
		buffer.WriteString(
			fmt.Sprintf(
				"%s\t%d\t%s\t%t\t%d\t%s\t%t\t%t\t%t\t%t\n",
				gene,
				info.Rank,
				strings.Join(info.NMFterms, ", "),
				info.NMFknown,
				info.Domain,
				strings.Join(info.SAFEterms, ", "),
				info.SAFEknown,
				info.NMFinSAFE,
				info.SAFEinNMF,
				info.Concordant,
			),
		)
	}
	afero.WriteFile(fs.Instance, outfile, buffer.Bytes(), 0644)
}
