package gene

func dict(ids []map[string]string, fromType, toType string) map[string]string {
	mapping := make(map[string]string, 0)
	for _, id := range ids {
		if id[fromType] != "" {
			mapping[id[fromType]] = id[toType]
		}
	}
	return mapping
}
