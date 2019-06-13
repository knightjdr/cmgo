package rbo

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Overlapsize", func() {
	It("should return length of input slice", func() {
		intersections := [][]string{
			{"a", "b", "c"},
			{"a", "b", "c", "d"},
			{"a"},
		}
		expected := []int{3, 4, 1}
		for i, intersection := range intersections {
			Expect(X(intersection)).To(Equal(expected[i]), "should return length of intersection")
		}
	})
})
