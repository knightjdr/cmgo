package saint

// FilterByBaitNumber filters a SAINT file by a minimum bait number
// required per prey.
func (saint *SAINT) FilterByBaitNumber(minimumBaits int) {
	if minimumBaits <= 1 {
		return
	}

	preysOccurences := countPreyOccurences(saint)
	removePreysNotPassingCutoff(saint, preysOccurences, minimumBaits)
}

func countPreyOccurences(saint *SAINT) map[string]int {
	preys := make(map[string]int, 0)
	for _, row := range *saint {
		preys[row.PreyGene]++
	}
	return preys
}

func removePreysNotPassingCutoff(saint *SAINT, preysOccurences map[string]int, minimumBaits int) {
	for i := len(*saint) - 1; i >= 0; i-- {
		if preysOccurences[(*saint)[i].PreyGene] < minimumBaits {
			*saint = append((*saint)[:i], (*saint)[i+1:]...)
		}
	}
}
