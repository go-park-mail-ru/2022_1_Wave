package TrackUseCase

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal/domain"
	Gateway "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/gateway"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/track/trackProto"
)

type UseCase interface {
	GetAll() ([]*trackProto.TrackDataTransfer, error)
	GetLastId() (int64, error)
	Create(transfer *trackProto.Track) error
	Update(transfer *trackProto.Track) error
	Delete(int64) error
	GetById(int64) (*trackProto.TrackDataTransfer, error)
	GetPopular() ([]*trackProto.TrackDataTransfer, error)
	GetTracksFromAlbum(int64) ([]*trackProto.TrackDataTransfer, error)
	GetPopularTracksFromArtist(int64) ([]*trackProto.TrackDataTransfer, error)
	GetSize() (int64, error)
	Like(arg int64, userId int64) error
	Listen(arg int64) error
	SearchByTitle(arg string) ([]*trackProto.TrackDataTransfer, error)
	GetFavorites(int64) ([]*trackProto.TrackDataTransfer, error)
	AddToFavorites(userId int64, trackId int64) error
	RemoveFromFavorites(userId int64, trackId int64) error
}

type trackUseCase struct {
	albumAgent  domain.AlbumAgent
	artistAgent domain.ArtistAgent
	trackAgent  domain.TrackAgent
}

func NewTrackUseCase(albumAgent domain.AlbumAgent, artistAgent domain.ArtistAgent, trackAgent domain.TrackAgent) *trackUseCase {
	return &trackUseCase{
		albumAgent:  albumAgent,
		artistAgent: artistAgent,
		trackAgent:  trackAgent,
	}
}

func (useCase trackUseCase) CastToDTO(track *trackProto.Track) (*trackProto.TrackDataTransfer, error) {
	artist, err := useCase.artistAgent.GetById(track.ArtistId)
	if err != nil {
		return nil, err
	}

	trackDto, err := Gateway.CastTrackToDtoWithoutArtistName(track)
	if err != nil {
		return nil, err
	}

	trackDto.Artist = artist.Name
	return trackDto, nil
}

func (useCase trackUseCase) GetAll() ([]*trackProto.TrackDataTransfer, error) {
	tracks, err := useCase.trackAgent.GetAll()

	if err != nil {
		return nil, err
	}

	dto := make([]*trackProto.TrackDataTransfer, len(tracks))

	for idx, obj := range tracks {
		result, err := useCase.CastToDTO(obj)
		if err != nil {
			return nil, err
		}
		dto[idx] = result
	}
	return dto, nil
}

func (useCase trackUseCase) GetLastId() (int64, error) {
	id, err := useCase.trackAgent.GetLastId()
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (useCase trackUseCase) Create(track *trackProto.Track) error {
	err := useCase.trackAgent.Create(track)
	return err
}

func (useCase trackUseCase) Update(track *trackProto.Track) error {
	err := useCase.trackAgent.Update(track)
	return err
}

func (useCase trackUseCase) Delete(id int64) error {
	err := useCase.trackAgent.Delete(id)
	return err
}

func (useCase trackUseCase) GetById(id int64) (*trackProto.TrackDataTransfer, error) {
	tracks, err := useCase.trackAgent.GetById(id)
	if err != nil {
		return nil, err
	}

	dto, err := useCase.CastToDTO(tracks)
	if err != nil {
		return nil, err
	}
	return dto, err
}

func (useCase trackUseCase) GetPopular() ([]*trackProto.TrackDataTransfer, error) {
	tracks, err := useCase.trackAgent.GetPopular()

	if err != nil {
		return nil, err
	}

	dto := make([]*trackProto.TrackDataTransfer, len(tracks))

	for idx, obj := range tracks {
		result, err := useCase.CastToDTO(obj)
		if err != nil {
			return nil, err
		}
		dto[idx] = result
	}
	return dto, nil
}

func (useCase trackUseCase) GetSize() (int64, error) {
	return useCase.trackAgent.GetSize()
}

func (useCase trackUseCase) SearchByTitle(title string) ([]*trackProto.TrackDataTransfer, error) {
	tracks, err := useCase.trackAgent.SearchByTitle(title)

	if err != nil {
		return nil, err
	}

	dto := make([]*trackProto.TrackDataTransfer, len(tracks))

	for idx, obj := range tracks {
		result, err := useCase.CastToDTO(obj)
		if err != nil {
			return nil, err
		}
		dto[idx] = result
	}
	return dto, nil
}

func (useCase trackUseCase) GetPopularTracksFromArtist(id int64) ([]*trackProto.TrackDataTransfer, error) {
	tracks, err := useCase.trackAgent.GetPopularTracksFromArtist(id)

	if err != nil {
		return nil, err
	}

	dto := make([]*trackProto.TrackDataTransfer, len(tracks))

	for idx, obj := range tracks {
		result, err := useCase.CastToDTO(obj)
		if err != nil {
			return nil, err
		}
		dto[idx] = result
	}
	return dto, nil
}

func (useCase trackUseCase) GetTracksFromAlbum(id int64) ([]*trackProto.TrackDataTransfer, error) {
	tracks, err := useCase.trackAgent.GetTracksFromAlbum(id)

	if err != nil {
		return nil, err
	}

	dto := make([]*trackProto.TrackDataTransfer, len(tracks))

	for idx, obj := range tracks {
		result, err := useCase.CastToDTO(obj)
		if err != nil {
			return nil, err
		}
		dto[idx] = result
	}
	return dto, nil
}

func (useCase trackUseCase) Like(trackId int64, userId int64) error {
	err := useCase.trackAgent.Like(trackId, userId)
	return err
}

func (useCase trackUseCase) Listen(trackId int64) error {
	err := useCase.trackAgent.Listen(trackId)
	return err
}

func (useCase trackUseCase) GetFavorites(id int64) ([]*trackProto.TrackDataTransfer, error) {
	tracks, err := useCase.trackAgent.GetFavorites(id)

	if err != nil {
		return nil, err
	}

	dto := make([]*trackProto.TrackDataTransfer, len(tracks))

	for idx, obj := range tracks {
		result, err := useCase.CastToDTO(obj)
		if err != nil {
			return nil, err
		}
		dto[idx] = result
	}
	return dto, nil
}

func (useCase trackUseCase) AddToFavorites(userId int64, albumId int64) error {
	return useCase.trackAgent.AddToFavorites(userId, albumId)
}

func (useCase trackUseCase) RemoveFromFavorites(userId int64, albumId int64) error {
	return useCase.trackAgent.RemoveFromFavorites(userId, albumId)
}
