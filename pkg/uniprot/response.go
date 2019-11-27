package uniprot

import (
	"github.com/tidwall/gjson"
)

// Entry contains UniProt information for a specific ID.
type Entry struct {
	Accession          string    `json:"accession"`
	Features           []Feature `json:"features"`
	Gene               Gene
	ID                 string   `json:"id"`
	Organism           Organism `json:"organism"`
	SecondaryAccession []string `json:"secondaryAccession"`
	Sequence           Sequence `json:"sequence"`
}

// Feature is a sequence topology feature
type Feature struct {
	Begin       int    `json:"begin,string"`
	Category    string `json:"category"`
	Description string `json:"description"`
	End         int    `json:"end,string"`
	Type        string `json:"type"`
}

// Gene name information.
type Gene struct {
	Symbol   string
	Synonyms []string
}

// Organism information for an entry.
type Organism struct {
	Common     string
	Names      []OrganismName `json:"names"`
	Scientific string
	Taxonomy   int `json:"taxonomy"`
}

// OrganismName can include common and scientific names.
type OrganismName struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// Sequence contains amino acid sequence information.
type Sequence struct {
	Length int `json:"length"`
	Mass   int `json:"mass"`
}

// ParseAdditionalFields parses other
func (e *Entry) ParseAdditionalFields(body []byte) {
	json := string(body)
	(*e).Gene.Symbol = gjson.Get(json, "gene.0.name.value").String()
	(*e).Gene.Synonyms = parseSynonyms(json)
	parseOrganism(e)
}

func parseSynonyms(jsonString string) []string {
	synonyms := make([]string, 0)

	result := gjson.Get(jsonString, "gene.0.synonyms")
	result.ForEach(func(index, obj gjson.Result) bool {
		obj.ForEach(func(key, value gjson.Result) bool {
			synonyms = append(synonyms, value.String())
			return true
		})
		return true
	})

	return synonyms
}

func parseOrganism(e *Entry) {
	for _, name := range (*e).Organism.Names {
		if name.Type == "common" {
			(*e).Organism.Common = name.Value
		} else if name.Type == "scientific" {
			(*e).Organism.Scientific = name.Value
		}
	}
}
