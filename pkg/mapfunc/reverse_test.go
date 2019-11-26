package mapfunc_test

import (
	"github.com/knightjdr/cmgo/pkg/mapfunc"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Reverse", func() {
	Describe("Map of type map[string]string", func() {
		It("should reverse keys and values", func() {
			m := map[string]string{
				"a": "value1",
				"d": "value2",
				"c": "value3",
			}

			expected := map[string]string{
				"value1": "a",
				"value2": "d",
				"value3": "c",
			}
			Expect(mapfunc.ReverseStringString(m)).To(Equal(expected), "should return string values")
		})
	})
})
