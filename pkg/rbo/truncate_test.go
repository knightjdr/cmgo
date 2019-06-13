package rbo

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Truncate", func() {
	Context("User supplied k is below allowable limit", func() {
		It("should return truncate slices", func() {
			k := 3
			S := []string{"a", "b", "c", "d"}
			T := []string{"a", "b", "c", "d", "e"}
			actualS, actualT := truncate(S, T, k)
			Expect(actualS).To(Equal(S[:3]), "should return slice S truncated to len 3")
			Expect(actualT).To(Equal(T[:3]), "should return slice T truncated to len 3")
		})
	})

	Context("User supplied k is equal to allowable limit", func() {
		It("should return slices as is", func() {
			k := 5
			S := []string{"a", "b", "c", "d"}
			T := []string{"a", "b", "c", "d", "e"}
			actualS, actualT := truncate(S, T, k)
			Expect(actualS).To(Equal(S), "should return slice S as input")
			Expect(actualT).To(Equal(T), "should return slice T as input")
		})
	})

	Context("User supplied k is above to allowable limit", func() {
		It("should return slices as is", func() {
			k := 6
			S := []string{"a", "b", "c", "d"}
			T := []string{"a", "b", "c", "d", "e"}
			actualS, actualT := truncate(S, T, k)
			Expect(actualS).To(Equal(S), "should return slice S as input")
			Expect(actualT).To(Equal(T), "should return slice T as input")
		})
	})
})
