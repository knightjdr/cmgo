package dbgenes_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDbgenes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dbgenes Suite")
}
