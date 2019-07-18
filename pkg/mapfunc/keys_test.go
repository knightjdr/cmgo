package mapfunc_test

import (
	"sort"

	"github.com/knightjdr/cmgo/pkg/mapfunc"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Keys", func() {
	Describe("Map of type map[int]float64", func() {
		It("should return integer keys", func() {
			m := map[int]float64{
				3: 0.1,
				1: 0.1,
				5: 0.2,
			}

			actual := mapfunc.KeysIntFloat(m)
			sort.Ints(actual)
			expected := []int{1, 3, 5}
			Expect(actual).To(Equal(expected), "should return integer keys")
		})
	})

	Describe("Map of type map[string]bool", func() {
		It("should return integer keys", func() {
			m := map[string]bool{
				"a": true,
				"d": true,
				"c": true,
			}

			actual := mapfunc.KeysStringBool(m)
			sort.Strings(actual)
			expected := []string{"a", "c", "d"}
			Expect(actual).To(Equal(expected), "should return string keys")
		})
	})

	Describe("Map of type map[string]string", func() {
		It("should return integer string", func() {
			m := map[string]string{
				"a": "test",
				"d": "test",
				"c": "test",
			}

			actual := mapfunc.KeysStringString(m)
			sort.Strings(actual)
			expected := []string{"a", "c", "d"}
			Expect(actual).To(Equal(expected), "should return string keys")
		})
	})
})
