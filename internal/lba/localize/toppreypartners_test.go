package localize

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Top prey partners", func() {
	Context("allow all preys", func() {
		It("should return the top associated preys", func() {
			baitsPerPrey := map[string][]string{
				"a": []string{"A", "B", "C", "D"},
				"b": []string{"A", "C", "D"},
				"c": []string{"A", "B"},
				"d": []string{"E"},
			}
			preysPerBait := map[string][]string{
				"A": []string{"a", "b", "c"},
				"B": []string{"a", "c"},
				"C": []string{"a", "b"},
				"D": []string{"a", "b"},
				"E": []string{"d"},
			}

			actualFoldChange, actualTopPreys := topPreyPartners(baitsPerPrey, preysPerBait, 2, 0)
			expectedFoldChange := map[string]map[string]float64{
				"a": map[string]float64{
					"b": 1.25,
					"c": 1.25,
				},
				"b": map[string]float64{
					"a": 1.25,
					"c": 0.833,
				},
				"c": map[string]float64{
					"a": 1.25,
					"b": 0.833,
				},
				"d": map[string]float64{},
			}
			expectedTopPreys := map[string][]string{
				"a": []string{"b", "c"},
				"b": []string{"a", "c"},
				"c": []string{"a", "b"},
				"d": []string{},
			}

			for prey, partnerDetails := range actualFoldChange {
				for partner, fc := range partnerDetails {
					Expect(fc).To(BeNumerically("~", expectedFoldChange[prey][partner], 0.01), fmt.Sprintf("prey %s and partner %s", prey, partner))
				}
			}
			Expect(actualTopPreys).To(Equal(expectedTopPreys), "should return top preys by occurence")
		})
	})

	Context("require minimum fold change", func() {
		It("should return the top associated preys", func() {
			baitsPerPrey := map[string][]string{
				"a": []string{"A", "B", "C", "D"},
				"b": []string{"A", "C", "D"},
				"c": []string{"A", "B"},
				"d": []string{"E"},
			}
			preysPerBait := map[string][]string{
				"A": []string{"a", "b", "c"},
				"B": []string{"a", "c"},
				"C": []string{"a", "b"},
				"D": []string{"a", "b"},
				"E": []string{"d"},
			}

			_, actualTopPreys := topPreyPartners(baitsPerPrey, preysPerBait, 2, 1)
			expectedTopPreys := map[string][]string{
				"a": []string{"b", "c"},
				"b": []string{"a"},
				"c": []string{"a"},
				"d": []string{},
			}

			Expect(actualTopPreys).To(Equal(expectedTopPreys), "should return top preys by occurence")
		})
	})
})
