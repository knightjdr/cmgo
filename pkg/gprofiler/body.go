package gprofiler

// RequestBody for POST data to gProfiler.
type RequestBody struct {
	AllResults                  bool     `json:"all_results"`
	Background                  []string `json:"background"`
	Combined                    bool     `json:"combined"`
	DomainScope                 string   `json:"domain_scope"` // 'annotated', 'known' or 'custom'.
	MeasureUnderrepresentation  bool     `json:"measure_underrepresentation"`
	NoEvidences                 bool     `json:"no_evidences"`
	NoIEA                       bool     `json:"no_iea"`
	Ordered                     bool     `json:"ordered"`
	Organism                    string   `json:"organism"`
	Query                       []string `json:"query"`
	SignificanceThresholdMethod string   `json:"significance_threshold_method"` // 'bonferroni', 'fdr' or 'gSCS'
	Sources                     []string `json:"sources"`                       // An empty source list will query all datasources.
	UserThreshold               float64  `json:"user_threshold"`
}

func (r *RequestBody) addDefaults() {
	if r.DomainScope == "" {
		r.DomainScope = "annotated"
	}
	if len(r.Background) > 0 && r.DomainScope != "custom" {
		r.DomainScope = "custom"
	}

	if r.Organism == "" {
		r.Organism = "hsapiens"
	}

	if r.SignificanceThresholdMethod == "" {
		r.SignificanceThresholdMethod = "gSCS"
	}

	if r.UserThreshold == 0 {
		r.UserThreshold = 0.01
	}
}
