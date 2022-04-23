package ArtistUseCase

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal"
	AlbumPostgres "github.com/go-park-mail-ru/2022_1_Wave/internal/app/album/repository"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
)

//var UseCase structsUseCase.UseCase

type ArtistUseCase struct {
	ArtistRepo *domain.ArtistRepo
	AlbumRepo  *domain.AlbumRepo
	TrackRepo  *domain.TrackRepo
}

type ArtistUseCaseInterface interface {
	CastToDTO(artist domain.Artist) (*domain.ArtistDataTransfer, error)
	GetAll() ([]domain.ArtistDataTransfer, error)
	GetLastId() (id int, err error)
	Create(dom domain.Artist) error
	Update(dom domain.Artist) error
	Delete(id int) error
	GetById(id int) (*domain.ArtistDataTransfer, error)
	GetPopular() ([]domain.ArtistDataTransfer, error)
	GetSize() (int, error)
}

func (useCase ArtistUseCase) CastToDTO(artist domain.Artist) (*domain.ArtistDataTransfer, error) {
	coverPath, err := artist.CreatePath(internal.PngFormat)
	if err != nil {
		return nil, err
	}

	albums, err := (*useCase.AlbumRepo).GetAlbumsFromArtist(artist.Id)
	if err != nil {
		return nil, err
	}

	albumsDto := make([]domain.AlbumDataTransfer, len(albums))

	for idx, album := range albums {
		albumDto, err := AlbumPostgres.GetFullAlbumByArtist(*useCase.TrackRepo, album, artist)
		if err != nil {
			return nil, err
		}
		albumsDto[idx] = *albumDto
	}

	return &domain.ArtistDataTransfer{
		Id:     artist.Id,
		Name:   artist.Name,
		Cover:  coverPath,
		Likes:  artist.CountLikes,
		Albums: albumsDto,
	}, nil
}

func MakeArtistUseCase(artistRepo domain.ArtistRepo, albumRepo domain.AlbumRepo, trackRepo domain.TrackRepo) ArtistUseCase {
	return ArtistUseCase{
		ArtistRepo: &artistRepo,
		AlbumRepo:  &albumRepo,
		TrackRepo:  &trackRepo}
}

func (useCase ArtistUseCase) GetAll() ([]domain.ArtistDataTransfer, error) {
	artists, err := (*useCase.ArtistRepo).GetAll()
	if err != nil {
		return nil, err
	}

	dto := make([]domain.ArtistDataTransfer, len(artists))

	for idx, obj := range artists {
		data, err := useCase.CastToDTO(obj)
		if err != nil {
			return nil, err
		}
		dto[idx] = *data
	}

	return dto, nil
}

func (useCase ArtistUseCase) GetLastId() (id int, err error) {
	return (*useCase.ArtistRepo).GetLastId()
}

func (useCase ArtistUseCase) Create(dom domain.Artist) error {
	return (*useCase.ArtistRepo).Insert(dom)
}

func (useCase ArtistUseCase) Update(dom domain.Artist) error {
	return (*useCase.ArtistRepo).Update(dom)
}

func (useCase ArtistUseCase) Delete(id int) error {
	return (*useCase.ArtistRepo).Delete(id)
}

func (useCase ArtistUseCase) GetById(id int) (*domain.ArtistDataTransfer, error) {
	artist, err := (*useCase.ArtistRepo).SelectByID(id)
	if err != nil {
		return nil, err
	}
	dto, err := useCase.CastToDTO(*artist)
	if err != nil {
		return nil, err
	}

	return dto, nil
}

func (useCase ArtistUseCase) GetPopular() ([]domain.ArtistDataTransfer, error) {
	artists, err := (*useCase.ArtistRepo).GetPopular()
	if err != nil {
		return nil, err
	}

	dto := make([]domain.ArtistDataTransfer, len(artists))

	for idx, obj := range artists {
		data, err := useCase.CastToDTO(obj)
		if err != nil {
			return nil, err
		}
		dto[idx] = *data
	}

	return dto, nil
}

func (useCase ArtistUseCase) GetSize() (int, error) {
	return (*useCase.ArtistRepo).GetSize()
}
