package subset

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefineColumns(t *testing.T) {
	columns := []string{"a", "b", "c", "d"}

	// TEST1: ranks are found in column names
	ranks1 := []string{"a", "d"}
	ranks2 := []string{"c"}
	wanted1Indices := []int{0, 3}
	wanted2Indices := []int{2}
	rank1Indices, rank2Indices, err := defineColumns(columns, ranks1, ranks2)
	assert.Nil(t, err, "Should not return an error when rank names can be found in columns")
	assert.Equal(t, wanted1Indices, rank1Indices, "Should return column indices for first compartment ranks")
	assert.Equal(t, wanted2Indices, rank2Indices, "Should return column indices for second compartment ranks")

	// TEST1: ranks are found in column names
	ranks1 = []string{"a", "e"}
	ranks2 = []string{"f"}
	wantErr := errors.New("ranks cannot be matched to columns")
	_, _, err = defineColumns(columns, ranks1, ranks2)
	assert.NotNil(t, err, "Should return error when missing arguments")
	assert.Equal(t, wantErr, err, "Should return correct error message")
}
