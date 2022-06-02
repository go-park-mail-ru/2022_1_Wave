package internal

// errors
const (
	IndexOutOfRange = "index out of range"
	Unauthorized    = "unauthorized"
)

// success
const (
	SuccessCreated             = "success created"
	SuccessUpdated             = "success updated"
	SuccessDeleted             = "success deleted"
	SuccessAdded               = "success added"
	ToPlaylist                 = "to playlist"
	SuccessRemoved             = "success removed"
	FromPlaylist               = "from playlist"
	SuccessLiked               = "success liked"
	SuccessListened            = "success listened"
	SuccessAddedToFavorites    = "success added to favorites"
	SuccessRemoveFromFavorites = "success remove from favorites"
)

// query
const (
	PlaylistId = "playlistId"
	TrackId    = "trackId"
)

// db types
const (
	Postgres = "postgres"
)

// db errors
const (
	ErrorNothingToUpdate = "nothing to update (check id)"
)

// album fields
const (
	FieldId     = "id"
	FieldToFind = "toFind"
	Hash        = "hash"
)

// others
const (
	AssetsPrefix  = "assets/"
	AlbumPreName  = "album_"
	ArtistPreName = "artist_"
	TrackPreName  = "track_"
	Top           = 20
	SearchTop     = 5
)

// network
const (
	Tcp = "tcp"
)

// auth
const (
	SessionIdKey = "session_id"
)

// clients
const (
	Grpc = "gRPC"
)

// formats
const (
	ImgFormat = ".webp"
	Mp3Format = ".mp3"
)
