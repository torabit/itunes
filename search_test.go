package itunes

import (
	"context"
	"net/http"
	"testing"
)

func TestSearch(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "test_data/search.json")
	defer server.Close()

	result, err := client.Search(
		context.Background(),
		Term("Helsinki Lambda Club"),
		Media(MediaMusic),
		Attribute(AttributeArtistTerm),
		Country(CountryJP),
		Entity(EntityMusicTrack),
		Limit(2),
	)

	if err != nil {
		t.Error(err)
	}

	if result.ResultCount != 2 {
		t.Fatalf("got: %v; want %v", result.ResultCount, 2)
	}

	cases := map[string]struct {
		result   string
		expected string
	}{
		"artist name_0": {result.Results[0].ArtistName, "Helsinki Lambda Club"},
		"track name_0":  {result.Results[0].TrackName, "眠ったふりして"},
		"artist name_1": {result.Results[1].ArtistName, "Helsinki Lambda Club"},
		"track name_1":  {result.Results[1].TrackName, "PIZZASHAKE"},
	}

	for name, tt := range cases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			if tt.result != tt.expected {
				t.Fatalf("want %v; got %v", tt.expected, tt.result)
			}
		})
	}

}
