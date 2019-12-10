// Package uniprot reads a uniprot database.
package uniprot

import (
	"bufio"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/knightjdr/cmgo/pkg/fs"
)

// Entries is a map of UniProt IDs with their information.
type Entries map[string]Entry

// Entry contains information about each UniProt ID.
type Entry struct {
	Symbol string
}

// Read all entries in a UniProt file (.dat format) specified by the organsim argument.
func Read(filename string, taxonID int) *Entries {
	file, err := fs.Instance.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	re := createRe()
	taxonString := fmt.Sprintf("%d", taxonID)

	entries := &Entries{}
	entry := initializeEntry()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "AC") {
			id := getValue(re["id"], line)
			if entry["id"] == "" {
				entry["id"] = id
			}
		} else if strings.HasPrefix(line, "GN") {
			entry["symbol"] = getValue(re["symbol"], line)
		} else if strings.HasPrefix(line, "OX") {
			entry["species"] = getValue(re["organism"], line)
		} else if strings.HasPrefix(line, "//") {
			if entry["species"] == taxonString && entry["id"] != "" {
				(*entries)[entry["id"]] = createEntry(entry)
			}
			entry = initializeEntry()
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	return entries
}

func createRe() map[string]*regexp.Regexp {
	id := regexp.MustCompile(`^AC   (\w+);`)
	organism := regexp.MustCompile(`^OX   NCBI_TaxID=(\d+);`)
	symbol := regexp.MustCompile(`^GN   Name=([^;]+);`)
	return map[string]*regexp.Regexp{
		"id":       id,
		"organism": organism,
		"symbol":   symbol,
	}
}

func initializeEntry() map[string]string {
	return map[string]string{
		"id":      "",
		"species": "",
		"symbol":  "",
	}
}

func getValue(re *regexp.Regexp, line string) string {
	matches := re.FindStringSubmatch(line)
	if len(matches) > 0 {
		return strings.Split(matches[1], " ")[0]
	}
	return ""
}

func createEntry(entry map[string]string) Entry {
	return Entry{
		Symbol: entry["symbol"],
	}
}
