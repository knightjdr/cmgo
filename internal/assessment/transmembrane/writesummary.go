package transmembrane

import (
	"bytes"
	"fmt"
	"sort"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/spf13/afero"
)

func writeSummary(summary map[string]preySummary, outfile string) {
	var buffer bytes.Buffer

	writeHeader(&buffer)
	writeBody(&buffer, summary)

	afero.WriteFile(fs.Instance, outfile, buffer.Bytes(), 0644)
}

func writeHeader(buffer *bytes.Buffer) {
	buffer.WriteString("prey\tuniprot\tlocalization\tcytosolic baits\t" +
		"lumenal baits\tcytosolic score\tlumenal score\tmax cytosolic score\tmax lumenal score\t" +
		"sequence length\tcytosolic fraction\tlumenal fraction\t" +
		"NMF orientation score\tsequence orientation score\n",
	)
}

func writeBody(buffer *bytes.Buffer, summary map[string]preySummary) {
	outputOrder := orderPreys(summary)

	for _, prey := range outputOrder {
		data := summary[prey]
		buffer.WriteString(
			fmt.Sprintf(
				"%s\t%s\t%s\t%d\t%d\t%0.3f\t%0.3f\t%0.3f\t%0.3f\t%d\t%0.4f\t%0.4f\t%0.4f\t%0.4f\n",
				prey,
				data.uniprotID,
				data.localization,
				data.cytosolicBaits,
				data.lumenalBaits,
				data.cytosolicScore,
				data.lumenalScore,
				data.maxCytosolicScore,
				data.maxLumenalScore,
				data.length,
				data.cytosolicFraction,
				data.lumenalFraction,
				(data.cytosolicScore/data.maxCytosolicScore)-(data.lumenalScore/data.maxLumenalScore),
				data.cytosolicFraction-data.lumenalFraction,
			),
		)
	}
}

func orderPreys(summary map[string]preySummary) []string {
	keys := make([]string, len(summary))

	i := 0
	for key := range summary {
		keys[i] = key
		i++
	}

	sort.Strings(keys)
	return keys
}
