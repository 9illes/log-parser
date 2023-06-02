package domain

import (
	"regexp"
	"strconv"
)

// Insight is a function that filter on properties of a request.
type Insight func(*Request)

// IsMedia is an insight that determines if the request is for a media file.
func IsMedia(r *Request) {
	re := regexp.MustCompile(`(?i)\.(css|js|pdf|jpg|jpeg|png|webp|woff|woff2|ttf|eot|pdf)`)
	r.Properties["is_media"] = re.Match([]byte(r.URI))
	r.Properties["is_page"] = !r.Properties["is_media"].(bool)
}

// IsBot is an insight that determines if the request is from a bot.
func IsBot(r *Request) {
	re := regexp.MustCompile(`(?i)(bot|spider|crawl|golang|curl|statuscake)`)
	r.Properties["is_bot"] = re.Match([]byte(r.UserAgent))
}

func IsCacheHit(r *Request) {
	re := regexp.MustCompile(`(?i)(hit)`)
	r.Properties["is_cache_hit"] = re.Match([]byte(r.EdgeDefaultResultType))
}

func IsSlow(threshold float64) Insight {
	return func(r *Request) {
		ttfb, _ := strconv.ParseFloat(r.TTFB, 8)
		r.Properties["is_slow"] = ttfb > threshold
	}
}
