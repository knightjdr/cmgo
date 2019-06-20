package bioplex_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBioplex(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Bioplex Suite")
}
