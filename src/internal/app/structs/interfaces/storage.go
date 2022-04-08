package utilsInterfaces

type GlobalStorageInterface interface {
	Open() (GlobalStorageInterface, error)
	Init(quantity int) (GlobalStorageInterface, error)
	Close() error
	GetAlbumRepo() *RepoInterface
	GetAlbumCoverRepo() *RepoInterface
	GetArtistRepo() *RepoInterface
	GetTrackRepo() *RepoInterface
	GetAlbumRepoLen() (int, error)
	GetAlbumCoverRepoLen() (int, error)
	GetArtistRepoLen() (int, error)
	GetTrackRepoLen() (int, error)
}
