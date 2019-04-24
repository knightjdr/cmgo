package strfunc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBeforeLast(t *testing.T) {
	// TEST1: string contains a delimiter
	str := "file.txt"
	wanted := "file"
	assert.Equal(t, wanted, BeforeLast(str, "."), "Should return string before delimiter")

	// TEST2: string contains multiple delimiters
	str = "file.txt.svg"
	wanted = "file.txt"
	assert.Equal(t, wanted, BeforeLast(str, "."), "Should return string before last instance of delimiter")

	// TEST3: string does not contain delimiter
	str = "file"
	wanted = "file"
	assert.Equal(t, wanted, BeforeLast(str, "."), "Should return entire string")

	// TEST4: last instance of delimiter is first character
	str = ".file"
	wanted = ""
	assert.Equal(t, wanted, BeforeLast(str, "."), "Should return nil string")
}
