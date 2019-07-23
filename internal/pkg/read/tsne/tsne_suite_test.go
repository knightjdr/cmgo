package tsne_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTsne(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tsne Suite")
}
