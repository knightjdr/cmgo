package correlation

import (
	"github.com/knightjdr/cmgo/pkg/matrix"
)

// CoefficientMatrix calculates the correlation coefficient between the rows or columns
// of a matrix. Use the "row" argument to specify calculating correlation between
// rows (true) or columns (false). Currently, only the Pearson method is implemented.
func CoefficientMatrix(inputMatrix [][]float64, row bool, method string) [][]float64 {
	x := inputMatrix
	if !row {
		x = matrix.Transpose(inputMatrix)
	}

	var n int
	if row {
		n = len(x)
	} else {
		n = len(x[0])
	}

	cc := make([][]float64, n)
	for i := 0; i < n; i++ {
		cc[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			cc[i][j] = Pearson(x[i], x[j])
		}
	}

	return cc
}
