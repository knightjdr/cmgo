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
				"W": []string{"a", "b", "c"},
				"X": []string{"a", "b"},
				"Y": []string{"a", "c"},
				"Z": []string{"b", "d"},
			}
			preysPerBait := map[string][]string{
				"a": []string{"W", "X", "Y"},
				"b": []string{"W", "X", "Z"},
				"c": []string{"W", "Y"},
				"d": []string{"Z"},
			}

			actualFoldChange, actualTopPreys := topPreyPartners(baitsPerPrey, preysPerBait, 2, 0)
			expectedFoldChange := map[string]map[string]float64{
				"W": map[string]float64{
					"X": 1.33,
					"Y": 1.33,
					"Z": 0.67,
				},
				"X": map[string]float64{
					"W": 1.33,
					"Y": 1,
					"Z": 1,
				},
				"Y": map[string]float64{
					"W": 1.33,
					"X": 1,
				},
				"Z": map[string]float64{
					"W": 0.67,
					"X": 1,
				},
			}
			expectedTopPreys := map[string][]string{
				"W": []string{"X", "Y"},
				"X": []string{"W", "Y"},
				"Y": []string{"W", "X"},
				"Z": []string{"X", "W"},
			}

			for prey, partnerDetails := range actualFoldChange {
				for partner, fc := range partnerDetails {
					Expect(fc).To(BeNumerically("~", expectedFoldChange[prey][partner], 0.01), fmt.Sprintf("prey %s and partner %s", prey, partner))
				}
			}
			Expect(actualTopPreys).To(Equal(expectedTopPreys), "should return top preys by fold change")
		})
	})

	Context("require minimum fold change", func() {
		It("should return the top associated preys", func() {
			baitsPerPrey := map[string][]string{
				"W": []string{"a", "b", "c"},
				"X": []string{"a", "b"},
				"Y": []string{"a", "c"},
				"Z": []string{"b", "d"},
			}
			preysPerBait := map[string][]string{
				"a": []string{"W", "X", "Y"},
				"b": []string{"W", "X", "Z"},
				"c": []string{"W", "Y"},
				"d": []string{"Z"},
			}

			_, actualTopPreys := topPreyPartners(baitsPerPrey, preysPerBait, 3, 1)
			expectedTopPreys := map[string][]string{
				"W": []string{"X", "Y"},
				"X": []string{"W"},
				"Y": []string{"W"},
				"Z": []string{},
			}
			Expect(actualTopPreys).To(Equal(expectedTopPreys), "should return top preys by fold change")
		})
	})
})
