package genes_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGenes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Genes Suite")
}
