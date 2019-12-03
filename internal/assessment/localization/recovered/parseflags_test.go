package recovered

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
				"-compartmentID", "GO:111111",
				"-genes", "genes.txt",
				"-goAnnotations", "annotations.txt",
				"-localizationID", "13",
				"-outFile", "out.txt",
			}
			fileOptions := map[string]interface{}{}

			expected := parameters{
				compartmentID:  "GO:111111",
				genes:          "genes.txt",
				goAnnotations:  "annotations.txt",
				localizationID: "13",
				outFile:        "out.txt",
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
				"-compartmentID", "GO:111111",
				"-genes", "genes.txt",
				"-goAnnotations", "annotations.txt",
				"-localizationID", "13",
			}
			fileOptions := map[string]interface{}{}

			options, err := parseFlags(fileOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(options.outFile).To(Equal("compartment-recovered.txt"), "should set output file name")
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
				"compartmentID":  "GO:111111",
				"genes":          "file-genes.txt",
				"goAnnotations":  "file-annotations.txt",
				"localizationID": "13",
				"outFile":        "file-out.txt",
			}

			expected := parameters{
				compartmentID:  "GO:111111",
				genes:          "file-genes.txt",
				goAnnotations:  "file-annotations.txt",
				localizationID: "13",
				outFile:        "file-out.txt",
			}
			options, err := parseFlags(fileOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(options).To(Equal(expected), "should set options")
		})
	})
})
