package gene

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var maxQueryLength = 1500 - 70

type uniprotService struct {
	Query  []string
	Result map[string][]string
	URL    string
}

func fetchUniProt(s *uniprotService) {
	mapping := make(map[string][]string, 0)

	for start := 0; start < len(s.Query); {
		querySubset := []string{}
		querySubset, start = getNextQuerySubset(s.Query, start)
		query := strings.Join(querySubset, "%20")

		url := s.URL
		if url == "" {
			url = fmt.Sprintf("https://www.uniprot.org/uploadlists?from=P_REFSEQ_AC&to=ACC&format=tab&query=%s", query)
		}

		client := &http.Client{}
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("User-Agent", "email:jknight@lunenfeld.ca")
		res, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}
		defer res.Body.Close()

		responseData, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatalln(err)
		}
		responseString := string(responseData)

		parseUniProtMapping(&mapping, responseString)
	}

	s.Result = mapping
}

func getNextQuerySubset(query []string, start int) ([]string, int) {
	subset := make([]string, 0)

	end := start
	queryLength := 0
	for i := start; i < len(query); i++ {
		queryLength += len(query[i])
		if queryLength > maxQueryLength {
			break
		}

		end = i
		subset = append(subset, query[i])
	}

	return subset, end + 1
}

func parseUniProtMapping(mapping *map[string][]string, str string) {
	lines := strings.Split(str, "\n")

	for i := 1; i < len(lines); i++ {
		addUniProtMapping(mapping, lines[i])
	}
}

func addUniProtMapping(mapping *map[string][]string, line string) {
	fields := strings.Split(line, "\t")

	if len(fields) > 1 {
		refseq := fields[0]
		uniprot := fields[1]

		if _, ok := (*mapping)[refseq]; !ok {
			(*mapping)[refseq] = make([]string, 0)
		}
		(*mapping)[refseq] = append((*mapping)[refseq], uniprot)
	}
}
