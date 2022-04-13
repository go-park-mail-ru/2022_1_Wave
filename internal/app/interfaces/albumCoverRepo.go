package utilsInterfaces

import "github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"

type AlbumCoverRepoInterface interface {
	Insert(domain.AlbumCover) error
	Update(domain.AlbumCover) error
	Delete(int) error
	SelectByID(int) (*domain.AlbumCover, error)
	GetAll() ([]domain.AlbumCover, error)
	GetLastId() (id int, err error)
	//GetType() reflect.Type
	GetSize() (int, error)
}
