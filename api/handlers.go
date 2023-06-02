package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/9illes/log-parser/domain"
)

const Version = "0.1"

// HTTP struct expose HTTP handlers
type HTTP struct {
	Version string `json:"version"`
}

func NewHTTP() *HTTP {
	return &HTTP{
		Version: Version,
	}
}

// LoadHandler is the handler for the load journal request, it return a paginated list of requests in JSON format
func (api *HTTP) LoadHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	response := &Response{}

	requests, err := domain.LoadJournal(
		r.URL.Query().Get("site"),
		r.URL.Query().Get("type"),
		r.URL.Query().Get("date"),
	)

	if err != nil {
		response.ErrorMsg = "can't load journal"
	}

	pageNumber := 1
	if r.URL.Query().Has("page") {
		strPageNumber := r.URL.Query().Get("page")
		pageNumber, err = strconv.Atoi(strPageNumber)
	}

	requests, nextPage := domain.Paginate(requests, pageNumber)

	response.Entries = requests
	response.NextPage = nextPage

	w.Header().Set("Content-Type", "application/json")

	b, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, string(b))
}
