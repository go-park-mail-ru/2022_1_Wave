package ArtistUseCase

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/album/albumProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/artist/artistProto"
	Gateway "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/gateway"
)

type ArtistUseCase interface {
	GetAll(userId int64) ([]*artistProto.ArtistDataTransfer, error)
	GetLastId() (int64, error)
	Create(transfer *artistProto.Artist) error
	Update(transfer *artistProto.Artist) error
	Delete(int64) error
	GetById(artistId int64, userId int64) (*artistProto.ArtistDataTransfer, error)
	GetPopular(userId int64) ([]*artistProto.ArtistDataTransfer, error)
	GetSize() (int64, error)
	SearchByName(userId int64, name string) ([]*artistProto.ArtistDataTransfer, error)
	GetFavorites(int64) ([]*artistProto.ArtistDataTransfer, error)
	AddToFavorites(userId int64, artistId int64) error
	RemoveFromFavorites(userId int64, artistId int64) error
	Like(arg int64, userId int64) error
	LikeCheckByUser(arg int64, userId int64) (bool, error)
}

type artistUseCase struct {
	albumAgent  domain.AlbumAgent
	artistAgent domain.ArtistAgent
	trackAgent  domain.TrackAgent
}

func NewArtistUseCase(albumAgent domain.AlbumAgent, artistAgent domain.ArtistAgent, trackAgent domain.TrackAgent) *artistUseCase {
	return &artistUseCase{
		albumAgent:  albumAgent,
		artistAgent: artistAgent,
		trackAgent:  trackAgent,
	}
}

func (useCase artistUseCase) CastToDTO(userId int64, artist *artistProto.Artist) (*artistProto.ArtistDataTransfer, error) {
	coverPath, err := Gateway.PathToArtistCover(artist, internal.ImgFormat)
	if err != nil {
		return nil, err
	}

	albums, err := useCase.albumAgent.GetAlbumsFromArtist(artist.Id)
	if err != nil {
		return nil, err
	}

	albumsDto := map[int64]*albumProto.AlbumDataTransfer{}

	for _, album := range albums {
		albumDto, err := Gateway.GetFullAlbumByArtist(userId, useCase.trackAgent, useCase.albumAgent, album, artist)
		if err != nil {
			return nil, err
		}
		albumsDto[album.Id] = albumDto
	}

	liked, err := useCase.artistAgent.LikeCheckByUser(userId, artist.Id)
	if err != nil {
		liked = false
	}

	return &artistProto.ArtistDataTransfer{
		Id:      artist.Id,
		Name:    artist.Name,
		Cover:   coverPath,
		Likes:   artist.CountLikes,
		Albums:  albumsDto,
		IsLiked: liked,
	}, nil
}

func (useCase artistUseCase) GetAll(userId int64) ([]*artistProto.ArtistDataTransfer, error) {
	albums, err := useCase.artistAgent.GetAll()

	if err != nil {
		return nil, err
	}

	dto := make([]*artistProto.ArtistDataTransfer, len(albums))

	for idx, obj := range albums {
		result, err := useCase.CastToDTO(userId, obj)
		if err != nil {
			return nil, err
		}
		dto[idx] = result
	}
	return dto, nil
}

func (useCase artistUseCase) GetLastId() (int64, error) {
	id, err := useCase.artistAgent.GetLastId()
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (useCase artistUseCase) Create(artist *artistProto.Artist) error {
	err := useCase.artistAgent.Create(artist)
	return err
}

func (useCase artistUseCase) Update(artist *artistProto.Artist) error {
	err := useCase.artistAgent.Update(artist)
	return err
}

func (useCase artistUseCase) Delete(id int64) error {
	err := useCase.artistAgent.Delete(id)
	return err
}

func (useCase artistUseCase) GetById(id int64, userId int64) (*artistProto.ArtistDataTransfer, error) {
	artist, err := useCase.artistAgent.GetById(id)
	if err != nil {
		return nil, err
	}

	dto, err := useCase.CastToDTO(userId, artist)
	if err != nil {
		return nil, err
	}
	return dto, err
}

func (useCase artistUseCase) GetPopular(userId int64) ([]*artistProto.ArtistDataTransfer, error) {
	artists, err := useCase.artistAgent.GetPopular()

	if err != nil {
		return nil, err
	}

	dto := make([]*artistProto.ArtistDataTransfer, len(artists))

	for idx, obj := range artists {
		result, err := useCase.CastToDTO(userId, obj)
		if err != nil {
			return nil, err
		}
		dto[idx] = result
	}
	return dto, nil
}

func (useCase artistUseCase) GetSize() (int64, error) {
	return useCase.artistAgent.GetSize()
}

func (useCase artistUseCase) SearchByName(userId int64, title string) ([]*artistProto.ArtistDataTransfer, error) {
	artists, err := useCase.artistAgent.SearchByName(title)

	if err != nil {
		return nil, err
	}

	dto := make([]*artistProto.ArtistDataTransfer, len(artists))

	for idx, obj := range artists {
		result, err := useCase.CastToDTO(userId, obj)
		if err != nil {
			return nil, err
		}
		dto[idx] = result
	}
	return dto, nil
}

func (useCase artistUseCase) GetFavorites(userId int64) ([]*artistProto.ArtistDataTransfer, error) {
	artists, err := useCase.artistAgent.GetFavorites(userId)

	if err != nil {
		return nil, err
	}

	dto := make([]*artistProto.ArtistDataTransfer, len(artists))

	for idx, obj := range artists {
		result, err := useCase.CastToDTO(userId, obj)
		if err != nil {
			return nil, err
		}
		dto[idx] = result
	}
	return dto, nil
}

func (useCase artistUseCase) AddToFavorites(userId int64, albumId int64) error {
	return useCase.artistAgent.AddToFavorites(userId, albumId)
}

func (useCase artistUseCase) RemoveFromFavorites(userId int64, albumId int64) error {
	return useCase.artistAgent.RemoveFromFavorites(userId, albumId)
}

func (useCase artistUseCase) Like(artistId int64, userId int64) error {
	err := useCase.artistAgent.Like(userId, artistId)
	return err
}

func (useCase artistUseCase) LikeCheckByUser(artistId int64, userId int64) (bool, error) {
	return useCase.artistAgent.LikeCheckByUser(userId, artistId)
}
