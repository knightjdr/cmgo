package tsnecytoscape_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTsnecytoscape(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tsnecytoscape Suite")
}
