package geneontology

import (
	"bufio"
	"log"
	"regexp"
	"strings"

	"github.com/knightjdr/cmgo/pkg/fs"
)

// GOhierarchy contains a parsed GO hierarchy split into namespaces.
type GOhierarchy map[string]map[string]*GOterm

func newGOhierarchy() *GOhierarchy {
	h := make(GOhierarchy, 3)
	hp := &h
	(*hp)["BP"] = make(map[string]*GOterm, 0)
	(*hp)["CC"] = make(map[string]*GOterm, 0)
	(*hp)["MF"] = make(map[string]*GOterm, 0)
	return hp
}

// GOsynonym contains synonyms for a GO term.
// Exact: an exact equivalent; interchangeable with the term name
// Broad: the synonym is broader than the term name
// Narrow: the synonym is narrower or more precise than the term name
// Related: the terms are related in some imprecise way
type GOsynonym map[string][]string

func newGOsynonym() GOsynonym {
	s := make(GOsynonym, 4)
	s["Broad"] = make([]string, 0)
	s["Exact"] = make([]string, 0)
	s["Narrow"] = make([]string, 0)
	s["Related"] = make([]string, 0)
	return s
}

// GOterm contains details about a GO term.
type GOterm struct {
	Children      []string
	DirectParents []string
	Name          string
	Parents       []string
	Synonyms      GOsynonym
}

var longNamespaceMap = map[string]string{
	"biological_process": "BP",
	"cellular_component": "CC",
	"molecular_function": "MF",
}

// OBO reads a go hiearchy in .obo format.
func OBO(filename string) *GOhierarchy {
	file, err := fs.Instance.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	// Regex to resuse.
	reID := regexp.MustCompile(`GO:\d+`)
	reName := regexp.MustCompile(`name: (.+)`)
	reNamespace := regexp.MustCompile(`namespace: ([a-z_]+)`)
	reSynonym := regexp.MustCompile(`synonym: "([^"]+)" ([A-Z]+)`)

	hierarchy := newGOhierarchy()

	// Variables to record for each entry
	var id string
	var name string
	var namespace string
	var obsolete bool
	parents := make([]string, 0)
	synonyms := newGOsynonym()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "id:") {
			id = reID.FindString(line)
		} else if strings.HasPrefix(line, "name:") {
			matches := reName.FindStringSubmatch(line)
			name = matches[1]
		} else if strings.HasPrefix(line, "namespace:") {
			matches := reNamespace.FindStringSubmatch(line)
			namespace = longNamespaceMap[matches[1]]
		} else if strings.HasPrefix(line, "is_a:") {
			parent := reID.FindString(line)
			parents = append(parents, parent)
		} else if strings.HasPrefix(line, "is_obsolete: true") {
			obsolete = true
		} else if strings.HasPrefix(line, "relationship: part_of") {
			parent := reID.FindString(line)
			parents = append(parents, parent)
		} else if strings.HasPrefix(line, "synonym:") {
			matches := reSynonym.FindStringSubmatch(line)
			synonym := matches[1]
			synonymType := strings.Title(strings.ToLower(matches[2]))
			synonyms[synonymType] = append(synonyms[synonymType], synonym)
		} else if line == "" {
			if id != "" && !obsolete {
				(*hierarchy)[namespace][id] = &GOterm{
					DirectParents: parents,
					Name:          name,
					Synonyms:      synonyms,
				}
			}
			id = ""
			name = ""
			namespace = ""
			obsolete = false
			parents = make([]string, 0)
			synonyms = newGOsynonym()
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	return hierarchy
}
