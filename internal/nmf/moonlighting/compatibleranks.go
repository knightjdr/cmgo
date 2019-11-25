package moonlighting

func defineCompatibleRanks(dissimilarityMatrix [][]float64) []map[int]bool {
	matrixDimension := len(dissimilarityMatrix)
	compatibleRanks := make([]map[int]bool, matrixDimension)
	allocateMemory(&compatibleRanks)

	for rowIndex, row := range dissimilarityMatrix {
		for columnIndex := rowIndex + 1; columnIndex < matrixDimension; columnIndex++ {
			if row[columnIndex] > 0 {
				allocateAndAssignMoonlightingPair(&compatibleRanks, columnIndex, rowIndex)
			}
		}
	}

	return compatibleRanks
}

func allocateMemory(compatibleRanks *[]map[int]bool) {
	for rowIndex := range *compatibleRanks {
		(*compatibleRanks)[rowIndex] = make(map[int]bool, 0)
	}
}

func allocateAndAssignMoonlightingPair(compatibleRanks *[]map[int]bool, columnIndex, rowIndex int) {
	(*compatibleRanks)[columnIndex][rowIndex] = true
	(*compatibleRanks)[rowIndex][columnIndex] = true
}
