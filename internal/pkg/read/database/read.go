// Package database reads a fasta database to an array.
package database

import (
	"bufio"
	"log"
	"regexp"
	"strings"

	"github.com/knightjdr/cmgo/pkg/fs"
)

// Fasta details for an entry in a database fasta file
type Fasta struct {
	Entrez   string
	Refseq   string
	Sequence string
	Symbol   string
}

func appendEntry(entries []Fasta, entry Fasta, sequence strings.Builder, includeSequence bool) []Fasta {
	if entry.Refseq != "" {
		if includeSequence {
			entry.Sequence = sequence.String()
		}
		return append(entries, entry)
	}
	return entries
}

// Read reads a fasta database.
func Read(filename string, includeSequence bool) []Fasta {
	file, err := fs.Instance.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	re := regexp.MustCompile(`^>([^|]+)\|gn\|([^:]+):(\d+)\|`)
	entry := Fasta{}
	entries := make([]Fasta, 0)
	var sequence strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, ">") {
			matches := re.FindStringSubmatch(line)
			if len(matches) > 0 {
				entries = appendEntry(entries, entry, sequence, includeSequence)
				entry.Entrez = matches[3]
				entry.Refseq = matches[1]
				entry.Symbol = matches[2]
			}
			entry.Sequence = ""
			sequence.Reset()
		} else {
			sequence.WriteString(line)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	entries = appendEntry(entries, entry, sequence, includeSequence)

	return entries
}
