package ArtistUseCase

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal"
	AlbumUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/app/album/usecase"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	utilsInterfaces "github.com/go-park-mail-ru/2022_1_Wave/internal/app/interfaces"
	trackUseCase "github.com/go-park-mail-ru/2022_1_Wave/internal/app/track/usecase"
	"reflect"
)

//var UseCase structsUseCase.UseCase

type ArtistUseCase struct {
	ArtistRepo *utilsInterfaces.ArtistRepoInterface
}

func (useCase ArtistUseCase) CastToDTO(artist domain.Artist, album AlbumUseCase.AlbumUseCase, track trackUseCase.TrackUseCase) (*domain.ArtistDataTransfer, error) {
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

func MakeArtistUseCase(repo utilsInterfaces.ArtistRepoInterface) ArtistUseCase {
	return ArtistUseCase{ArtistRepo: &repo}
}

func (useCase ArtistUseCase) GetAll(album AlbumUseCase.AlbumUseCase, track trackUseCase.TrackUseCase) ([]domain.ArtistDataTransfer, error) {
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

func (useCase ArtistUseCase) GetById(track trackUseCase.TrackUseCase, album AlbumUseCase.AlbumUseCase, id int) (*domain.ArtistDataTransfer, error) {
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

func (useCase ArtistUseCase) GetPopular(track trackUseCase.TrackUseCase, album AlbumUseCase.AlbumUseCase) ([]domain.ArtistDataTransfer, error) {
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

func (useCase ArtistUseCase) GetType() reflect.Type {
	return reflect.TypeOf(domain.Artist{})
}

func (useCase ArtistUseCase) GetSize() (int, error) {
	return (*useCase.ArtistRepo).GetSize()
}
