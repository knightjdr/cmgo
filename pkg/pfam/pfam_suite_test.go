package pfam_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPfam(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pfam Suite")
}
