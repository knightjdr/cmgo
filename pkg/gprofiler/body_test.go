package gprofiler_test

import (
	"github.com/knightjdr/cmgo/pkg/gprofiler"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Body", func() {
	Describe("AddDefaults", func() {
		Context("no user defined request data", func() {
			It("should add default values to POST body", func() {
				body := gprofiler.RequestBody{}
				body.AddDefaults()

				expected := gprofiler.RequestBody{
					DomainScope:                 "annotated",
					Organism:                    "hsapiens",
					SignificanceThresholdMethod: "gSCS",
					UserThreshold:               0.01,
				}
				Expect(body).To(Equal(expected))
			})
		})

		Context("user defined background", func() {
			It("should set DomainScope to custom when user supplies a background", func() {
				body := gprofiler.RequestBody{
					Background:  []string{"a", "b", "c"},
					DomainScope: "annotated",
				}
				body.AddDefaults()

				Expect(body.DomainScope).To(Equal("custom"))
			})
		})
	})
})
