package utilsInterfaces

import "github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"

type GlobalStorageInterface interface {
	Open() (GlobalStorageInterface, error)
	Init(quantity int) (GlobalStorageInterface, error)
	Close() error

	GetAlbumRepo() domain.AlbumRepo
	GetAlbumCoverRepo() domain.AlbumCoverRepo
	GetArtistRepo() domain.ArtistRepo
	GetTrackRepo() domain.TrackRepo
	GetSessionRepo() domain.SessionRepo
	GetUserRepo() domain.UserRepo
}
