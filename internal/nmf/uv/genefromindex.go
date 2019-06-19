package uv

func geneFromIndex(nonCharacterizingRowIndicies [][]int, rowNames []string) [][]string {
	nonCharacterizingGenes := make([][]string, len(nonCharacterizingRowIndicies))
	for i, rowIndices := range nonCharacterizingRowIndicies {
		nonCharacterizingGenes[i] = make([]string, len(rowIndices))
		for j, index := range rowIndices {
			nonCharacterizingGenes[i][j] = rowNames[index]
		}
	}
	return nonCharacterizingGenes
}
