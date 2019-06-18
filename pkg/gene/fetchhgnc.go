package gene

import (
	"encoding/json"
	"log"
	"net/http"
)

type hgncService struct {
	Result []map[string]string
	URL    string
}

func fetchHGNC(s *hgncService) {
	if s.URL == "" {
		s.URL = "http://rest.genenames.org/fetch/status/Approved"
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", s.URL, nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	var result Data
	json.NewDecoder(res.Body).Decode(&result)
	result.Response.DefineUniProt()

	s.Result = result.Response.ParseIDtoMap()
}
