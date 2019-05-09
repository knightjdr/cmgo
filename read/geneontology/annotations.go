package geneontology

import (
	"encoding/csv"
	"io"
	"log"
	"strings"

	"github.com/knightjdr/cmgo/fs"
)

// GOannotations contains GO annotations for each gene and a mapping for gene name
// to UniProt ID. The "Genes" value will be a map of GO namespace/gene name/GO ID,
// with information about each annotatedID
type GOannotations struct {
	Genes          *map[string]map[string]map[string]*GOannotation
	UniProtMapping map[string]string
}

func newGOannotations() GOannotations {
	h := map[string]map[string]map[string]*GOannotation{
		"BP": make(map[string]map[string]*GOannotation, 0),
		"CC": make(map[string]map[string]*GOannotation, 0),
		"MF": make(map[string]map[string]*GOannotation, 0),
	}
	return GOannotations{
		Genes:          &h,
		UniProtMapping: make(map[string]string, 0),
	}
}

// GOannotation contains information about each GO annotation
type GOannotation struct {
	Sources map[string]bool
}

func mapAnnotationLine(line []string) map[string]string {
	return map[string]string{
		"gene":      line[2],
		"id":        line[4],
		"namespace": shortNamespaceMap[line[8]],
		"source":    line[14],
		"uniprot":   line[1],
	}
}

var shortNamespaceMap = map[string]string{
	"C": "CC",
	"F": "MF",
	"P": "BP",
}

// Annotations reads a go-annotation.gaf file.
func Annotations(filename string) GOannotations {
	file, err := fs.Instance.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	reader := csv.NewReader(file)
	reader.Comma = '\t'
	reader.FieldsPerRecord = -1
	reader.LazyQuotes = true

	// Read annoations.
	annotations := newGOannotations()
	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}

		if !strings.HasPrefix(line[0], "!") {
			annotation := mapAnnotationLine(line)
			if _, ok := (*annotations.Genes)[annotation["namespace"]][annotation["gene"]]; !ok {
				(*annotations.Genes)[annotation["namespace"]][annotation["gene"]] = make(map[string]*GOannotation, 0)
			}
			if _, ok := (*annotations.Genes)[annotation["namespace"]][annotation["gene"]][annotation["id"]]; !ok {
				(*annotations.Genes)[annotation["namespace"]][annotation["gene"]][annotation["id"]] = &GOannotation{
					Sources: make(map[string]bool, 0),
				}
			}
			(*annotations.Genes)[annotation["namespace"]][annotation["gene"]][annotation["id"]].Sources[annotation["source"]] = true
			annotations.UniProtMapping[annotation["gene"]] = annotation["uniprot"]
		}
	}

	return annotations
}
