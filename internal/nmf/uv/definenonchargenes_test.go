package uv

import (
	"sort"

	readNMF "github.com/knightjdr/cmgo/internal/pkg/read/nmf"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Define non characterizing genes", func() {
	It("should return a 2D slice with non-characterizing genes for each rank", func() {
		charaterizingGenes := [][]string{
			{"a", "b", "c"},
			{"c", "d", "e"},
		}
		nmfLocalizations := readNMF.NMFlocalization{
			"a": readNMF.GeneLocalization{
				Compartment: 1,
			},
			"b": readNMF.GeneLocalization{
				Compartment: 1,
			},
			"c": readNMF.GeneLocalization{
				Compartment: 2,
			},
			"d": readNMF.GeneLocalization{
				Compartment: 2,
			},
			"e": readNMF.GeneLocalization{
				Compartment: 2,
			},
			"f": readNMF.GeneLocalization{
				Compartment: 1,
			},
			"g": readNMF.GeneLocalization{
				Compartment: 2,
			},
			"h": readNMF.GeneLocalization{
				Compartment: 1,
			},
		}

		actual := defineNonCharacterizingGenes(charaterizingGenes, nmfLocalizations)
		sort.Strings(actual[0])
		sort.Strings(actual[1])
		expected := [][]string{
			{"f", "h"},
			{"g"},
		}
		Expect(actual).To(Equal(expected))
	})
})
