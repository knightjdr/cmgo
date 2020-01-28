package isolation_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestIsolation(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Isolation Suite")
}
