// Package gprofiler performs and term enrichment at g:Profiler
package gprofiler

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// Fetch submits a gene list and parse results from g:Profiler
func Fetch(body RequestBody) []EnrichedTerm {
	body.addDefaults()
	data, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}

	url := "https://biit.cs.ut.ee/gprofiler/api/gost/profile/"
	res, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	var result Response
	json.NewDecoder(res.Body).Decode(&result)

	return result.Result
}
