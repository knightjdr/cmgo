package random_test

import (
	"github.com/knightjdr/cmgo/pkg/random"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SliceIntN", func() {
	It("should generate a slice of length N", func() {
		n := 4
		s := []int{1, 2, 3, 4, 5}
		actual := random.SliceIntN(s, n)
		Expect(len(actual)).To(Equal(4))
	})

	Describe("if numbers are randomly selected, expect occurrences to be equal", func() {
		It("should see each element occurring with approximately expected frequency", func() {
			n := 4
			s := []int{0, 1, 2, 3, 4}
			occurence := make([]int, 5)

			for i := 0; i < 1000; i++ {
				randomSlice := random.SliceIntN(s, n)
				for _, item := range randomSlice {
					occurence[item]++
				}
			}

			for i := 0; i < 5; i++ {
				Expect(occurence[i]).To(BeNumerically("~", 800, 20))
			}
		})
	})

})
