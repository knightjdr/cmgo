package goenrich_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGoenrich(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Goenrich Suite")
}
