package gene

// ID contains gene information.
type ID struct {
	EnsemblGene string `json:"ensembl_gene_id"`
	Entrez      string `json:"entrez_id"`
	HGNC        string `json:"hgnc_id"`
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
	UniProt     string
	UniProtList []string `json:"uniprot_ids"`
}

// Response contains individual gene entries.
type Response struct {
	Docs []ID `json:"docs"`
}

// Data for data received from HGNC.
type Data struct {
	Response Response `json:"response"`
}

// DefineUniProt adds the primary UniProt ID.
func (r *Response) DefineUniProt() {
	for i, doc := range r.Docs {
		if doc.UniProtList != nil && len(doc.UniProtList) > 0 {
			r.Docs[i].UniProt = doc.UniProtList[0]
		}
	}
}

// ParseIDtoMap converts the Gene struct to a map[string]string.
func (r *Response) ParseIDtoMap() []map[string]string {
	idMap := make([]map[string]string, len(r.Docs))
	for i, doc := range r.Docs {
		idMap[i] = map[string]string{
			"EnsemblGene": doc.EnsemblGene,
			"Entrez":      doc.Entrez,
			"HGNC":        doc.HGNC,
			"Name":        doc.Name,
			"Symbol":      doc.Symbol,
			"UniProt":     doc.UniProt,
		}
	}
	return idMap
}
