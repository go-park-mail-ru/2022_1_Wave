package domain

import (
	auth_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/auth"
	user_microservice_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/user"
)

type GlobalStorageInterface interface {
	Open() (GlobalStorageInterface, error)
	Init(quantity int64) (GlobalStorageInterface, error)
	Close() error

	GetAlbumRepo() AlbumRepo
	GetAlbumCoverRepo() AlbumCoverRepo
	GetArtistRepo() ArtistRepo
	GetTrackRepo() TrackRepo
	GetSessionRepo() auth_domain.AuthRepo
	GetUserRepo() user_microservice_domain.UserRepo
}
