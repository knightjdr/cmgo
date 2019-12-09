package prediction_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPrediction(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Prediction Suite")
}
