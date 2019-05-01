package dbgenes

import (
	"bufio"
	"log"
	"regexp"
	"strings"

	"github.com/knightjdr/cmgo/fs"
	"github.com/knightjdr/cmgo/slice"
)

func readDatabase(filename string) []string {
	file, err := fs.Instance.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	genes := make([]string, 0)
	re := regexp.MustCompile(`\|([\w-]+):\d+\|`)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, ">") {
			matches := re.FindStringSubmatch(line)
			if len(matches) > 0 {
				genes = append(genes, matches[1])
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	genes = slice.UniqueStrings(genes)

	return genes
}
