package rbo_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRbo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rbo Suite")
}
