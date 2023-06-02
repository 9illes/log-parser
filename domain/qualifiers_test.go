package domain

import (
	"testing"
)

func TestIsMediaRequest(t *testing.T) {

	var tests = []struct {
		name string
		uri  string
		want *Request
	}{
		{"a page request", "/foo", &Request{IsMedia: false, IsPage: true}},
		{"a png", "/foo.png", &Request{IsMedia: true, IsPage: false}},
		{"a png camelcase", "/foo.PnG", &Request{IsMedia: true, IsPage: false}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sut := &Request{
				URI: tt.uri,
			}

			IsMediaRequest(sut)

			if sut.IsMedia != tt.want.IsMedia {
				t.Errorf("%s : IsMedia got %t, want %t", tt.name, sut.IsMedia, tt.want.IsMedia)
			}
		})
	}
}

// Sample fuzz test
func FuzzIsMedia(f *testing.F) {

	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, URI string) {
		IsMediaRequest(&Request{URI: URI})
	})
}

func TestIsBotRequest(t *testing.T) {

	var tests = []struct {
		name      string
		userAgent string
		want      *Request
	}{
		{"browser", "Mozilla/5.0 (Windows NT 10.0; rv:112.0) Gecko/20100101 Firefox/112.0", &Request{IsBot: false}},
		{"bot", "garbageBOTgarbage", &Request{IsBot: true}},
		{"bot/crawler", "garbagecrawlergarbage", &Request{IsBot: true}},
		{"googlebot", " 	Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5X Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/W.X.Y.Z Mobile Safari/537.36 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)", &Request{IsBot: true}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sut := &Request{
				UserAgent: tt.userAgent,
			}

			IsBotRequest(sut)

			if sut.IsBot != tt.want.IsBot {
				t.Errorf("%s : IsBot got %t, want %t", tt.name, sut.IsBot, tt.want.IsBot)
			}
		})
	}
}
