package entity

import "testing"

func TestToString(t *testing.T) {
	cases := map[string]struct {
		in   EntityType
		want string
	}{
		"MovieArtist":     {MovieArtist, "movieArtist"},
		"Movie":           {Movie, "movie"},
		"PodcastAuthor":   {PodcastAuthor, "podcastAuthor"},
		"Podcast":         {Podcast, "podcast"},
		"MusicArtist":     {MusicArtist, "musicArtist"},
		"MusicTrack":      {MusicTrack, "musicTrack"},
		"Mix":             {Mix, "mix"},
		"Song":            {Song, "song"},
		"AudiobookAuthor": {AudiobookAuthor, "audiobookAuthor"},
		"Audiobook":       {Audiobook, "audiobook"},
		"ShortFilmArtist": {ShortFilmArtist, "shortFilmArtist"},
		"ShortFilm":       {ShortFilm, "shortFilm"},
		"TvEpisode":       {TvEpisode, "tvEpisode"},
		"TvSeason":        {TvSeason, "tvSeason"},
		"Software":        {Software, "software"},
		"IPadSoftware":    {IPadSoftware, "iPadSoftware"},
		"MacSoftware":     {MacSoftware, "macSoftware"},
		"Ebook":           {Ebook, "ebook"},
		"AllArtist":       {AllArtist, "allArtist"},
		"AllTrack":        {AllTrack, "allTrack"},
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
