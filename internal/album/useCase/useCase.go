package AlbumUseCase

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/album/albumProto"
	Gateway "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/gateway"
)

type AlbumUseCase interface {
	GetAll(userId int64) ([]*albumProto.AlbumDataTransfer, error)
	GetAllCovers() ([]*albumProto.AlbumCoverDataTransfer, error)
	GetLastId() (int64, error)
	GetLastCoverId() (int64, error)
	Create(*albumProto.Album) error
	CreateCover(*albumProto.AlbumCover) error
	Update(*albumProto.Album) error
	UpdateCover(*albumProto.AlbumCover) error
	Delete(int64) error
	DeleteCover(int64) error
	GetById(userId int64, albumId int64) (*albumProto.AlbumDataTransfer, error)
	GetCoverById(int64) (*albumProto.AlbumCoverDataTransfer, error)
	GetPopular(userId int64) ([]*albumProto.AlbumDataTransfer, error)
	GetAlbumsFromArtist(userId int64, artistId int64) ([]*albumProto.AlbumDataTransfer, error)
	GetSize() (int64, error)
	SearchByTitle(userId int64, title string) ([]*albumProto.AlbumDataTransfer, error)
	GetFavorites(userId int64) ([]*albumProto.AlbumDataTransfer, error)
	AddToFavorites(userId int64, albumId int64) error
	RemoveFromFavorites(userId int64, albumId int64) error
	Like(arg int64, userId int64) error
	LikeCheckByUser(arg int64, userId int64) (bool, error)
	GetPopularAlbumOfWeek(userId int64) ([]*albumProto.AlbumDataTransfer, error)
}

type albumUseCase struct {
	albumAgent  domain.AlbumAgent
	artistAgent domain.ArtistAgent
	trackAgent  domain.TrackAgent
}

func NewAlbumUseCase(albumAgent domain.AlbumAgent, artistAgent domain.ArtistAgent, trackAgent domain.TrackAgent) *albumUseCase {
	return &albumUseCase{
		albumAgent:  albumAgent,
		artistAgent: artistAgent,
		trackAgent:  trackAgent,
	}
}

func (useCase albumUseCase) CastToDTO(userId int64, album *albumProto.Album) (*albumProto.AlbumDataTransfer, error) {
	artist, err := useCase.artistAgent.GetById(album.ArtistId)
	if err != nil {
		return nil, err
	}
	return Gateway.GetFullAlbumByArtist(userId, useCase.trackAgent, useCase.albumAgent, album, artist)
}

func (useCase albumUseCase) CastCoverToDTO(cover *albumProto.AlbumCover) (*albumProto.AlbumCoverDataTransfer, error) {
	return &albumProto.AlbumCoverDataTransfer{
		Quote:  cover.Quote,
		IsDark: cover.IsDark,
	}, nil
}

func (useCase albumUseCase) GetAll(userId int64) ([]*albumProto.AlbumDataTransfer, error) {
	albums, err := useCase.albumAgent.GetAll()

	if err != nil {
		return nil, err
	}

	dto := make([]*albumProto.AlbumDataTransfer, len(albums))

	for idx, obj := range albums {
		result, err := useCase.CastToDTO(userId, obj)
		if err != nil {
			return nil, err
		}
		dto[idx] = result
	}
	return dto, nil
}

func (useCase albumUseCase) GetAllCovers() ([]*albumProto.AlbumCoverDataTransfer, error) {
	covers, err := useCase.albumAgent.GetAllCovers()

	if err != nil {
		return nil, err
	}

	dto := make([]*albumProto.AlbumCoverDataTransfer, len(covers))

	for idx, obj := range covers {
		result, err := useCase.CastCoverToDTO(obj)
		if err != nil {
			return nil, err
		}
		dto[idx] = result
	}
	return dto, nil
}

func (useCase albumUseCase) GetLastId() (int64, error) {
	id, err := useCase.albumAgent.GetLastId()
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (useCase albumUseCase) GetLastCoverId() (int64, error) {
	id, err := useCase.albumAgent.GetLastCoverId()
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (useCase albumUseCase) Create(album *albumProto.Album) error {
	err := useCase.albumAgent.Create(album)
	return err
}

func (useCase albumUseCase) CreateCover(cover *albumProto.AlbumCover) error {
	err := useCase.albumAgent.CreateCover(cover)
	return err
}

func (useCase albumUseCase) Update(album *albumProto.Album) error {
	err := useCase.albumAgent.Update(album)
	return err
}

func (useCase albumUseCase) UpdateCover(cover *albumProto.AlbumCover) error {
	err := useCase.albumAgent.UpdateCover(cover)
	return err
}

func (useCase albumUseCase) Delete(id int64) error {
	err := useCase.albumAgent.Delete(id)
	return err
}

func (useCase albumUseCase) DeleteCover(id int64) error {
	err := useCase.albumAgent.DeleteCover(id)
	return err
}

func (useCase albumUseCase) GetById(userId int64, id int64) (*albumProto.AlbumDataTransfer, error) {
	album, err := useCase.albumAgent.GetById(id)
	if err != nil {
		return nil, err
	}

	dto, err := useCase.CastToDTO(userId, album)
	if err != nil {
		return nil, err
	}
	return dto, err
}

func (useCase albumUseCase) GetCoverById(id int64) (*albumProto.AlbumCoverDataTransfer, error) {
	cover, err := useCase.albumAgent.GetCoverById(id)
	if err != nil {
		return nil, err
	}

	dto, err := useCase.CastCoverToDTO(cover)
	if err != nil {
		return nil, err
	}
	return dto, err
}

func (useCase albumUseCase) GetPopular(userId int64) ([]*albumProto.AlbumDataTransfer, error) {
	albums, err := useCase.albumAgent.GetPopular()

	if err != nil {
		return nil, err
	}

	dto := make([]*albumProto.AlbumDataTransfer, len(albums))

	for idx, obj := range albums {
		result, err := useCase.CastToDTO(userId, obj)
		if err != nil {
			return nil, err
		}
		dto[idx] = result
	}
	return dto, nil

}

func (useCase albumUseCase) GetAlbumsFromArtist(userId int64, artistId int64) ([]*albumProto.AlbumDataTransfer, error) {
	albums, err := useCase.albumAgent.GetAlbumsFromArtist(artistId)

	if err != nil {
		return nil, err
	}

	dto := make([]*albumProto.AlbumDataTransfer, len(albums))

	for idx, obj := range albums {
		result, err := useCase.CastToDTO(userId, obj)
		if err != nil {
			return nil, err
		}
		dto[idx] = result
	}
	return dto, nil
}

func (useCase albumUseCase) GetSize() (int64, error) {
	return useCase.albumAgent.GetSize()
}

func (useCase albumUseCase) SearchByTitle(userId int64, title string) ([]*albumProto.AlbumDataTransfer, error) {
	albums, err := useCase.albumAgent.SearchByTitle(title)

	if err != nil {
		return nil, err
	}

	dto := make([]*albumProto.AlbumDataTransfer, len(albums))

	for idx, obj := range albums {
		result, err := useCase.CastToDTO(userId, obj)
		if err != nil {
			return nil, err
		}
		dto[idx] = result
	}
	return dto, nil
}

func (useCase albumUseCase) GetFavorites(id int64) ([]*albumProto.AlbumDataTransfer, error) {
	albums, err := useCase.albumAgent.GetFavorites(id)

	if err != nil {
		return nil, err
	}

	dto := make([]*albumProto.AlbumDataTransfer, len(albums))

	for idx, obj := range albums {
		result, err := useCase.CastToDTO(id, obj)
		if err != nil {
			return nil, err
		}
		dto[idx] = result
	}
	return dto, nil
}

func (useCase albumUseCase) AddToFavorites(userId int64, albumId int64) error {
	return useCase.albumAgent.AddToFavorites(userId, albumId)
}

func (useCase albumUseCase) RemoveFromFavorites(userId int64, albumId int64) error {
	return useCase.albumAgent.RemoveFromFavorites(userId, albumId)
}

func (useCase albumUseCase) Like(albumId int64, userId int64) error {
	err := useCase.albumAgent.Like(userId, albumId)
	return err
}

func (useCase albumUseCase) LikeCheckByUser(albumId int64, userId int64) (bool, error) {
	return useCase.albumAgent.LikeCheckByUser(userId, albumId)
}

func (useCase albumUseCase) GetPopularAlbumOfWeek(userId int64) ([]*albumProto.AlbumDataTransfer, error) {
	albums, err := useCase.albumAgent.GetPopularAlbumOfWeekTop20()

	if err != nil {
		return nil, err
	}

	dto := make([]*albumProto.AlbumDataTransfer, len(albums))

	for idx, obj := range albums {
		result, err := useCase.CastToDTO(userId, obj)
		if err != nil {
			return nil, err
		}
		dto[idx] = result
	}
	return dto, nil

}
