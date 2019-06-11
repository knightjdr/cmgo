package subset

import (
	"errors"
)

func matchRank(columns, ranks []string) ([]int, error) {
	columnIndices := make([]int, len(ranks))

	var err error
	for i, rank := range ranks {
		match := false
		for j, column := range columns {
			if rank == column {
				columnIndices[i] = j
				match = true
				break
			}
		}
		if !match {
			err = errors.New("ranks cannot be matched to columns")
			break
		}
	}
	return columnIndices, err
}

func defineColumns(columns, ranks1, ranks2 []string) ([]int, []int, error) {
	rank1Indices, err1 := matchRank(columns, ranks1)
	rank2Indices, err2 := matchRank(columns, ranks2)

	var err error
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}

	return rank1Indices, rank2Indices, err
}
