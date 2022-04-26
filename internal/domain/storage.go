package domain

type GlobalStorageInterface interface {
	Open() (GlobalStorageInterface, error)
	Init(quantity int64) (GlobalStorageInterface, error)
	Close() error

	GetAlbumRepo() AlbumRepo
	GetAlbumCoverRepo() AlbumCoverRepo
	GetArtistRepo() ArtistRepo
	GetTrackRepo() TrackRepo
	GetSessionRepo() SessionRepo
	GetUserRepo() UserRepo
}
