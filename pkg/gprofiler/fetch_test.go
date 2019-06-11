package gprofiler_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"
	. "github.com/knightjdr/cmgo/pkg/gprofiler"
)

var _ = Describe("Fetch", func() {
	Describe("testing...", func() {
		It("should...", func() {
			body := RequestBody{
				Query: []string{"CASQ2", "CASQ1", "GSTO1", "DMD", "GSTM2"},
			}
			result := Fetch(body)
			fmt.Println(result)
		})
	})
})
