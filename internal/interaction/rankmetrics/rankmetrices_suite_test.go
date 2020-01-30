package rankmetrics_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRankMetrics(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RankMetrics Suite")
}
