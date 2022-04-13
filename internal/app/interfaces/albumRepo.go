package utilsInterfaces

import "github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"

type AlbumRepoInterface interface {
	Insert(domain.Album) error
	Update(domain.Album) error
	Delete(int) error
	SelectByID(int) (*domain.Album, error)
	GetAll() ([]domain.Album, error)
	GetPopular() ([]domain.Album, error)
	GetLastId() (id int, err error)
	//GetType() reflect.Type
	GetSize() (int, error)

	//todo пока кастыль
	GetAlbumsFromArtist(artist int) ([]domain.Album, error)
}
