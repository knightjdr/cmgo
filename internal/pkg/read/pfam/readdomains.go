// Package pfam reads domain and motif information from a file.
package pfam

import (
	"strconv"

	"github.com/knightjdr/cmgo/internal/pkg/read/csv"
)

// Domain name and start and end amino acid.
type Domain struct {
	End   int
	Name  string
	Start int
}

// Domains for each UniProt entry.
type Domains map[string][]Domain

// ReadDomains reads and returns a list of domains for each UniProt ID.
func ReadDomains(filename string) Domains {
	reader := csv.Read(filename, false)

	domains := make(map[string][]Domain, 0)
	for {
		eof, line := csv.Readline(reader)
		if eof {
			break
		}

		id, domain := mapDomainLine(line)
		allocateDomainMemory(domains, id)
		domains[id] = append(domains[id], domain)
	}

	return domains
}

func mapDomainLine(line []string) (string, Domain) {
	end, _ := strconv.Atoi(line[2])
	id := line[0]
	start, _ := strconv.Atoi(line[1])
	domain := Domain{
		End:   end,
		Name:  line[6],
		Start: start,
	}

	return id, domain
}

func allocateDomainMemory(domains Domains, uniprotID string) {
	if _, ok := domains[uniprotID]; !ok {
		domains[uniprotID] = make([]Domain, 0)
	}
}
