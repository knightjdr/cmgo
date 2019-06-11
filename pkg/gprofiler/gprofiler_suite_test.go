package gprofiler_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGprofiler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gprofiler Suite")
}
