package uniprot

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net/http"
)

// Entries is a map of UniProt IDs with entry information.
type Entries map[string]Entry

type uniprotService struct {
	Result *Entries
	URL    string
}

func fetchUniprot(s *uniprotService, ids []string) {
	numberOfIDs := len(ids)
	entries := make(Entries, numberOfIDs)

	client := &http.Client{}
	for i, id := range ids {
		if id != "" {
			url := createURL(s, id)
			req := initializeRequest(url)

			res, err := client.Do(req)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()

			entries[id] = handleResponse(res.Body)
		}
		showProgress(i, numberOfIDs)
	}

	s.Result = &entries
}

func createURL(s *uniprotService, id string) string {
	if s.URL == "" {
		s.URL = "http://www.ebi.ac.uk/proteins/api/proteins"
	}

	return fmt.Sprintf("%s/%s", s.URL, id)
}

func initializeRequest(url string) *http.Request {
	req, _ := http.NewRequest("GET", url, nil)
	setHeaders(req)
	return req
}

func setHeaders(req *http.Request) {
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
}

func handleResponse(response io.Reader) Entry {
	var result Entry
	body := parseBody(response)
	json.Unmarshal(body, &result)
	result.ParseAdditionalFields(body)
	return result
}

func parseBody(body io.Reader) []byte {
	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		log.Fatal(err)
	}
	return bodyBytes
}

func showProgress(i, queries int) {
	iteration := i + 1
	if math.Mod(float64(iteration), 25) == 0 {
		fmt.Println(fmt.Sprintf("UniProt: %d of %d", iteration, queries))
	}
}
