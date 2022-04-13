package utilsInterfaces

import "github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"

type TrackRepoInterface interface {
	Insert(domain.Track) error
	Update(domain.Track) error
	Delete(int) error
	SelectByID(int) (*domain.Track, error)
	GetAll() ([]domain.Track, error)
	GetPopular() ([]domain.Track, error)
	GetLastId() (id int, err error)
	//GetType() reflect.Type
	GetSize() (int, error)
	GetTracksFromAlbum(albumId int) ([]domain.Track, error)
	GetPopularTracksFromArtist(artistId int) ([]domain.Track, error)
}
