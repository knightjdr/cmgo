package transmembrane

func countOrganelleBaitsPerPrey(baitsPerPrey map[string]map[string]bool, cytosolicBaits, lumenalBaits []string) map[string]map[string]int {
	organelleBaitsPerPrey := make(map[string]map[string]int, 0)

	for prey, baits := range baitsPerPrey {
		organelleBaitsPerPrey[prey] = map[string]int{
			"cytosolic": countBaits(baits, cytosolicBaits),
			"lumenal":   countBaits(baits, lumenalBaits),
		}
	}

	return organelleBaitsPerPrey
}

func countBaits(baitsPerPrey map[string]bool, organelleBaits []string) int {
	count := 0
	for _, bait := range organelleBaits {
		if _, ok := baitsPerPrey[bait]; ok {
			count++
		}
	}
	return count
}
