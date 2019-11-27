// Package uniprot will get uniprot entries for a list of IDs.
package uniprot

// GetProteins returns regions
func GetProteins(ids []string, url string) *Entries {
	service := uniprotService{URL: url}
	fetchUniprot(&service, ids)
	return service.Result
}
