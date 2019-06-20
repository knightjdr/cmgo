package hydropathy

func diff(saintRefseq []string, bioplexEntrez []string, refseqEntrez map[string]string) []string {
	bioplexDict := make(map[string]bool, len(bioplexEntrez))
	for _, id := range bioplexEntrez {
		bioplexDict[id] = true
	}

	uniqueToSaint := make([]string, 0)
	for _, refseqID := range saintRefseq {
		entrez := ""
		if _, ok := refseqEntrez[refseqID]; ok {
			entrez = refseqEntrez[refseqID]
		}

		if _, ok := bioplexDict[entrez]; !ok {
			uniqueToSaint = append(uniqueToSaint, refseqID)
		}
	}
	return uniqueToSaint
}
