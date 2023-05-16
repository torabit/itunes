package itunes

import (
	"testing"
)

func TestLookupOptions(t *testing.T) {
	cases := map[string]struct {
		input    LookupOption
		expected string
	}{
		"ID":          {ID(909253), "id=909253"},
		"AMGArtistID": {AMGArtistID(468749), "amgArtistId=468749"},
		"AMGAlbumID":  {AMGAlbumID(15175), "amgAlbumId=15175"},
		"AMGVideoID":  {AMGVideoID(17120), "amgVideoId=17120"},
		"UPC":         {UPC(720642462928), "upc=720642462928"},
		"ISBN":        {ISBN(9780316069359), "isbn=9780316069359"},
	}
	for name, tt := range cases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			v := processOptions()
			tt.input(&v)

			actual := v.urlParams.Encode()

			if actual != tt.expected {
				t.Errorf("processLookupOption(%v) = %v; want %v", tt.input, actual, tt.expected)
			}
		})
	}
}
