package mapfunc_test

import (
	"sort"

	"github.com/knightjdr/cmgo/pkg/mapfunc"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Values", func() {
	Describe("Map of type map[string]string", func() {
		It("should return string values", func() {
			m := map[string]string{
				"a": "value1",
				"d": "value2",
				"c": "value3",
			}

			actual := mapfunc.ValuesStringString(m)
			sort.Strings(actual)
			expected := []string{"value1", "value2", "value3"}
			Expect(actual).To(Equal(expected), "should return string values")
		})
	})
})
