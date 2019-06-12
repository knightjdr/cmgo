package gprofiler

import (
	"sort"
)

// EnrichedTerm contains information on an enriched term.
type EnrichedTerm struct {
	Genes            []string
	ID               string `json:"native"`
	Intersections    [][]string
	IntersectionSize int `json:"intersection_size"`
	Name             string
	QuerySize        int `json:"query_size"`
	Precision        float64
	Pvalue           float64 `json:"p_value"`
	Recall           float64
	Source           string
	TermSize         int `json:"term_size"`
}

// GenesMetaData contains g:Profiler information on queried genes.
type GenesMetaData struct {
	Query map[string]Query
}

// MetaData contains g:Profiler query meta data.
type MetaData struct {
	GenesMetaData GenesMetaData `json:"genes_metadata"`
}

// Query contains ENS to gene name mapping data for queried genes.
type Query struct {
	ENSG    []string `json:"ensgs"`
	Mapping map[string][]string
}

// Response for data received from gProfiler.
type Response struct {
	MetaData MetaData `json:"meta"`
	Result   []EnrichedTerm
}

// AddIntersectionGenes adds the list of genes with the enriched term to each enriched
// term struct.
func (r *Response) AddIntersectionGenes(queryName string) {
	ensemblIDs := r.MetaData.GenesMetaData.Query[queryName].ENSG
	genesToEnsembl := r.MetaData.GenesMetaData.Query[queryName].Mapping

	ensemblToGenes := make(map[string]string, len(genesToEnsembl))
	for gene, ids := range genesToEnsembl {
		ensemblToGenes[ids[0]] = gene
	}

	for i, enrichedTerm := range r.Result {
		for j, evidence := range enrichedTerm.Intersections {
			if len(evidence) > 0 {
				gene := ensemblToGenes[ensemblIDs[j]]
				r.Result[i].Genes = append(r.Result[i].Genes, gene)
			}
		}
		sort.Strings(r.Result[i].Genes)
	}
}
