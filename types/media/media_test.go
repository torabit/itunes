package media

import "testing"

func TestToString(t *testing.T) {
	cases := map[string]struct {
		in   MediaType
		want string
	}{
		"Movie":      {Movie, "movie"},
		"Podcast":    {Podcast, "podcast"},
		"Music":      {Music, "music"},
		"MusicVideo": {MusicVideo, "musicVideo"},
		"AudioBook":  {AudioBook, "audioBook"},
		"ShortFilm":  {ShortFilm, "shortfilm"},
		"TvShow":     {TvShow, "tvShow"},
		"Software":   {Software, "software"},
		"Ebook":      {Ebook, "ebook"},
		"All":        {All, "all"},
	}

	for name, tt := range cases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := ToString(tt.in)
			if got != tt.want {
				t.Fatalf("want = %+v, but got = %+v", tt.want, got)
			}
		})
	}
}
