package AlbumCoverUseCase

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	"reflect"
)

type AlbumCoverUseCase struct {
	AlbumCoverRepo *domain.AlbumCoverRepo
}

func (useCase AlbumCoverUseCase) CastToDTO(cover domain.AlbumCover) (domain.AlbumCoverDataTransfer, error) {
	return domain.AlbumCoverDataTransfer{
		Quote:  cover.Quote,
		IsDark: cover.IsDark,
	}, nil
}

func MakeAlbumCoverUseCase(repo domain.AlbumCoverRepo) AlbumCoverUseCase {
	return AlbumCoverUseCase{AlbumCoverRepo: &repo}
}

func (useCase AlbumCoverUseCase) GetAll() ([]domain.AlbumCoverDataTransfer, error) {
	albums, err := (*useCase.AlbumCoverRepo).GetAll()
	if err != nil {
		return nil, err
	}

	dto := make([]domain.AlbumCoverDataTransfer, len(albums))

	for idx, obj := range albums {
		data, err := useCase.CastToDTO(obj)
		if err != nil {
			return nil, err
		}
		dto[idx] = data
	}

	return dto, nil
}

func (useCase AlbumCoverUseCase) GetLastId() (id int, err error) {
	return (*useCase.AlbumCoverRepo).GetLastId()
}

func (useCase AlbumCoverUseCase) Create(dom domain.AlbumCover) error {
	return (*useCase.AlbumCoverRepo).Insert(dom)
}

func (useCase AlbumCoverUseCase) Update(dom domain.AlbumCover) error {
	return (*useCase.AlbumCoverRepo).Update(dom)
}

func (useCase AlbumCoverUseCase) Delete(id int) error {
	return (*useCase.AlbumCoverRepo).Delete(id)
}

func (useCase AlbumCoverUseCase) GetById(id int) (*domain.AlbumCoverDataTransfer, error) {
	album, err := (*useCase.AlbumCoverRepo).SelectByID(id)
	if err != nil {
		return nil, err
	}
	dto, err := useCase.CastToDTO(*album)
	if err != nil {
		return nil, err
	}

	return &dto, nil
}

func (useCase AlbumCoverUseCase) GetType() reflect.Type {
	return reflect.TypeOf(domain.AlbumCover{})
}

func (useCase AlbumCoverUseCase) GetSize() (int, error) {
	return (*useCase.AlbumCoverRepo).GetSize()
}
