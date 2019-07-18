package localize

func filterEnrichment(enrichment map[string][]Enrichment, goIDs map[string]string) map[string][]Enrichment {
	validEnrichment := make(map[string][]Enrichment, 0)
	// Remove invalid localizations.
	for gene, terms := range enrichment {
		validEnrichment[gene] = make([]Enrichment, 0)
		for _, term := range terms {
			if _, ok := goIDs[term.ID]; ok {
				validEnrichment[gene] = append(validEnrichment[gene], term)
			}
		}
	}

	return validEnrichment
}
