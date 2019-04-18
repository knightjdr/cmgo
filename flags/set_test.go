package flags

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetFloat(t *testing.T) {
	// TEST1: set from command line argument
	args := map[string]interface{}{
		"arg": "1.5",
	}
	fileOptions := map[string]interface{}{
		"arg": "1.0",
	}
	assert.Equal(t, 1.5, SetFloat("arg", args, fileOptions, 0.5), "Should return command line argument")

	// TEST2: set from file argument
	args = map[string]interface{}{}
	fileOptions = map[string]interface{}{
		"arg": "1.0",
	}
	assert.Equal(t, 1.0, SetFloat("arg", args, fileOptions, 0.5), "Should return file parameter")

	// TEST3: set from default value
	args = map[string]interface{}{}
	fileOptions = map[string]interface{}{}
	assert.Equal(t, 0.5, SetFloat("arg", args, fileOptions, 0.5), "Should return default value")
}

func TestSetInt(t *testing.T) {
	// TEST1: set from command line argument
	args := map[string]interface{}{
		"arg": "3",
	}
	fileOptions := map[string]interface{}{
		"arg": "2",
	}
	assert.Equal(t, 3, SetInt("arg", args, fileOptions, 1), "Should return command line argument")

	// TEST2: set from file argument
	args = map[string]interface{}{}
	fileOptions = map[string]interface{}{
		"arg": "2",
	}
	assert.Equal(t, 2, SetInt("arg", args, fileOptions, 1), "Should return file parameter")

	// TEST3: set from default value
	args = map[string]interface{}{}
	fileOptions = map[string]interface{}{}
	assert.Equal(t, 1, SetInt("arg", args, fileOptions, 1), "Should return default value")
}

func TestSetString(t *testing.T) {
	// TEST1: set from command line argument
	args := map[string]interface{}{
		"arg": "a",
	}
	fileOptions := map[string]interface{}{
		"arg": "b",
	}
	assert.Equal(t, "a", SetString("arg", args, fileOptions, "c"), "Should return command line argument")

	// TEST2: set from file argument
	args = map[string]interface{}{}
	fileOptions = map[string]interface{}{
		"arg": "b",
	}
	assert.Equal(t, "b", SetString("arg", args, fileOptions, "c"), "Should return file parameter")

	// TEST3: set from default value
	args = map[string]interface{}{}
	fileOptions = map[string]interface{}{}
	assert.Equal(t, "c", SetString("arg", args, fileOptions, "c"), "Should return default value")
}
