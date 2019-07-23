// Package matrix contains methods for operating on matrices.
package matrix

// Transpose transposes a 2D matrix.
func Transpose(matrix [][]float64) [][]float64 {
	// Matrix dimensions.
	colNum := len(matrix[0])
	rowNum := len(matrix)

	transposed := make([][]float64, colNum)
	for i := range transposed {
		transposed[i] = make([]float64, rowNum)
	}
	for i, row := range matrix {
		for j, value := range row {
			transposed[j][i] = value
		}
	}
	return transposed
}
