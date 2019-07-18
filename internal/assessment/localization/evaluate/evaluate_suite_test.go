package evaluate_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestEvaluate(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Evaluate Suite")
}
