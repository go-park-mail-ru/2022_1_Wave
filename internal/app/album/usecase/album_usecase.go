package AlbumUseCase

import (
	AlbumPostgres "github.com/go-park-mail-ru/2022_1_Wave/internal/app/album/repository"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
)

type AlbumUseCase struct {
	TrackRepo      *domain.TrackRepo
	ArtistRepo     *domain.ArtistRepo
	AlbumRepo      *domain.AlbumRepo
	AlbumCoverRepo *domain.AlbumCoverRepo
}

type AlbumUseCaseInterface interface {
	CastToDTO(album domain.Album) (*domain.AlbumDataTransfer, error)
	GetAll() ([]domain.AlbumDataTransfer, error)
	GetLastId() (id int64, err error)
	Create(dom domain.Album) error
	Update(dom domain.Album) error
	Delete(id int64) error
	GetById(id int64) (*domain.AlbumDataTransfer, error)
	GetPopular() ([]domain.AlbumDataTransfer, error)
	GetAlbumsFromArtist(artist int64) ([]domain.AlbumDataTransfer, error)
	GetSize() (int64, error)
}

func MakeAlbumUseCase(track domain.TrackRepo,
	artist domain.ArtistRepo,
	album domain.AlbumRepo,
	albumCover domain.AlbumCoverRepo) AlbumUseCase {
	return AlbumUseCase{
		TrackRepo:      &track,
		ArtistRepo:     &artist,
		AlbumRepo:      &album,
		AlbumCoverRepo: &albumCover,
	}
}

func (useCase AlbumUseCase) CastToDTO(album domain.Album) (*domain.AlbumDataTransfer, error) {
	artist, err := (*useCase.ArtistRepo).SelectByID(album.ArtistId)
	if err != nil {
		return nil, err
	}
	return AlbumPostgres.GetFullAlbumByArtist(*useCase.TrackRepo, album, *artist)
}

func (useCase AlbumUseCase) GetAll() ([]domain.AlbumDataTransfer, error) {
	albums, err := (*useCase.AlbumRepo).GetAll()
	if err != nil {
		return nil, err
	}

	dto := make([]domain.AlbumDataTransfer, len(albums))

	for idx, obj := range albums {
		data, err := useCase.CastToDTO(obj)
		if err != nil {
			return nil, err
		}
		dto[idx] = *data
	}

	return dto, nil
}

func (useCase AlbumUseCase) GetLastId() (id int64, err error) {
	return (*useCase.AlbumRepo).GetLastId()
}

func (useCase AlbumUseCase) Create(dom domain.Album) error {
	return (*useCase.AlbumRepo).Insert(dom)
}

func (useCase AlbumUseCase) Update(dom domain.Album) error {
	return (*useCase.AlbumRepo).Update(dom)
}

func (useCase AlbumUseCase) Delete(id int64) error {
	return (*useCase.AlbumRepo).Delete(id)
}

func (useCase AlbumUseCase) GetById(id int64) (*domain.AlbumDataTransfer, error) {
	album, err := (*useCase.AlbumRepo).SelectByID(id)
	if err != nil {
		return nil, err
	}
	dto, err := useCase.CastToDTO(*album)
	if err != nil {
		return nil, err
	}

	return dto, nil
}

func (useCase AlbumUseCase) GetPopular() ([]domain.AlbumDataTransfer, error) {
	albums, err := (*useCase.AlbumRepo).GetPopular()
	if err != nil {
		return nil, err
	}

	dto := make([]domain.AlbumDataTransfer, len(albums))

	for idx, obj := range albums {
		data, err := useCase.CastToDTO(obj)
		if err != nil {
			return nil, err
		}
		dto[idx] = *data
	}

	return dto, nil
}

func (useCase AlbumUseCase) GetAlbumsFromArtist(artist int64) ([]domain.AlbumDataTransfer, error) {
	albums, err := (*useCase.AlbumRepo).GetAlbumsFromArtist(artist)
	if err != nil {
		return nil, err
	}

	dto := make([]domain.AlbumDataTransfer, len(albums))

	for idx, obj := range albums {
		data, err := useCase.CastToDTO(obj)
		if err != nil {
			return nil, err
		}
		dto[idx] = *data
	}

	return dto, nil
}

func (useCase AlbumUseCase) GetSize() (int64, error) {
	return (*useCase.AlbumRepo).GetSize()
}
