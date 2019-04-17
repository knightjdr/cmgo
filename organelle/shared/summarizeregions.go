package shared

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	"github.com/knightjdr/cmgo/customsort"
	"github.com/knightjdr/cmgo/fs"
	"github.com/knightjdr/cmgo/slice"
	"github.com/spf13/afero"
)

func summarizeRegions(proteins []string, regions map[string]map[string]bool, outfile string) {
	regionCount := make(map[string]int, 0)
	regionProteins := make(map[string][]string, 0)
	for _, protein := range proteins {
		if _, ok := regions[protein]; ok {
			for region := range regions[protein] {
				regionCount[region]++
				regionProteins[region] = append(regionProteins[region], protein)
			}
		}
	}
	order := customsort.ByMapValueInt(regionCount, "descending")

	var buffer bytes.Buffer
	buffer.WriteString("region\tno. preys\tpreys\tpreys not containing region\n")
	sort.Strings(proteins)
	buffer.WriteString(fmt.Sprintf("-\t%d\t%s\t\n", len(proteins), strings.Join(proteins, ", ")))
	for _, region := range order {
		currProteins := regionProteins[region.Key]
		sort.Strings(currProteins)

		diff := slice.Diff(proteins, currProteins)
		sort.Strings(diff)
		buffer.WriteString(fmt.Sprintf("%s\t%d\t%s\t%s\n", region.Key, region.Value, strings.Join(currProteins, ", "), strings.Join(diff, ", ")))
	}
	afero.WriteFile(fs.Instance, outfile, buffer.Bytes(), 0644)
}
