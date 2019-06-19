package uv

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestUv(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Uv Suite")
}
