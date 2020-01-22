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
	header := "prey\tcompartment\tGO term(s)\tGO ID(s)\t" +
		"bait-organelle recovery component\tbait-organelle specificity component\tstudy component\ttext component\ttotal score\t" +
		"organelle specific baits\tHPA supporting\tFractionation supporting\tbest text term\n"
	buffer.WriteString(header)
}

func writeBody(buffer *bytes.Buffer, scores preyScore, inputFiles fileContent) {
	outputOrder := orderPreys(scores.Bait)

	for _, prey := range outputOrder {
		baitString := strings.Join((*scores.Bait)[prey].organelleBaits, ";")
		compartment := inputFiles.predictions[prey]
		supportingFractionation := strings.Join((*scores.Study)[prey].fractionation, ";")
		supportingHPA := strings.Join((*scores.Study)[prey].hpa, ";")
		buffer.WriteString(
			fmt.Sprintf(
				"%s\t%d\t%s\t%s\t%0.5f\t%0.5f\t%0.5f\t%0.5f\t%0.5f\t%s\t%s\t%s\t%s\n",
				prey,
				compartment,
				strings.Join(inputFiles.predictionSummary[compartment].GOterms, ";"),
				strings.Join(inputFiles.predictionSummary[compartment].GOid, ";"),
				(*scores.Bait)[prey].scoreOrganelle,
				(*scores.Bait)[prey].scoreSpecificity,
				(*scores.Study)[prey].score,
				(*scores.Text)[prey].score,
				((*scores.Bait)[prey].scoreOrganelle+(*scores.Bait)[prey].scoreSpecificity+(*scores.Study)[prey].score+(*scores.Text)[prey].score)/4,
				baitString,
				supportingHPA,
				supportingFractionation,
				(*scores.Text)[prey].GOID,
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
