package nmf_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestNmf(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Nmf Suite")
}
