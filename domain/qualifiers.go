package domain

import (
	"regexp"
	"strconv"
)

// Qualifier is a function that filter on properties of a request.
type Qualifier func(*Request)

// IsMedia is an qualifier that determines if the request is for a media file.
func IsMediaRequest(r *Request) {
	re := regexp.MustCompile(`(?i)\.(css|js|pdf|jpg|jpeg|png|webp|woff|woff2|ttf|eot|pdf)`)
	r.IsMedia = re.Match([]byte(r.URI))
	r.IsPage = !r.IsMedia
}

// IsBot is an qualifier that determines if the request is from a bot.
func IsBotRequest(r *Request) {
	re := regexp.MustCompile(`(?i)(bot|spider|crawl|golang|curl|statuscake)`)
	r.IsBot = re.Match([]byte(r.UserAgent))
	// IsGoogleBot is an qualifier that determines if the request is from GoogleBot.
	re = regexp.MustCompile(`(?i)(googlebot)`)
	r.IsGoogleBot = re.Match([]byte(r.UserAgent))
}

// IsCacheHit is an qualifier that determines if the request is a cache hit.
func IsCacheHitRequest(r *Request) {
	re := regexp.MustCompile(`(?i)(hit)`)
	r.IsCacheHit = re.Match([]byte(r.EdgeDefaultResultType))
}

func IsSlowRequest(threshold float64) Qualifier {
	return func(r *Request) {
		ttfb, _ := strconv.ParseFloat(r.TTFB, 8)
		r.IsSlow = ttfb > threshold
	}
}
