package goenrich

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
				"-baits", "baitA,baitB",
				"-namespace", "GO:BP",
				"-outFile", "out.txt",
				"-saint", "saint.txt",
			}
			fileOptions := map[string]interface{}{}

			expected := parameters{
				baits:     []string{"baitA", "baitB"},
				namespace: "GO:BP",
				outFile:   "out.txt",
				saint:     "saint.txt",
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
				"-baits", "baitA,baitB",
				"-saint", "saint.txt",
			}
			fileOptions := map[string]interface{}{}

			options, err := parseFlags(fileOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(options.namespace).To(Equal("GO:CC"), "should set default GO namespace")
			Expect(options.outFile).To(Equal("go-enrichment.txt"), "should set default out file")
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
				"baits":     "baitA,baitB",
				"namespace": "GO:BP",
				"outFile":   "file-out.txt",
				"saint":     "file-saint.txt",
			}

			expected := parameters{
				baits:     []string{"baitA", "baitB"},
				namespace: "GO:BP",
				outFile:   "file-out.txt",
				saint:     "file-saint.txt",
			}
			options, err := parseFlags(fileOptions)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(options).To(Equal(expected), "should set options")
		})
	})
})
