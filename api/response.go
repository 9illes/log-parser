package api

import "github.com/9illes/log-parser/domain"

// Response is the response to a load request
type Response struct {
	Entries  []*domain.Request `json:"entries"`
	NextPage int               `json:"next_page"`
	ErrorMsg string            `json:"error"`
}
