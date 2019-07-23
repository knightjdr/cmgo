package tsnecytoscape

import (
	"github.com/knightjdr/cmgo/internal/pkg/read/tsne"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Determine network transformation", func() {
	Context("negative min x and y, positive max x", func() {
		It("should parse arguments", func() {
			coordinates := map[string]tsne.Coordinate{
				"a": tsne.Coordinate{X: 20.7, Y: -100},
				"b": tsne.Coordinate{X: -200, Y: 0},
				"c": tsne.Coordinate{X: 300, Y: 200.5},
				"d": tsne.Coordinate{X: 75, Y: 100},
			}
			expected := map[string]float64{
				"scale":      2,
				"translateX": 200,
				"translateY": 100,
			}
			Expect(networkTransformation(coordinates, 1000)).To(Equal(expected))
		})
	})

	Context("positive min x and y, positive max x", func() {
		It("should parse arguments", func() {
			coordinates := map[string]tsne.Coordinate{
				"a": tsne.Coordinate{X: 90, Y: 100},
				"b": tsne.Coordinate{X: 200, Y: 10},
				"c": tsne.Coordinate{X: 275, Y: 200.5},
				"d": tsne.Coordinate{X: 75, Y: 100},
			}
			expected := map[string]float64{
				"scale":      5,
				"translateX": -75,
				"translateY": -10,
			}
			Expect(networkTransformation(coordinates, 1000)).To(Equal(expected))
		})
	})
})
