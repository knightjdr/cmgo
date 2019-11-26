package transmembrane

import (
	"bytes"
	"fmt"

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
	buffer.WriteString("prey\tlocalization\tcytosolic baits\tlumenal baits\tcytosolic score\tlumenal score\n")
}

func writeBody(buffer *bytes.Buffer, summary map[string]preySummary) {
	for prey, data := range summary {
		buffer.WriteString(
			fmt.Sprintf(
				"%s\t%s\t%d\t%d\t%0.3f\t%0.3f\n",
				prey,
				data.localization,
				data.cytosolicBaits,
				data.lumenalBaits,
				data.cytosolicScore,
				data.lumenalScore,
			),
		)
	}
}
