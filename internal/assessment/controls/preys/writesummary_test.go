package preys

import (
	"github.com/knightjdr/cmgo/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Write control summary", func() {
	It("should write all preys with metrics", func() {
		oldFs := fs.Instance
		defer func() { fs.Instance = oldFs }()
		fs.Instance = afero.NewMemMapFs()

		fs.Instance.MkdirAll("test", 0755)

		summary := map[string]*preyInteraction{
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

		expected := "prey\taverage\tmax\tmin\taverage BirA-FLAG\tmax BirA-FLAG\tmin BirA-FLAG\taverage BirA-GFP\tmax BirA-GFP\tmin BirA-GFP\taverage empty\tmax empty\tmin empty\tBirA-Flag\tBirA-FGFP\tempty\n" +
			"BirA_R118G_H0QFJ5\t70.00\t200\t0\t75.00\t100\t50\t120.00\t200\t40\t30.00\t30\t0\t50|100\t40|200\t30\n" +
			"APC\t10.00\t30\t0\t10.00\t10\t0\t0.00\t0\t0\t25.00\t30\t20\t10\t\t20|30\n"

		writeSummary(summary, "test/out.txt")
		bytes, _ := afero.ReadFile(fs.Instance, "test/out.txt")
		Expect(string(bytes)).To(Equal(expected))
	})
})
