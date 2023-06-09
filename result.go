package itunes

import "time"

// iTunesResult represent iTunes response outer most structure
type iTunesResult struct {
	ResultCount int      `json:"resultCount"`
	Results     []Result `json:"results"`
}

type Result struct {
	WrapperType            string    `json:"wrapperType"`
	Kind                   string    `json:"kind"`
	ArtistId               int64     `json:"artistId"`
	CollectionId           int64     `json:"collectionId"`
	TrackId                int64     `json:"trackId"`
	ArtistName             string    `json:"artistName"`
	CollectionName         string    `json:"collectionName"`
	TrackName              string    `json:"trackName"`
	CollectionCensoredName string    `json:"collectionCensoredName"`
	TrackCensoredName      string    `json:"trackCensoredName"`
	ArtistViewUrl          string    `json:"artistViewUrl"`
	CollectionViewUrl      string    `json:"collectionViewUrl"`
	TrackViewUrl           string    `json:"trackViewUrl"`
	PreviewUrl             string    `json:"previewUrl"`
	ArtworkUrl30           string    `json:"artworkUrl30"`
	ArtworkUrl60           string    `json:"artworkUrl60"`
	ArtworkUrl100          string    `json:"artworkUrl100"`
	CollectionPrice        float64   `json:"collectionPrice"`
	TrackPrice             float64   `json:"trackPrice"`
	ReleaseDate            time.Time `json:"releaseDate"`
	CollectionExplicitness string    `json:"collectionExplicitness"`
	TrackExplicitness      string    `json:"trackExplicitness"`
	DiscCount              int64     `json:"discCount"`
	DiscNumber             int64     `json:"discNumber"`
	TrackCount             int64     `json:"trackCount"`
	TrackNumber            int64     `json:"trackNumber"`
	TrackTimeMillis        int64     `json:"trackTimeMillis"`
	Country                string    `json:"country"`
	Currency               string    `json:"currency"`
	PrimaryGenreName       string    `json:"primaryGenreName"`
	IsStreamable           bool      `json:"isStreamable"`
}
