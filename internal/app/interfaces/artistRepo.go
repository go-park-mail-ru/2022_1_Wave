package utilsInterfaces

import "github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"

type ArtistRepoInterface interface {
	Insert(domain.Artist) error
	Update(domain.Artist) error
	Delete(int) error
	SelectByID(int) (*domain.Artist, error)
	GetAll() ([]domain.Artist, error)
	GetPopular() ([]domain.Artist, error)
	GetLastId() (id int, err error)
	//GetType() reflect.Type
	GetSize() (int, error)
}
