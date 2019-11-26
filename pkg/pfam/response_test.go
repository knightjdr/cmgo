package pfam

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Response", func() {
	It("should add metadata identifier to domain", func() {
		json := Response{
			Regions{
				Domains: []Domain{
					{
						End:      161,
						Metadata: Metadata{Identified: "Domain 1"},
						Name:     "",
						Start:    14,
					},
					{
						End:      200,
						Metadata: Metadata{Identified: "Domain 2"},
						Name:     "",
						Start:    170,
					},
				},
			},
		}

		expected := Response{
			Regions{
				Domains: []Domain{
					{
						End:      161,
						Metadata: Metadata{Identified: "Domain 1"},
						Name:     "Domain 1",
						Start:    14,
					},
					{
						End:      200,
						Metadata: Metadata{Identified: "Domain 2"},
						Name:     "Domain 2",
						Start:    170,
					},
				},
			},
		}
		json.AddDomainNames()
		Expect(json).To(Equal(expected))
	})
})
