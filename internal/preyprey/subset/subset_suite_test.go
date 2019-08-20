package subset_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSubset(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Subset Suite")
}
