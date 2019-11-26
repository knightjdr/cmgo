package pfam

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
)

// Features is a map of UniProt IDs with domain and motif information.
type Features map[string]Regions

type pfamService struct {
	Result *Features
	URL    string
}

func fetchRegions(s *pfamService, ids []string) {
	if s.URL == "" {
		s.URL = "http://pfam.xfam.org/protein"
	}

	numberOfIDs := len(ids)
	features := make(Features, numberOfIDs)

	client := &http.Client{}
	for i, id := range ids {
		url := fmt.Sprintf("%s/%s/graphic", s.URL, id)
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Content-Type", "application/json")
		res, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()

		var result Response
		json.NewDecoder(res.Body).Decode(&result)
		if len(result) > 0 {
			result.AddDomainNames()
			features[id] = result[0]
		} else {
			features[id] = Regions{}
		}
		showProgress(i, numberOfIDs)
	}

	s.Result = &features
}

func showProgress(i, queries int) {
	iteration := i + 1
	if math.Mod(float64(iteration), 25) == 0 {
		fmt.Println(fmt.Sprintf("%d of %d", iteration, queries))
	}
}
