package prediction

func calculateBaitComponent(options parameters, inputFiles fileContent) *preyBaitScore {
	compartmentsPerBait := getBaitLocalizationsAsCompartments(options.baitExpected, inputFiles)
	baitsPerPrey := getBaitsPerPrey(inputFiles.baitInteractors)

	return calculateBaitScoreComponent(inputFiles, compartmentsPerBait, baitsPerPrey)
}

func getBaitsPerPrey(baitInteractors map[string][]string) map[string][]string {
	baitsPerPrey := make(map[string][]string)

	for bait, preys := range baitInteractors {
		for _, prey := range preys {
			if _, ok := baitsPerPrey[prey]; !ok {
				baitsPerPrey[prey] = make([]string, 0)
			}
			baitsPerPrey[prey] = append(baitsPerPrey[prey], bait)
		}
	}

	return baitsPerPrey
}
