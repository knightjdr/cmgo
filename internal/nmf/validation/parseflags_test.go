package validation

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = BeforeSuite(func() {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
})

var _ = Describe("Parseflags", func() {
	Context("all command line arguments", func() {
		It("should parse arguments", func() {
			os.Args = []string{
				"cmd",
				"-basisMatrix", "basis.csv",
				"-maxGenesPerRank", "50",
				"-minRankValue", "0.5",
				"-outFile", "out.svg",
				"-withinRankMax", "0.6",
			}
			fileOptions := map[string]interface{}{}

			expected := parameters{
				basisMatrix:     "basis.csv",
				maxGenesPerRank: 50,
				minRankValue:    0.5,
				outFile:         "out.svg",
				withinRankMax:   0.6,
			}
			options, err := parseFlags(fileOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(options).To(Equal(expected), "should set options")
		})
	})

	Context("only required command line arguments", func() {
		It("should set defaults", func() {
			os.Args = []string{
				"cmd",
				"-basisMatrix", "basis.csv",
			}
			fileOptions := map[string]interface{}{}

			options, err := parseFlags(fileOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(options.maxGenesPerRank).To(Equal(100), "should set default maxGenesPerRank")
			Expect(options.minRankValue).To(Equal(0.25), "should set default minRankValue")
			Expect(options.withinRankMax).To(Equal(0.75), "should set default withinRankMax")
		})
	})

	Context("missing required command line arguments", func() {
		It("should set defaults", func() {
			os.Args = []string{
				"cmd",
			}
			fileOptions := map[string]interface{}{}

			_, err := parseFlags(fileOptions)
			Expect(err).Should(HaveOccurred())
		})
	})

	Context("argument passed via input file", func() {
		It("should set defaults", func() {
			os.Args = []string{
				"cmd",
			}
			fileOptions := map[string]interface{}{
				"basisMatrix":     "file-basis.csv",
				"maxGenesPerRank": 50,
				"minRankValue":    0.5,
				"outFile":         "file-out.svg",
				"withinRankMax":   0.6,
			}

			expected := parameters{
				basisMatrix:     "file-basis.csv",
				maxGenesPerRank: 50,
				minRankValue:    0.5,
				outFile:         "file-out.svg",
				withinRankMax:   0.6,
			}
			options, err := parseFlags(fileOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(options).To(Equal(expected), "should set options")
		})
	})
})
