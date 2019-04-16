package flags

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	// TEST1
	os.Args = []string{
		"cmd",
		"-optiona=a",
		"--optionb", "1",
	}
	wanted := map[string]interface{}{
		"optiona": "a",
		"optionb": "1",
	}
	options := Parse()
	assert.Equal(t, wanted, options, "Should return command line arguments as interface")
}
