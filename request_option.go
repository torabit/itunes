package itunes

import (
	"net/url"
	"strconv"
)

type RequestOption func(*requestOption)

type requestOption struct {
	urlParams url.Values
}

// The URL-encoded text string you want to search for. For example: jack+johnson.
func Term(term string) RequestOption {
	return func(o *requestOption) {
		o.urlParams.Set("term", term)
	}
}

// The two-letter country code for the store you want to search.
// The search uses the default store front for the specified country.
// For example: US. The default is US.
func Country(code string) RequestOption {
	return func(o *requestOption) {
		o.urlParams.Set("country", code)
	}
}

// The media type you want to search for. For example: movie. The default is all.
func Media(media string) RequestOption {
	return func(o *requestOption) {
		o.urlParams.Set("media", media)
	}
}

// The type of results you want returned, relative to the specified media type.
// For example: movieArtist for a movie media type search.
// The default is the track entity associated with the specified media type.
func Entity(entity string) RequestOption {
	return func(o *requestOption) {
		o.urlParams.Set("entity", entity)
	}
}

// The attribute you want to search for in the stores, relative to the specified media type.
// For example, if you want to search for an artist by name specify entity=allArtist&attribute=allArtistTerm.
// In this example, if you search for term=maroon, iTunes returns "Maroon 5" in the search results,
// instead of all artists who have ever recorded a song with the word “maroon” in the title.
// The default is all attributes associated with the specified media type.
func Attribute(attribute string) RequestOption {
	return func(o *requestOption) {
		o.urlParams.Set("attribute", attribute)
	}
}

// The number of search results you want the iTunes Store to return.
// For example: 25. The default is 50.
func Limit(amount int) RequestOption {
	return func(o *requestOption) {
		if amount > 200 {
			amount = 200
		}

		if amount < 1 {
			amount = 1
		}
		o.urlParams.Set("limit", strconv.Itoa(amount))
	}
}

// The language, English or Japanese, you want to use when returning search results.
// Specify the language using the five-letter codename. For example: en_us.
// The default is en_us (English).
func Lang(lang string) RequestOption {
	return func(o *requestOption) {
		o.urlParams.Set("lang", lang)
	}
}

// The search result key version you want to receive back from your search.
// The default is 2.
func Version(version int) RequestOption {
	return func(o *requestOption) {
		v := 2

		if version == 1 {
			v = 1
		}
		o.urlParams.Set("version", strconv.Itoa(v))
	}
}

func Explicit(isExplicit bool) RequestOption {
	return func(o *requestOption) {
		explicit := "Yes"
		if !isExplicit {
			explicit = "No"
		}

		o.urlParams.Set("explicit", explicit)
	}
}

func processOptions(options ...RequestOption) requestOption {
	o := requestOption{
		urlParams: url.Values{},
	}
	for _, opt := range options {
		opt(&o)
	}

	return o
}
