package prediction

func calculateBaitComponent(options parameters, inputFiles fileContent) *preyBaitScore {
	baitCompartments := getBaitLocalizationsAsCompartments(options.baitExpected, inputFiles)

	return calculateBaitScoreComponent(inputFiles.predictions, baitCompartments, inputFiles.baitInteractors)
}
