package hydropathy

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHydropathy(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Hydropathy Suite")
}
