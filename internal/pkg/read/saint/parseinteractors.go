package saint

// ParseInteractors returns significant interactors for each bait.
func (s *SAINT) ParseInteractors(fdr float64) map[string][]string {
	interactorsPerBait := make(map[string][]string, 0)

	for _, row := range *s {
		if row.FDR <= fdr {
			allocateMemoryInteractorsPerBait(&interactorsPerBait, row.Bait)
			interactorsPerBait[row.Bait] = append(interactorsPerBait[row.Bait], row.PreyGene)
		}
	}

	return interactorsPerBait
}

func allocateMemoryInteractorsPerBait(interactorsPerBait *map[string][]string, bait string) {
	if _, ok := (*interactorsPerBait)[bait]; !ok {
		(*interactorsPerBait)[bait] = make([]string, 0)
	}
}
