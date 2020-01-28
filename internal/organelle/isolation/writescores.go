package isolation

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	"github.com/knightjdr/cmgo/internal/pkg/read/localization"
	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/spf13/afero"
)

func writeScores(scores *isolationScores, nmfSummary localization.Summary, outfile string) {
	outputOrder := getOutputOrder(scores)

	var buffer bytes.Buffer
	buffer.WriteString("compartment\tname\tisolation score\tedges within\tedges outside\n")
	for _, compartment := range outputOrder {
		score := (*scores)[compartment]
		compartmentName := strings.Join(nmfSummary[compartment].DisplayTerms, ", ")
		buffer.WriteString(fmt.Sprintf("%d\t%s\t%0.3f\t%d\t%d\n", compartment, compartmentName, score.isolation, score.edgesWithin, score.edgesOutside))
	}
	afero.WriteFile(fs.Instance, outfile, buffer.Bytes(), 0644)
}

func getOutputOrder(scores *isolationScores) []int {
	keys := make([]int, len(*scores))

	i := 0
	for compartment := range *scores {
		keys[i] = compartment
		i++
	}
	sort.Ints(keys)

	return keys
}
