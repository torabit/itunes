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

func ToString(mediaType MediaType) string {

	stringType := ""

	switch mediaType {
	case Movie:
		stringType = "movie"
	case Podcast:
		stringType = "podcast"
	case Music:
		stringType = "music"
	case MusicVideo:
		stringType = "musicVideo"
	case AudioBook:
		stringType = "audioBook"
	case ShortFilm:
		stringType = "shortFilm"
	case TvShow:
		stringType = "tvShow"
	case Software:
		stringType = "software"
	case Ebook:
		stringType = "ebook"
	case All:
		stringType = "all"
	}

	return stringType
}
