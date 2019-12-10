package uniprot

// CreateGeneNameMap creates a map of official gene names to UniProt ID.
func (u *Entries) CreateGeneNameMap() map[string]string {
	geneToUniProt := make(map[string]string, 0)

	for id, entry := range *u {
		if entry.Symbol != "" {
			geneToUniProt[entry.Symbol] = id
		}
	}

	return geneToUniProt
}
