package transmembrane

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
				"-basisMatrix", "basis.txt",
				"-cytosolicBaits", "a,b,c",
				"-cytosolicCompartments", "1,5",
				"-fdr", "0.05",
				"-lumenalBaits", "d,e",
				"-lumenalCompartments", "2",
				"-minRankValue", "0.05",
				"-outFile", "out.txt",
				"-saint", "saint.txt",
			}
			fileOptions := map[string]interface{}{}

			expected := parameters{
				basisMatrix:           "basis.txt",
				cytosolicBaits:        []string{"a", "b", "c"},
				cytosolicCompartments: []string{"1", "5"},
				fdr:                   0.05,
				lumenalBaits:          []string{"d", "e"},
				lumenalCompartments:   []string{"2"},
				minRankValue:          0.05,
				outFile:               "out.txt",
				saint:                 "saint.txt",
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
				"-basisMatrix", "basis.txt",
				"-saint", "saint.txt",
			}
			fileOptions := map[string]interface{}{}

			options, err := parseFlags(fileOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(options.fdr).To(Equal(0.01), "should set default minimum FDR")
			Expect(options.minRankValue).To(Equal(0.15), "should set default minimum NMF score")
			Expect(options.outFile).To(Equal("transmembrane.txt"), "should set default out file")
		})
	})

	Context("missing required command line arguments", func() {
		It("should report error", func() {
			os.Args = []string{
				"cmd",
			}
			fileOptions := map[string]interface{}{}

			_, err := parseFlags(fileOptions)
			Expect(err).Should(HaveOccurred())
		})
	})

	Context("argument passed via input file", func() {
		It("should set variables from file", func() {
			os.Args = []string{
				"cmd",
			}
			fileOptions := map[string]interface{}{
				"basisMatrix":           "file-basis.txt",
				"cytosolicBaits":        "a,b,c",
				"cytosolicCompartments": "1,5",
				"fdr":                   "0.05",
				"lumenalBaits":          "d,e",
				"lumenalCompartments":   "2",
				"minRankValue":          "0.05",
				"outFile":               "file-out.txt",
				"saint":                 "file-saint.txt",
			}

			expected := parameters{
				basisMatrix:           "file-basis.txt",
				cytosolicBaits:        []string{"a", "b", "c"},
				cytosolicCompartments: []string{"1", "5"},
				fdr:                   0.05,
				lumenalBaits:          []string{"d", "e"},
				lumenalCompartments:   []string{"2"},
				minRankValue:          0.05,
				outFile:               "file-out.txt",
				saint:                 "file-saint.txt",
			}
			options, err := parseFlags(fileOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(options).To(Equal(expected), "should set options")
		})
	})
})
