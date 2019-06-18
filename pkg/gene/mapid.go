// Package gene implements methods for mapping gene IDs
package gene

// MapIDs will map a slice of IDs of the type "fromType" to the type "toType"
// and return a map of those conversions.
func MapIDs(ids []string, fromType, toType, url string) map[string]string {
	service := hgncService{URL: url}
	fetchHGNC(&service)
	mappingAll := dict(service.Result, fromType, toType)

	mappingRequested := make(map[string]string, 0)
	for _, id := range ids {
		var target string
		if _, ok := mappingAll[id]; ok {
			target = mappingAll[id]
		}
		mappingRequested[id] = target
	}
	return mappingRequested
}
