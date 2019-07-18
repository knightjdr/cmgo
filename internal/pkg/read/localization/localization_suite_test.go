package localization_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLocalization(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Localization Suite")
}
