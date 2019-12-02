package countgo_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCountgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Countgo Suite")
}
