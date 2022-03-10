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

// db constants
const (
	varchar512 = 512
	varchar255 = 255
)

// album constraints
const (
	AlbumTitleLen = varchar255
)

// song constraints
const (
	SongTitleLen = varchar255
	SongLinkLen  = varchar255
)

// artist constraints
const (
	ArtistNameLen      = varchar255
	ArtistPhotoLinkLen = varchar255
)

// album errors
const (
	ErrorAlbumIdIsNegative        = "album's id is negative"
	ErrorAlbumMaxTitleLen         = "album's title length is over than max"
	ErrorAuthorIdIsNegative       = "author's id is negative"
	ErrorCountLikesIsNegative     = "number of count likes is negative"
	ErrorCountListeningIsNegative = "number of listening is negative"
	ErrorCoverIdIsNegative        = "cover's id is negative"
)

// songs errors
const (
	ErrorSongMaxNameLen      = "song's title length is over than max"
	ErrorSongMaxPhotoLinkLen = "length of song's link to source is over than max"
)

// artists errors
const (
	ErrorArtistMaxNameLen       = "authors's name length is over than max"
	ErrorArtistsMaxPhotoLinkLen = "length of authors's link to photo is over than max"
)
