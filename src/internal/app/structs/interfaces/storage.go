package utilsInterfaces

type GlobalStorageInterface interface {
	Open() error
	Init(quantity int) (GlobalStorageInterface, error)
	Close() error
	GetAlbumRepo() *RepoInterface
	GetArtistRepo() *RepoInterface
	GetTrackRepo() *RepoInterface
}
