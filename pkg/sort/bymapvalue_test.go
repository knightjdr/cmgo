package sort_test

import (
	"github.com/knightjdr/cmgo/pkg/sort"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ByMapValue", func() {
	Describe("Sort map[string]int", func() {
		Context("in descending order", func() {
			It("should return sorted map", func() {
				m := map[string]int{
					"a": 5,
					"b": 20,
					"c": 7,
				}

				expected := []sort.KeyValueStringInt{
					{Key: "b", Value: 20},
					{Key: "c", Value: 7},
					{Key: "a", Value: 5},
				}
				Expect(sort.ByMapValueStringInt(m, "descending")).To(Equal(expected), "should sort in descending order")
			})
		})

		Context("in ascending order", func() {
			It("should return sorted map", func() {
				m := map[string]int{
					"a": 5,
					"b": 20,
					"c": 7,
				}

				expected := []sort.KeyValueStringInt{
					{Key: "a", Value: 5},
					{Key: "c", Value: 7},
					{Key: "b", Value: 20},
				}
				Expect(sort.ByMapValueStringInt(m, "ascending")).To(Equal(expected), "should sort in ascending order")
			})
		})
	})

	Describe("Sort map[int]float64", func() {
		Context("in descending order", func() {
			It("should return sorted map", func() {
				m := map[int]float64{
					2: 5,
					4: 20,
					7: 7,
				}

				expected := []sort.KeyValueIntFloat{
					{Key: 4, Value: 20},
					{Key: 7, Value: 7},
					{Key: 2, Value: 5},
				}
				Expect(sort.ByMapValueIntFloat64(m, "descending")).To(Equal(expected), "should sort in descending order")
			})
		})

		Context("in ascending order", func() {
			It("should return sorted map", func() {
				m := map[int]float64{
					2: 5,
					4: 20,
					7: 7,
				}

				expected := []sort.KeyValueIntFloat{
					{Key: 2, Value: 5},
					{Key: 7, Value: 7},
					{Key: 4, Value: 20},
				}
				Expect(sort.ByMapValueIntFloat64(m, "ascending")).To(Equal(expected), "should sort in ascending order")
			})
		})
	})
})
