package internal

// errors
const (
	InvalidBody     = "invalid body format"
	InvalidJson     = "error to unpacking json"
	IndexOutOfRange = "index out of range"

	AlbumIsNotExist  = "album is not exist"
	TrackIsNotExist  = "track is not exist"
	ArtistIsNotExist = "artist is not exist"

	IsNotExist = "is not exist"
)

// success
const (
	SuccessCreated = "success created"
	SuccessUpdated = "success updated"
	SuccessDeleted = "success deleted"
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

// track constraints
const (
	TrackTitleLen   = varchar255
	TrackLinkLen    = varchar255
	TrackMp4LinkLen = varchar255
)

// artist constraints
const (
	ArtistNameLen      = varchar255
	ArtistPhotoLinkLen = varchar255
)

// album errors
const (
	ErrorAlbumIdIsNegative             = "album's id is negative"
	ErrorAlbumMaxTitleLen              = "album's title length is over than max"
	ErrorCoverIdIsNegative             = "album's cover id is negative"
	ErrorAlbumAuthorIdIsNegative       = "author's id of album is negative"
	ErrorAlbumCountLikesIsNegative     = "number of album's count likes is negative"
	ErrorAlbumCountListeningIsNegative = "number of album's listening is negative"
	ErrorAlbumCoverIdIsNegative        = "cover's id is negative"
)

// tracks errors
const (
	ErrorTrackIdIsNegative             = "track's id is negative"
	ErrorTrackMaxTitleLen              = "track's title length is over than max"
	ErrorTrackArtistIdIsNegative       = "author's id of track is negative"
	ErrorTrackCoverIdIsNegative        = "cover's id of track is negative"
	ErrorTrackMaxPhotoLinkLen          = "length of song's link to source is over than max"
	ErrorTrackMp4MaxLinkLen            = "length of Mp4's link to source is over than max"
	ErrorTrackCountLikesIsNegative     = "number of track's count likes is negative"
	ErrorTrackCountListeningIsNegative = "number of track's listening is negative"
)

// artists errors
const (
	ErrorArtistIdIsNegative             = "artist's id is negative"
	ErrorArtistMaxNameLen               = "authors's name length is over than max"
	ErrorArtistsMaxPhotoLinkLen         = "length of authors's link to photo is over than max"
	ErrorArtistCountFollowersIsNegative = "number of count followers is negative"
	ErrorArtistCountLikesIsNegative     = "number of artist's count likes is negative"
	ErrorArtistCountListeningIsNegative = "number of artist's listening is negative"
)

// db errors
const (
	ErrorDbIsEmpty = "database is empty"
)

// local db errors
const (
	ErrorLocalDbArtistsNotEnought = "sum of artists not equal to input quantity"
	ErrorLocalDbAlbumsNotEnought  = "sum of albums not equal to input quantity"
	ErrorLocalDbTracksNotEnought  = "sum of tracks not equal to input quantity"
)

// album fields
const (
	FieldId             = "id"
	FieldTitle          = "title"
	FieldArtist         = "artist"
	FieldAuthorId       = "authorId"
	FieldAlbumId        = "albumId"
	FieldCover          = "cover"
	FieldCountLikes     = "countLikes"
	FieldCountListening = "countListening"
	FieldDate           = "date"
	FieldCoverId        = "coverId"
	FieldName           = "name"
	FieldPhoto          = "photo"
	FieldCountFollowers = "countFollowers"
	FieldDuration       = "duration"
	FieldMp4            = "mp4"
)

// others
const (
	AssetsPrefix  = "assets/"
	AlbumPreName  = "album_"
	ArtistPreName = "artist_"
	TrackPreName  = "track_"
	PngFormat     = ".png"
	NullId        = 0
	BadId         = -1
	BadType       = "bad type"
)