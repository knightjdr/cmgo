package uniprot_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestUniprot(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Uniprot Suite")
}
