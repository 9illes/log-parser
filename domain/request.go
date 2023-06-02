package domain

import (
	"log"
	"net/url"

	"github.com/9illes/log-parser/cloudfront"
)

type Properties map[string]interface{}

type Request struct {
	Datetime              string     `json:"datetime"`
	Date                  string     `json:"date"`
	Time                  string     `json:"time"`
	IP                    string     `json:"ip"`
	Protocol              string     `json:"protocol"`
	UserAgent             string     `json:"user_agent"`
	HostHeader            string     `json:"host_header"`
	Method                string     `json:"method"`
	Status                string     `json:"status"`
	Cache                 string     `json:"cache"`
	URI                   string     `json:"uri"`
	Referer               string     `json:"referer"`
	ContentType           string     `json:"content_type"`
	TTFB                  string     `json:"ttfb"`
	EdgeDefaultResultType string     `json:"-"`
	IsCacheHit            bool       `json:"is_cache_hit"`
	IsPage                bool       `json:"is_page"`
	IsMedia               bool       `json:"is_media"`
	IsBot                 bool       `json:"is_bot"`
	IsGoogleBot           bool       `json:"is_google_bot"`
	IsSlow                bool       `json:"is_slow"`
	Properties            Properties `json:"-"`
}

func ParseLines(lines []*cloudfront.Line, qualifiers ...Qualifier) ([]*Request, error) {
	var requests []*Request

	for _, line := range lines {
		request, err := fromLine(line)
		if err != nil {
			log.Fatal(err)
		}
		for _, qualifier := range qualifiers {
			qualifier(request)
		}
		requests = append(requests, request)
	}

	return requests, nil
}

func fromLine(line *cloudfront.Line) (*Request, error) {

	userAgent, err := url.PathUnescape(line.UserAgent)
	if err != nil {
		log.Fatal(err)
	}

	contentType, err := url.PathUnescape(line.ContentType)
	if err != nil {
		log.Fatal(err)
	}

	return &Request{
		Datetime:              line.Date + " " + line.Time,
		Date:                  line.Date,
		Time:                  line.Time,
		IP:                    line.IP,
		Protocol:              line.Protocol,
		UserAgent:             userAgent,
		HostHeader:            line.HostHeader,
		Method:                line.Method,
		Status:                line.Status,
		Cache:                 line.EdgeDefaultResultType,
		URI:                   line.URI,
		Referer:               line.Referer,
		ContentType:           contentType,
		TTFB:                  line.TTFB,
		EdgeDefaultResultType: line.EdgeDefaultResultType,
		Properties:            Properties{},
	}, nil
}
