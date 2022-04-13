package utilsInterfaces

import (
	"reflect"
	"sync"
)

type RepoInterface interface {
	Insert(Domain, *sync.RWMutex) (RepoInterface, error)
	Update(Domain, *sync.RWMutex) (RepoInterface, error)
	Delete(uint64, *sync.RWMutex) (RepoInterface, error)
	SelectByID(uint64, *sync.RWMutex) (Domain, error)
	GetAll(*sync.RWMutex) ([]Domain, error)
	GetPopular(*sync.RWMutex) ([]Domain, error)
	GetLastId(*sync.RWMutex) (id uint64, err error)
	GetType(*sync.RWMutex) reflect.Type
	GetSize(mutex *sync.RWMutex) (uint64, error)

	//todo пока кастыль
	GetTracksFromAlbum(albumId uint64, mutex *sync.RWMutex) (interface{}, error)
	//todo пока кастыль
	GetAlbumsFromArtist(artist uint64, mutex *sync.RWMutex) (interface{}, error)
	//todo пока кастыль
	GetPopularTracksFromArtist(artistId uint64, mutex *sync.RWMutex) (interface{}, error)
}
