package rbo

import (
	"sort"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Intersection", func() {
	Context("Lists are longer than comparison limit", func() {
		It("should return intersecting elements", func() {
			d := 3
			S := []string{"a", "b", "c", "d", "e"}
			T := []string{"a", "b", "d", "c", "f"}
			actual := I(S, T, d)
			sort.Strings(actual)
			expected := []string{"a", "b"}
			Expect(actual).To(Equal(expected))
		})
	})

	Context("Lists are longer equal in length to comparison limit", func() {
		It("should return intersecting elements", func() {
			d := 5
			S := []string{"a", "b", "c", "d", "e"}
			T := []string{"a", "b", "d", "c", "f"}
			actual := I(S, T, d)
			sort.Strings(actual)
			expected := []string{"a", "b", "c", "d"}
			Expect(actual).To(Equal(expected))
		})
	})

	Context("Lists are shorter than comparison limit", func() {
		It("should return intersecting elements", func() {
			d := 10
			S := []string{"a", "b", "c", "d", "e"}
			T := []string{"a", "b", "d", "c", "f"}
			actual := I(S, T, d)
			sort.Strings(actual)
			expected := []string{"a", "b", "c", "d"}
			Expect(actual).To(Equal(expected))
		})
	})
})
