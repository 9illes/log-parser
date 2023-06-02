package domain

// Qualifier is a function that filter on properties of a request.
type FilterFunc func(*Request) bool

type FilteredRequests struct {
	Requests []*Request
	Filename string
	filters  []FilterFunc
}

func NewFilteredRequests(filename string, filters ...FilterFunc) *FilteredRequests {
	return &FilteredRequests{
		Filename: filename,
		filters:  filters,
	}
}

func (f *FilteredRequests) Append(r *Request) {
	for _, filter := range f.filters {
		if filter(r) == false {
			return
		}
	}
	f.Requests = append(f.Requests, r)
}

func FilterIsPage(r *Request) bool {
	return !r.IsMedia
}

func FilterIsAssets(r *Request) bool {
	return r.IsMedia
}

func FilterIsBot(r *Request) bool {
	return r.IsBot
}

// IsGoogleBot is an qualifier that determines if the request is from GoogleBot.
func FilterIsGoogleBot(r *Request) bool {
	return r.IsGoogleBot
}

func FilterAlwaysFalse(r *Request) bool {
	return false
}

func FilterAlwaysTrue(r *Request) bool {
	return true
}
