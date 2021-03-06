package genes

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
				"-background", "background.txt",
				"-genes", "genes.txt",
				"-namespace", "BP",
				"-outFile", "out.txt",
			}
			fileOptions := map[string]interface{}{}

			expected := parameters{
				background: "background.txt",
				genes:      "genes.txt",
				namespace:  "BP",
				outFile:    "out.txt",
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
				"-background", "background.txt",
				"-genes", "genes.txt",
			}
			fileOptions := map[string]interface{}{}

			options, err := parseFlags(fileOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(options.namespace).To(Equal("CC"), "should set default namespace")
			Expect(options.outFile).To(Equal("enrichment.txt"), "should set default out file")
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
				"background": "file-background.txt",
				"genes":      "file-genes.txt",
				"namespace":  "BP",
				"outFile":    "file-out.txt",
			}

			expected := parameters{
				background: "file-background.txt",
				genes:      "file-genes.txt",
				namespace:  "BP",
				outFile:    "file-out.txt",
			}
			options, err := parseFlags(fileOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(options).To(Equal(expected), "should set options")
		})
	})
})
