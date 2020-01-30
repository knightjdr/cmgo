package rankmetrics

import (
	"bytes"
	"fmt"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/knightjdr/cmgo/pkg/stats"
	"github.com/spf13/afero"
)

func writeSummary(summary map[int]*rankSummary, outfile string) {
	var buffer bytes.Buffer

	writeHeader(&buffer)
	writeBody(&buffer, summary)

	afero.WriteFile(fs.Instance, outfile, buffer.Bytes(), 0644)
}

func writeHeader(buffer *bytes.Buffer) {
	buffer.WriteString("prey rank\tlysines (mean)\tlysines (SD)\tgene with lysine data\tturnover rate (mean)\tturnover rate(SD)\tgenes with turnover data\n")
}

func writeBody(buffer *bytes.Buffer, summary map[int]*rankSummary) {
	numberOfRanks := len(summary)

	for i := 1; i <= numberOfRanks; i++ {
		lysineMean := stats.MeanInt(summary[i].Lysines)
		lysineSD := stats.SDInt(summary[i].Lysines)
		numberLysines := len(summary[i].Lysines)
		turnoverMean := stats.MeanFloat(summary[i].TurnoverRates)
		turnoverSD := stats.SDFloat(summary[i].TurnoverRates)
		numberTurnoverGenes := len(summary[i].TurnoverRates)
		buffer.WriteString(
			fmt.Sprintf(
				"%d\t%0.4f\t%0.4f\t%d\t%0.4f\t%0.4f\t%d\n",
				i,
				lysineMean,
				lysineSD,
				numberLysines,
				turnoverMean,
				turnoverSD,
				numberTurnoverGenes,
			),
		)
	}
}
