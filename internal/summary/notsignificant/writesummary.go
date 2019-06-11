package notsignificant

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/spf13/afero"
)

func writeSummary(summary map[string]*preySummary, outfile string) {
	preys := make([]string, len(summary))
	i := 0
	for prey := range summary {
		preys[i] = prey
		i++
	}
	sort.Strings(preys)

	var buffer bytes.Buffer
	buffer.WriteString("prey\tbaits\tmaximum spectral count\taverage control count\tbest FDR\n")
	for _, prey := range preys {
		baits := strings.Join(summary[prey].baits, ", ")
		ctrlAvg := fmt.Sprintf("%.2f", summary[prey].ctrlAvg)
		bestFDR := fmt.Sprintf("%.2f", summary[prey].bestFDR)
		maxSpec := fmt.Sprintf("%.2f", summary[prey].maxSpec)
		buffer.WriteString(fmt.Sprintf("%s\t%s\t%s\t%s\t%s\n", prey, baits, maxSpec, ctrlAvg, bestFDR))
	}
	afero.WriteFile(fs.Instance, outfile, buffer.Bytes(), 0644)
}
