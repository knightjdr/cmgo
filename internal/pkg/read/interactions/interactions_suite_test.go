package interactions_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestInteractions(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Interactions Suite")
}
