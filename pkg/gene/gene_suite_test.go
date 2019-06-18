package gene_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGene(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gene Suite")
}
