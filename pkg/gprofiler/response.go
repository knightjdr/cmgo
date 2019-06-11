package gprofiler

// Response for data received from gProfiler.
type Response struct {
	Result []EnrichedTerm
}

// EnrichedTerm contains information on an enriched term.
type EnrichedTerm struct {
	ID               string `json:"native"`
	IntersectionSize int    `json:"intersection_size"`
	Name             string
	QuerySize        int `json:"query_size"`
	Precision        float64
	Pvalue           float64 `json:"p_value"`
	Recall           float64
	Source           string
	TermSize         int `json:"term_size"`
}
