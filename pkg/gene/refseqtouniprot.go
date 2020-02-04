package gene

// RefseqToUniProt maps refseq IDs to an array of UniProt accessions.
func RefseqToUniProt(ids []string, url string) map[string][]string {
	service := uniprotService{
		Query: ids,
		URL:   url,
	}
	fetchUniProt(&service)

	return service.Result
}
