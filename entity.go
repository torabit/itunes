package itunes

type EntityType string

const (
	EntityMovieArtist     EntityType = "movieArtist"
	EntityMovie                      = "movie"
	EntityPodcastAuthor              = "podcastAuthor"
	EntityPodcast                    = "podcast"
	EntityMusicArtist                = "musicArtist"
	EntityMusicTrack                 = "musicTrack"
	EntityAlbum                      = "album"
	EntityMusicVideo                 = "musicVideo"
	EntityMix                        = "mix"
	EntitySong                       = "song"
	EntityAudiobookAuthor            = "audiobookAuthor"
	EntityAudiobook                  = "audiobook"
	EntityShortFilmArtist            = "shortFilmArtist"
	EntityShortFilm                  = "shortFilm"
	EntityTvEpisode                  = "tvEpisode"
	EntityTvSeason                   = "tvSeason"
	EntitySoftware                   = "software"
	EntityIPadSoftware               = "iPadSoftware"
	EntityMacSoftware                = "macSoftware"
	EntityEbook                      = "ebook"
	EntityAllArtist                  = "allArtist"
	EntityAllTrack                   = "allTrack"
)
