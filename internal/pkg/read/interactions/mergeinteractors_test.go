package interactions

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Merge BioGRID and IntAct interactors", func() {
	It("should merge two data sources", func() {
		biogrid := map[string][]string{
			"ACVR1":  []string{"FNTA"},
			"ARF3":   []string{"ARFIP1", "ARFIP2"},
			"ARFIP1": []string{"ARF3"},
			"ARFIP2": []string{"ARF3"},
			"FLNC":   []string{"MAP2K4"},
			"FNTA":   []string{"ACVR1"},
			"MAP2K4": []string{"FLNC"},
		}
		intact := map[string][]string{
			"ALDOA":  []string{"XRN1"},
			"ARF3":   []string{"ARFIP1", "ARFIP2"},
			"ARFIP1": []string{"ARF3"},
			"ARFIP2": []string{"ARF3"},
			"XRN1":   []string{"ALDOA"},
		}

		expected := map[string][]string{
			"ACVR1":  []string{"FNTA"},
			"ALDOA":  []string{"XRN1"},
			"ARF3":   []string{"ARFIP1", "ARFIP2"},
			"ARFIP1": []string{"ARF3"},
			"ARFIP2": []string{"ARF3"},
			"FLNC":   []string{"MAP2K4"},
			"FNTA":   []string{"ACVR1"},
			"MAP2K4": []string{"FLNC"},
			"XRN1":   []string{"ALDOA"},
		}
		Expect(mergeInteractors(biogrid, intact)).To(Equal(expected))
	})
})
