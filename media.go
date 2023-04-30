package itunes

type MediaType string

const (
	MediaMovie      MediaType = "movie"      // movieArtist, movie
	MediaPodcast              = "podcast"    // podcastAuthor, podcast
	MediaMusic                = "music"      // musicArtist, musicTrack, album, musicVideo, mix, song.
	MediaMusicVideo           = "musicVideo" // musicArtist, musicVideo
	MediaAudioBook            = "audiobook"  // audiobookAuthor, audiobook
	MediaShortFilm            = "shortFilm"  // shortFilmArtist, shortFilm
	MediaTvShow               = "tvShow"     // tvEpisode, tvSeason
	MediaSoftware             = "software"   // software, iPadSoftware, macSoftware
	MediaEbook                = "ebook"      // ebook
	MediaAll                  = "all"        // movie, album, allArtist, podcast, musicVideo, mix, audiobook, tvSeason, allTrack
)
