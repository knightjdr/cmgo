package prediction

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/spf13/afero"
)

func writeScores(scores preyScore, inputFiles fileContent, outFile string) {
	var buffer bytes.Buffer

	writeHeader(&buffer)
	writeBody(&buffer, scores, inputFiles)

	afero.WriteFile(fs.Instance, outFile, buffer.Bytes(), 0644)
}

func writeHeader(buffer *bytes.Buffer) {
	buffer.WriteString("prey\tcompartment\tGO term(s)\tGO ID(s)\tbait component\tdomain component\tstudy component\ttotal score\tbaits\tsupporting domains\tconflicting domains\tHPA supporting\tFractionation supporting\n")
}

func writeBody(buffer *bytes.Buffer, scores preyScore, inputFiles fileContent) {
	outputOrder := orderPreys(scores.Bait)

	for _, prey := range outputOrder {
		baitString := strings.Join((*scores.Bait)[prey].baits, ";")
		compartment := inputFiles.predictions[prey]
		conflictingDomainString := strings.Join((*scores.Domain)[prey].conflictingDomains, ";")
		supportingDomainString := strings.Join((*scores.Domain)[prey].supportingDomains, ";")
		supportingFractionation := strings.Join((*scores.Study)[prey].fractionation, ";")
		supportingHPA := strings.Join((*scores.Study)[prey].hpa, ";")
		buffer.WriteString(
			fmt.Sprintf(
				"%s\t%d\t%s\t%s\t%0.5f\t%0.5f\t%0.5f\t%0.5f\t%s\t%s\t%s\t%s\t%s\n",
				prey,
				compartment,
				strings.Join(inputFiles.predictionSummary[compartment].GOterms, ";"),
				strings.Join(inputFiles.predictionSummary[compartment].GOid, ";"),
				(*scores.Bait)[prey].score,
				(*scores.Domain)[prey].score,
				(*scores.Study)[prey].score,
				((*scores.Bait)[prey].score+(*scores.Domain)[prey].score+(*scores.Study)[prey].score)/3,
				baitString,
				supportingDomainString,
				conflictingDomainString,
				supportingHPA,
				supportingFractionation,
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
