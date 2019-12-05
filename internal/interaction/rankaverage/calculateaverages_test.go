package rankaverage

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Calculate prey interaction ranks and averages", func() {
	It("should return stats for each prey", func() {
		preys := []string{"preyA", "preyB", "preyC"}
		sortedPreysPerBait := map[string][]string{
			"baitA": []string{"preyA", "preyC"},
			"baitB": []string{"preyA", "preyB", "preyC"},
			"baitC": []string{"preyB"},
			"baitD": []string{"preyA", "preyB"},
		}

		expectedSummary := map[string]preySummary{
			"preyA": preySummary{
				mean:  1,
				ranks: []int{1, 1, 1},
				sd:    0,
			},
			"preyB": preySummary{
				mean:  1.667,
				ranks: []int{1, 2, 2},
				sd:    0.577,
			},
			"preyC": preySummary{
				mean:  2.500,
				ranks: []int{2, 3},
				sd:    0.707,
			},
		}

		actualSummary, actualMean, actualSD := calculateAverages(preys, sortedPreysPerBait)
		Expect(actualSummary).To(Equal(expectedSummary))
		Expect(actualMean).To(Equal(1.625))
		Expect(actualSD).To(Equal(0.744))
	})
})
