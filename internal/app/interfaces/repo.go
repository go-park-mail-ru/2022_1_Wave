package utilsInterfaces

import (
	"reflect"
)

type RepoInterface interface {
	Insert(Domain) (RepoInterface, error)
	Update(Domain) (RepoInterface, error)
	Delete(uint64) (RepoInterface, error)
	SelectByID(uint64) (Domain, error)
	GetAll() ([]Domain, error)
	GetPopular() ([]Domain, error)
	GetLastId() (id uint64, err error)
	GetType() reflect.Type
	GetSize() (uint64, error)

	//todo пока кастыль
	GetTracksFromAlbum(albumId uint64) (interface{}, error)
	//todo пока кастыль
	GetAlbumsFromArtist(artist uint64) (interface{}, error)
	//todo пока кастыль
	GetPopularTracksFromArtist(artistId uint64) (interface{}, error)
}
