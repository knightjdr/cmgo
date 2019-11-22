package moonlighting_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMoonlighting(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Moonlighting Suite")
}
