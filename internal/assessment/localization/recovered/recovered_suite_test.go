package recovered_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRecovered(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Recovered Suite")
}
