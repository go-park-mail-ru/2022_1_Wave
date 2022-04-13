package utilsInterfaces

import (
	"reflect"
)

type UseCaseInterface interface {
	GetAll() ([]Domain, error)
	GetLastId() (id uint64, err error)
	Create(Domain) (UseCaseInterface, error)
	Update(Domain) (UseCaseInterface, error)
	Delete(id uint64) (UseCaseInterface, error)
	GetById(id uint64) (Domain, error)
	GetPopular() ([]Domain, error)
	GetType() reflect.Type
	SetRepo(repoInterface RepoInterface) (UseCaseInterface, error)
	GetRepo() (RepoInterface, error)
	GetSize() (uint64, error)

	// -------------------------
	//todo пока кастыль
	GetTracksFromAlbum(albumId uint64) (interface{}, error)
	//todo пока кастыль
	GetPopularTracksFromArtist(artistId uint64) (interface{}, error)
}
