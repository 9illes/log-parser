package domain

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

const requestsPerPage = 200

// LoadJournal loads the journal json file
func LoadJournal(site, journalType, date string) ([]*Request, error) {
	content, err := ioutil.ReadFile(getJournalFullpath(site, journalType, date))
	if err != nil {
		log.Printf("Error when opening file: %+v", err)
		return []*Request{}, err
	}

	var requests []*Request
	err = json.Unmarshal(content, &requests)
	if err != nil {
		log.Printf("Error during Unmarshal(): %s", err)
	}

	return requests, nil
}

// Paginate returns a slice of requests and the next page number
func Paginate(requests []*Request, page int) ([]*Request, int) {

	length := len(requests)
	nextPage := page + 1

	if length == 0 {
		return []*Request{}, 0
	}

	begin := (page - 1) * requestsPerPage
	if begin > length {
		begin = length
	}

	end := begin + requestsPerPage
	if end > length {
		end = length
		nextPage = 0
	}

	return requests[begin:end], nextPage
}

// getJournalFullpath returns the full path to the journal file
func getJournalFullpath(site, journalType, date string) string {
	return fmt.Sprintf("./var/%s/%s-%s.json", site, date, journalType)
}
