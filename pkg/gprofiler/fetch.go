// Package gprofiler performs and term enrichment at g:Profiler
package gprofiler

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// Service defines the POST body, URL and contains parsed results.
type Service struct {
	Body   RequestBody
	Result []EnrichedTerm
	URL    string
}

// Fetch submits a gene list and parses results from g:Profiler
func Fetch(s *Service) {
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

	s.Result = result.Result
}
