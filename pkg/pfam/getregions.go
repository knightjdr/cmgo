// Package pfam will get regions for a list of UniProt IDs.
package pfam

// GetRegions returns regions
func GetRegions(ids []string, url string) *Features {
	service := pfamService{URL: url}
	fetchRegions(&service, ids)
	return service.Result
}
