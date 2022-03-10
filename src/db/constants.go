package db

// errors
const (
	InvalidBody     = "invalid body format"
	InvalidJson     = "error to unpacking json"
	IndexOutOfRange = "index out of range"

	AlbumIsNotExist  = "album is not exist"
	TrackIsNotExist  = "track is not exist"
	ArtistIsNotExist = "artist is not exist"
)

// success albums
const (
	SuccessCreatedAlbum = "success created album"
	SuccessUpdatedAlbum = "success updated album"
	SuccessDeletedAlbum = "success deleted album"
)

// success tracks
const (
	SuccessCreatedTrack = "success created track"
	SuccessUpdatedTrack = "success updated track"
	SuccessDeletedTrack = "success deleted track"
)

// success artists
const (
	SuccessCreatedArtist = "success created artist"
	SuccessUpdatedArtist = "success updated artist"
	SuccessDeletedArtist = "success deleted artist"
)
