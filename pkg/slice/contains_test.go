package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceContains(t *testing.T) {
	testIntSlice := []int{1, 2, 3}
	testStringSlice := []string{"a", "c", "d"}

	// TEST: slice contains tested string value
	shouldContainString := []string{"a", "c", "d"}
	for _, value := range shouldContainString {
		assert.True(t, ContainsString(value, testStringSlice), "should find value in []string")
	}

	// TEST: slice does not contain tested string value
	shouldNotContainString := []string{"aa", "b", "something"}
	for _, value := range shouldNotContainString {
		assert.False(t, ContainsString(value, testStringSlice), "should not find value in []string")
	}

	// TEST: slice contains tested int value
	shouldContainInt := []int{1, 2, 3}
	for _, value := range shouldContainInt {
		assert.True(t, ContainsInt(value, testIntSlice), "should find value in []int")
	}

	// TEST: slice does not contain tested int value
	shouldNotContainInt := []int{11, 22}
	for _, value := range shouldNotContainInt {
		assert.False(t, ContainsInt(value, testIntSlice), "should not find value in []int")
	}
}
