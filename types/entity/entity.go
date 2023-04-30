package entity

type EntityType string

const (
	MovieArtist     EntityType = "movieArtist"
	Movie                      = "movie"
	PodcastAuthor              = "podcastAuthor"
	Podcast                    = "podcast"
	MusicArtist                = "musicArtist"
	MusicTrack                 = "musicTrack"
	Album                      = "album"
	MusicVideo                 = "musicVideo"
	Mix                        = "mix"
	Song                       = "song"
	AudiobookAuthor            = "audiobookAuthor"
	Audiobook                  = "audiobook"
	ShortFilmArtist            = "shortFilmArtist"
	ShortFilm                  = "shortFilm"
	TvEpisode                  = "tvEpisode"
	TvSeason                   = "tvSeason"
	Software                   = "software"
	IPadSoftware               = "iPadSoftware"
	MacSoftware                = "macSoftware"
	Ebook                      = "ebook"
	AllArtist                  = "allArtist"
	AllTrack                   = "allTrack"
)

func ToString(entityType EntityType) string {
	value := ""

	switch entityType {
	case MovieArtist:
		value = "movieArtist"
	case Movie:
		value = "movie"
	case PodcastAuthor:
		value = "podcastAuthor"
	case Podcast:
		value = "podcast"
	case MusicArtist:
		value = "musicArtist"
	case MusicTrack:
		value = "musicTrack"
	case Mix:
		value = "mix"
	case Song:
		value = "song"
	case AudiobookAuthor:
		value = "audiobookAuthor"
	case Audiobook:
		value = "audiobook"
	case ShortFilmArtist:
		value = "shortFilmArtist"
	case ShortFilm:
		value = "shortFilm"
	case TvEpisode:
		value = "tvEpisode"
	case TvSeason:
		value = "tvSeason"
	case Software:
		value = "software"
	case IPadSoftware:
		value = "iPadSoftware"
	case MacSoftware:
		value = "macSoftware"
	case Ebook:
		value = "ebook"
	case AllArtist:
		value = "allArtist"
	case AllTrack:
		value = "allTrack"
	}

	return value
}
