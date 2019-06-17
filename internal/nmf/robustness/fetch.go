package robustness

import (
	"github.com/knightjdr/cmgo/pkg/gprofiler"
)

func fetchWrapper(s *gprofiler.Service) {
	gprofiler.Fetch(s)
}

var fetch = fetchWrapper