package preys_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPreys(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Preys Suite")
}
