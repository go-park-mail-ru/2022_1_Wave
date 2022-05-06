package TrackUseCase

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal/domain"
	Gateway "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/gateway"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/track/trackProto"
)

type UseCase interface {
	GetAll(userId int64) ([]*trackProto.TrackDataTransfer, error)
	GetLastId() (int64, error)
	Create(transfer *trackProto.Track) error
	Update(transfer *trackProto.Track) error
	Delete(int64) error
	GetById(trackId int64, userId int64) (*trackProto.TrackDataTransfer, error)
	GetPopular(userId int64) ([]*trackProto.TrackDataTransfer, error)
	GetTracksFromAlbum(albumId int64, userId int64) ([]*trackProto.TrackDataTransfer, error)
	GetPopularTracksFromArtist(artistId int64, userId int64) ([]*trackProto.TrackDataTransfer, error)
	GetSize() (int64, error)
	Like(arg int64, userId int64) error
	LikeCheckByUser(arg int64, userId int64) (bool, error)
	Listen(arg int64) error
	SearchByTitle(arg string, userId int64) ([]*trackProto.TrackDataTransfer, error)
	GetFavorites(userId int64) ([]*trackProto.TrackDataTransfer, error)
	AddToFavorites(userId int64, trackId int64) error
	RemoveFromFavorites(userId int64, trackId int64) error
	GetTracksFromPlaylist(playlistId int64, userId int64) ([]*trackProto.TrackDataTransfer, error)
}

type trackUseCase struct {
	albumAgent    domain.AlbumAgent
	artistAgent   domain.ArtistAgent
	trackAgent    domain.TrackAgent
	playlistAgent domain.PlaylistAgent
}

func NewTrackUseCase(albumAgent domain.AlbumAgent, artistAgent domain.ArtistAgent, trackAgent domain.TrackAgent) *trackUseCase {
	return &trackUseCase{
		albumAgent:  albumAgent,
		artistAgent: artistAgent,
		trackAgent:  trackAgent,
	}
}

func (useCase trackUseCase) CastToDTO(track *trackProto.Track, userId int64) (*trackProto.TrackDataTransfer, error) {
	artist, err := useCase.artistAgent.GetById(track.ArtistId)
	if err != nil {
		return nil, err
	}

	trackDto, err := Gateway.CastTrackToDtoWithoutArtistName(track, useCase.trackAgent, userId)
	if err != nil {
		return nil, err
	}

	trackDto.Artist = artist.Name

	return trackDto, nil
}

func (useCase trackUseCase) castArray(userId int64, tracks []*trackProto.Track) ([]*trackProto.TrackDataTransfer, error) {
	dto := make([]*trackProto.TrackDataTransfer, len(tracks))
	for idx, obj := range tracks {
		result, err := useCase.CastToDTO(obj, userId)
		if err != nil {
			return nil, err
		}
		dto[idx] = result
	}
	return dto, nil
}

func (useCase trackUseCase) GetAll(userId int64) ([]*trackProto.TrackDataTransfer, error) {
	tracks, err := useCase.trackAgent.GetAll()

	if err != nil {
		return nil, err
	}

	return useCase.castArray(userId, tracks)

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

func (useCase trackUseCase) GetById(id int64, userId int64) (*trackProto.TrackDataTransfer, error) {
	track, err := useCase.trackAgent.GetById(id)
	if err != nil {
		return nil, err
	}

	return useCase.CastToDTO(track, userId)
}

func (useCase trackUseCase) GetPopular(userId int64) ([]*trackProto.TrackDataTransfer, error) {
	tracks, err := useCase.trackAgent.GetPopular()

	if err != nil {
		return nil, err
	}

	return useCase.castArray(userId, tracks)
}

func (useCase trackUseCase) GetSize() (int64, error) {
	return useCase.trackAgent.GetSize()
}

func (useCase trackUseCase) SearchByTitle(title string, userId int64) ([]*trackProto.TrackDataTransfer, error) {
	tracks, err := useCase.trackAgent.SearchByTitle(title)

	if err != nil {
		return nil, err
	}

	return useCase.castArray(userId, tracks)
}

func (useCase trackUseCase) GetPopularTracksFromArtist(id int64, userId int64) ([]*trackProto.TrackDataTransfer, error) {
	tracks, err := useCase.trackAgent.GetPopularTracksFromArtist(id)

	if err != nil {
		return nil, err
	}

	return useCase.castArray(userId, tracks)
}

func (useCase trackUseCase) GetTracksFromAlbum(id int64, userId int64) ([]*trackProto.TrackDataTransfer, error) {
	tracks, err := useCase.trackAgent.GetTracksFromAlbum(id)

	if err != nil {
		return nil, err
	}

	return useCase.castArray(userId, tracks)
}

func (useCase trackUseCase) Like(trackId int64, userId int64) error {
	err := useCase.trackAgent.Like(trackId, userId)
	return err
}

func (useCase trackUseCase) LikeCheckByUser(trackId int64, userId int64) (bool, error) {
	return useCase.trackAgent.LikeCheckByUser(userId, trackId)
}

func (useCase trackUseCase) Listen(trackId int64) error {
	err := useCase.trackAgent.Listen(trackId)
	return err
}

func (useCase trackUseCase) GetFavorites(userId int64) ([]*trackProto.TrackDataTransfer, error) {
	tracks, err := useCase.trackAgent.GetFavorites(userId)

	if err != nil {
		return nil, err
	}

	return useCase.castArray(userId, tracks)
}

func (useCase trackUseCase) AddToFavorites(userId int64, albumId int64) error {
	return useCase.trackAgent.AddToFavorites(userId, albumId)
}

func (useCase trackUseCase) RemoveFromFavorites(userId int64, albumId int64) error {
	return useCase.trackAgent.RemoveFromFavorites(userId, albumId)
}

func (useCase trackUseCase) GetTracksFromPlaylist(playlistId int64, userId int64) ([]*trackProto.TrackDataTransfer, error) {
	tracks, err := useCase.trackAgent.GetTracksFromPlaylist(playlistId)
	if err != nil {
		return nil, err
	}

	return useCase.castArray(userId, tracks)
}
