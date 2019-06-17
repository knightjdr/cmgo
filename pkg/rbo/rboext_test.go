package rbo_test

import (
	"github.com/knightjdr/cmgo/pkg/rbo"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RBOext", func() {
	Context("Identical lists", func() {
		It("should return score of 1", func() {
			k := 0
			p := 0.9
			S := []string{"a", "b", "c", "d", "e"}
			T := []string{"a", "b", "c", "d", "e"}
			Expect(rbo.RBOext(S, T, p, k)).To(Equal(float64(1)))
		})

		It("should return score of 1 when floating point errors yielding scores > 1", func() {
			k := 0
			p := 0.9
			S := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t"}
			T := []string{"a", "b"}
			Expect(rbo.RBOext(S, T, p, k)).To(Equal(float64(1)))
		})
	})

	Context("Identical lists but different order", func() {
		It("should return score when first elements are swapped", func() {
			k := 0
			p := 0.9
			S := []string{"a", "b", "c", "d", "e"}
			T := []string{"b", "a", "c", "d", "e"}
			Expect(rbo.RBOext(S, T, p, k)).To(BeNumerically("~", 0.9, 0.00001))
		})

		It("should return score when elements have different orders", func() {
			k := 0
			p := 0.9
			S := []string{"a", "b", "c", "d", "e"}
			T := []string{"b", "d", "a", "c", "e"}
			Expect(rbo.RBOext(S, T, p, k)).To(BeNumerically("~", 0.828, 0.00001))
		})
	})

	Context("Different length lists", func() {
		It("should return score when shared items are identical", func() {
			k := 0
			p := 0.9
			S := []string{"a", "b", "c", "d", "e"}
			T := []string{"b", "a", "c", "d", "e", "f"}
			Expect(rbo.RBOext(S, T, p, k)).To(BeNumerically("~", 0.9, 0.00001))
		})

		It("should return score when discordant items", func() {
			k := 0
			p := 0.9
			S := []string{"a", "b", "c", "d", "e"}
			T := []string{"b", "a", "x", "d", "y", "f"}
			Expect(rbo.RBOext(S, T, p, k)).To(BeNumerically("~", 0.592335, 0.00001))
		})
	})

	Context("Disjointed lists", func() {
		It("should return score of 0", func() {
			k := 0
			p := 0.9
			S := []string{"a", "b", "c", "d", "e"}
			T := []string{"v", "w", "x", "y", "z"}
			Expect(rbo.RBOext(S, T, p, k)).To(Equal(float64(0)))
		})

		It("should return score of 0 when one list is empty", func() {
			k := 0
			p := 0.9
			S := []string{"a", "b", "c", "d", "e"}
			T := []string{}
			Expect(rbo.RBOext(S, T, p, k)).To(Equal(float64(0)))
		})
	})
})
