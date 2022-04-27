package ArtistUseCase

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal"
	AlbumUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/app/album/usecase"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	TrackUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/app/track/usecase"
)

//var UseCase structsUseCase.UseCase

type ArtistUseCase struct {
	ArtistRepo *domain.ArtistRepo
}

type ArtistUseCaseInterface interface {
	CastToDTO(artist domain.Artist, album AlbumUseCase.AlbumUseCaseInterface, track TrackUseCase.TrackUseCaseInterface) (*domain.ArtistDataTransfer, error)
	GetAll(album AlbumUseCase.AlbumUseCaseInterface, track TrackUseCase.TrackUseCaseInterface) ([]domain.ArtistDataTransfer, error)
	GetLastId() (id int, err error)
	Create(dom domain.Artist) error
	Update(dom domain.Artist) error
	Delete(id int) error
	GetById(track TrackUseCase.TrackUseCaseInterface, album AlbumUseCase.AlbumUseCaseInterface, id int) (*domain.ArtistDataTransfer, error)
	GetPopular(track TrackUseCase.TrackUseCaseInterface, album AlbumUseCase.AlbumUseCaseInterface) ([]domain.ArtistDataTransfer, error)
	GetSize() (int, error)
}

func (useCase ArtistUseCase) CastToDTO(artist domain.Artist, album AlbumUseCase.AlbumUseCaseInterface, track TrackUseCase.TrackUseCaseInterface) (*domain.ArtistDataTransfer, error) {
	coverPath, err := artist.CreatePath(internal.PngFormat)
	if err != nil {
		return nil, err
	}

	albumsDto, err := album.GetAlbumsFromArtist(artist.Id, track)

	return &domain.ArtistDataTransfer{
		Id:     artist.Id,
		Name:   artist.Name,
		Cover:  coverPath,
		Likes:  artist.CountLikes,
		Albums: albumsDto,
	}, nil
}

func MakeArtistUseCase(repo domain.ArtistRepo) ArtistUseCase {
	return ArtistUseCase{ArtistRepo: &repo}
}

func (useCase ArtistUseCase) GetAll(album AlbumUseCase.AlbumUseCaseInterface, track TrackUseCase.TrackUseCaseInterface) ([]domain.ArtistDataTransfer, error) {
	artists, err := (*useCase.ArtistRepo).GetAll()
	if err != nil {
		return nil, err
	}

	dto := make([]domain.ArtistDataTransfer, len(artists))

	for idx, obj := range artists {
		data, err := useCase.CastToDTO(obj, album, track)
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

func (useCase ArtistUseCase) GetById(track TrackUseCase.TrackUseCaseInterface, album AlbumUseCase.AlbumUseCaseInterface, id int) (*domain.ArtistDataTransfer, error) {
	artist, err := (*useCase.ArtistRepo).SelectByID(id)
	if err != nil {
		return nil, err
	}
	dto, err := useCase.CastToDTO(*artist, album, track)
	if err != nil {
		return nil, err
	}

	return dto, nil
}

func (useCase ArtistUseCase) GetPopular(track TrackUseCase.TrackUseCaseInterface, album AlbumUseCase.AlbumUseCaseInterface) ([]domain.ArtistDataTransfer, error) {
	artists, err := (*useCase.ArtistRepo).GetPopular()
	if err != nil {
		return nil, err
	}

	dto := make([]domain.ArtistDataTransfer, len(artists))

	for idx, obj := range artists {
		data, err := useCase.CastToDTO(obj, album, track)
		if err != nil {
			return nil, err
		}
		dto[idx] = *data
	}

	return dto, nil
}

//func (useCase ArtistUseCase) GetType() reflect.Type {
//	return reflect.TypeOf(domain.Artist{})
//}

func (useCase ArtistUseCase) GetSize() (int, error) {
	return (*useCase.ArtistRepo).GetSize()
}
