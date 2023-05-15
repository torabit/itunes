package itunes

import (
	"context"
	"net/http"
	"testing"
)

func TestLookupID(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "test_data/lookup_id.json")
	defer server.Close()

	result, err := client.Lookup(context.Background(), ID(1611115294))
	expectedCollectionName := "NIA"
	expectedArtistName := "Kaho Nakamura"
	if err != nil {
		t.Error(err)
	}
	if result.ResultCount != 1 {
		t.Errorf("got: %v; want %v", result.ResultCount, 1)
	}
	if result.Results[0].CollectionName != expectedCollectionName {
		t.Errorf("got: %v; want %v", result.Results[0].CollectionName, expectedCollectionName)
	}
	if result.Results[0].ArtistName != expectedArtistName {
		t.Errorf("got: %v; want %v", result.Results[0].CollectionName, expectedArtistName)
	}
}
