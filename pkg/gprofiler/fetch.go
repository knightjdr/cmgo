// Package gprofiler performs and term enrichment at g:Profiler
package gprofiler

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// Service handles
type Service struct {
	Body RequestBody
	URL  string
}

// Fetch submits a gene list and parse results from g:Profiler
func (s *Service) Fetch() []EnrichedTerm {
	if s.URL == "" {
		s.URL = "https://biit.cs.ut.ee/gprofiler/api/gost/profile/"
	}

	s.Body.AddDefaults()
	data, err := json.Marshal(s.Body)
	if err != nil {
		log.Fatal(err)
	}

	res, err := http.Post(s.URL, "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	var result Response
	json.NewDecoder(res.Body).Decode(&result)
	result.AddIntersectionGenes("query_1")

	return result.Result
}
