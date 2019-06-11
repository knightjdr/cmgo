package nmfsafe

import (
	"bytes"
	"fmt"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/spf13/afero"
)

func summarize(genes map[string]*localizationInfo, outfile string) {
	totalGenes := len(genes)

	// Count number of genes in NMF and SAFE (nmfGenes/safeGenes), as well
	// as genes assigned an actual GO term (nmfAssigned/safeAssigned)
	bothAssigned := 0
	concordant := 0
	nmfAssigned := 0
	nmfGenes := 0
	nmfKnown := 0
	safeAssigned := 0
	safeGenes := 0
	safeKnown := 0
	for _, info := range genes {
		if info.Concordant {
			concordant++
		}
		if info.Domain > 0 {
			safeGenes++
		}
		if len(info.NMFids) > 0 {
			nmfAssigned++
		}
		if info.NMFknown {
			nmfKnown++
		}
		if info.Rank > 0 {
			nmfGenes++
		}
		if len(info.SAFEids) > 0 {
			bothAssigned++
			safeAssigned++
		}
		if info.SAFEknown {
			safeKnown++
		}
	}

	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Total genes:\t%d\n", totalGenes))
	buffer.WriteString(fmt.Sprintf("Genes with NMF and SAFE assignment:\t%d\n", bothAssigned))
	buffer.WriteString(fmt.Sprintf("Concordant genes:\t%d\t%0.2f\n\n", concordant, (float64(concordant)/float64(bothAssigned))*100))

	buffer.WriteString("analysis\ttotal genes\tassigned term\t% assigned\tknown\t% known\n")
	buffer.WriteString(
		fmt.Sprintf(
			"NMF\t%d\t%d\t%0.2f\t%d\t%0.2f\n",
			nmfGenes,
			nmfAssigned,
			(float64(nmfAssigned)/float64(nmfGenes))*100,
			nmfKnown,
			(float64(nmfKnown)/float64(nmfGenes))*100,
		),
	)
	buffer.WriteString(
		fmt.Sprintf(
			"SAFE\t%d\t%d\t%0.2f\t%d\t%0.2f\n",
			safeGenes,
			safeAssigned,
			(float64(safeAssigned)/float64(safeGenes))*100,
			safeKnown,
			(float64(safeKnown)/float64(safeGenes))*100,
		),
	)
	afero.WriteFile(fs.Instance, outfile, buffer.Bytes(), 0644)
}
