package itunes

import "testing"

func TestOptions(t *testing.T) {
	t.Parallel()
	resultSet := processOptions(
		Term("example_song"),
		Country(CountryUS),
		Media(MediaMusic),
		Entity(EntitySong),
		Attribute(AttributeSongTerm),
		Limit(200),
		Lang(LangEnglish),
		Version(2),
		Explicit(true),
	)

	expected := "attribute=songTerm&country=US&entity=song&explicit=Yes&lang=en_us&limit=200&media=music&term=example_song&version=2"
	actual := resultSet.urlParams.Encode()
	if actual != expected {
		t.Errorf("Expected: %+v, got: %+v", expected, actual)
	}
}
