package preys

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Summarize interactions", func() {
	It("should calculate metrics for each prey", func() {
		baits := map[string]string{
			"128_7909": "bira-flag",
			"128_7910": "bira-flag",
			"128_8000": "bira-gfp",
			"128_8001": "bira-gfp",
			"128_8301": "empty",
			"128_8302": "empty",
		}
		interactions := map[string]map[string]int{
			"128_7909": map[string]int{
				"BirA_R118G_H0QFJ5": 100,
				"APC":               10,
			},
			"128_7910": map[string]int{
				"BirA_R118G_H0QFJ5": 50,
			},
			"128_8000": map[string]int{
				"BirA_R118G_H0QFJ5": 40,
			},
			"128_8001": map[string]int{
				"BirA_R118G_H0QFJ5": 200,
			},
			"128_8301": map[string]int{
				"APC": 20,
			},
			"128_8302": map[string]int{
				"BirA_R118G_H0QFJ5": 30,
				"APC":               30,
			},
		}

		expected := map[string]*preyInteraction{
			"BirA_R118G_H0QFJ5": &preyInteraction{
				Average: average{
					BirAFlag: 75,
					BirAGFP:  120,
					Empty:    30,
					Overall:  70,
				},
				BirAFlag: []int{50, 100},
				BirAGFP:  []int{40, 200},
				Empty:    []int{30},
				Max: maxMin{
					BirAFlag: 100,
					BirAGFP:  200,
					Empty:    30,
					Overall:  200,
				},
				Min: maxMin{
					BirAFlag: 50,
					BirAGFP:  40,
					Empty:    0,
					Overall:  0,
				},
			},
			"APC": &preyInteraction{
				Average: average{
					BirAFlag: 10,
					BirAGFP:  0,
					Empty:    25,
					Overall:  10,
				},
				BirAFlag: []int{10},
				BirAGFP:  []int{},
				Empty:    []int{20, 30},
				Max: maxMin{
					BirAFlag: 10,
					BirAGFP:  0,
					Empty:    30,
					Overall:  30,
				},
				Min: maxMin{
					BirAFlag: 0,
					BirAGFP:  0,
					Empty:    20,
					Overall:  0,
				},
			},
		}

		Expect(summarize(baits, interactions)).To(Equal(expected))
	})
})
