package localize

func addUniprotIDs(mapping *map[string]map[string]string, entrezToUniprot map[string]string) {
	for _, otherIDs := range *mapping {
		if _, ok := entrezToUniprot[otherIDs["Entrez"]]; ok {
			otherIDs["UniProt"] = entrezToUniprot[otherIDs["Entrez"]]
		}
	}
}
