package itunes

type MediaType string

const (
	Movie      MediaType = "movie"      // movieArtist, movie
	Podcast              = "podcast"    // podcastAuthor, podcast
	Music                = "music"      // musicArtist, musicTrack, album, musicVideo, mix, song.
	MusicVideo           = "musicVideo" // musicArtist, musicVideo
	AudioBook            = "audiobook"  // audiobookAuthor, audiobook
	ShortFilm            = "shortFilm"  // shortFilmArtist, shortFilm
	TvShow               = "tvShow"     // tvEpisode, tvSeason
	Software             = "software"   // software, iPadSoftware, macSoftware
	Ebook                = "ebook"      // ebook
	All                  = "all"        // movie, album, allArtist, podcast, musicVideo, mix, audiobook, tvSeason, allTrack
)
