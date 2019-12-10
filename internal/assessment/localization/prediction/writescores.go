package prediction

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/spf13/afero"
)

func writeScores(scores preyScore, outFile string) {
	var buffer bytes.Buffer

	writeHeader(&buffer)
	writeBody(&buffer, scores)

	afero.WriteFile(fs.Instance, outFile, buffer.Bytes(), 0644)
}

func writeHeader(buffer *bytes.Buffer) {
	buffer.WriteString("prey\tbait component\tdomain component\ttotal score\tbaits\tsupporting domains\tconflicting domains\n")
}

func writeBody(buffer *bytes.Buffer, scores preyScore) {
	outputOrder := orderPreys(scores.Bait)

	for _, prey := range outputOrder {
		baitString := strings.Join((*scores.Bait)[prey].baits, ";")
		conflictingDomainString := strings.Join((*scores.Domain)[prey].conflictingDomains, ";")
		supportingDomainString := strings.Join((*scores.Domain)[prey].supportingDomains, ";")
		buffer.WriteString(
			fmt.Sprintf(
				"%s\t%0.5f\t%0.5f\t%0.5f\t%s\t%s\t%s\n",
				prey,
				(*scores.Bait)[prey].score,
				(*scores.Domain)[prey].score,
				(*scores.Bait)[prey].score+(*scores.Domain)[prey].score,
				baitString,
				supportingDomainString,
				conflictingDomainString,
			),
		)
	}
}

func orderPreys(preyData *preyBaitScore) []string {
	keys := make([]string, len(*preyData))

	i := 0
	for key := range *preyData {
		keys[i] = key
		i++
	}

	sort.Strings(keys)
	return keys
}
