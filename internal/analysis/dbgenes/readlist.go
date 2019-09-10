package dbgenes

import (
	"bufio"
	"log"
	"regexp"
	"strings"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/knightjdr/cmgo/pkg/slice"
)

func readList(filename string) []string {
	file, err := fs.Instance.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	reGene := regexp.MustCompile(`^GN   Name=([^;]+);`)
	reOrgansim := regexp.MustCompile(`^OX   NCBI_TaxID=(\d+);`)

	genes := make([]string, 0)
	scanner := bufio.NewScanner(file)
	var gene string
	var species string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "GN") {
			matches := reGene.FindStringSubmatch(line)
			if len(matches) > 0 {
				gene = strings.Split(matches[1], " ")[0]
			}
		}
		if strings.HasPrefix(line, "OX") {
			matches := reOrgansim.FindStringSubmatch(line)
			if len(matches) > 0 {
				species = matches[1]
			}
		}
		if strings.HasPrefix(line, "//") {
			if species == "9606" && gene != "" {
				genes = append(genes, gene)
			}
			gene = ""
			species = ""
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	genes = slice.UniqueStrings(genes)

	return genes
}
