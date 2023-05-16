package itunes

import (
	"strconv"
)

// Lookup Options
type LookupOption func(*requestOption)

func ID(id int64) LookupOption {
	return func(o *requestOption) {
		o.urlParams.Set("id", strconv.FormatInt(id, 10))
	}
}

func AMGArtistID(amgArtistId int64) LookupOption {
	return func(o *requestOption) {
		o.urlParams.Set("amgArtistId", strconv.FormatInt(amgArtistId, 10))
	}
}

func AMGAlbumID(amgAlbumId int64) LookupOption {
	return func(o *requestOption) {
		o.urlParams.Set("amgAlbumId", strconv.FormatInt(amgAlbumId, 10))
	}
}

func AMGVideoID(amgVideoId int64) LookupOption {
	return func(o *requestOption) {
		o.urlParams.Set("amgVideoId", strconv.FormatInt(amgVideoId, 10))
	}
}

func UPC(upc int64) LookupOption {
	return func(o *requestOption) {
		o.urlParams.Set("upc", strconv.FormatInt(upc, 10))
	}
}

func ISBN(isbn int64) LookupOption {
	return func(o *requestOption) {
		o.urlParams.Set("isbn", strconv.FormatInt(isbn, 10))
	}
}
