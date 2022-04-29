package PlaylistUseCase

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal/domain"
	Gateway "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/gateway"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/playlist/playlistProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/track/trackProto"
)

type UseCase interface {
	GetAll(userId int64) ([]*playlistProto.PlaylistDataTransfer, error)
	GetLastId(userId int64) (int64, error)
	Create(userId int64, playlist *playlistProto.Playlist) error
	Update(userId int64, playlist *playlistProto.Playlist) error
	Delete(userId int64, playlistId int64) error
	GetById(userId int64, playlistId int64) ([]playlistProto.PlaylistDataTransfer, error)
	GetSize(userId int64) (int64, error)
}

type playlistUseCase struct {
	playlistAgent domain.PlaylistAgent
	artistAgent   domain.ArtistAgent
	trackAgent    domain.TrackAgent
}

func NewPlaylistUseCase(playlistAgent domain.PlaylistAgent, artistAgent domain.ArtistAgent, trackAgent domain.TrackAgent) *playlistUseCase {
	return &playlistUseCase{
		playlistAgent: playlistAgent,
		artistAgent:   artistAgent,
		trackAgent:    trackAgent,
	}
}

func (useCase playlistUseCase) CastToDTO(playlist *playlistProto.Playlist) (*playlistProto.PlaylistDataTransfer, error) {
	tracks := make([]*trackProto.TrackDataTransfer, len(playlist.TracksId))
	for idx, trackId := range playlist.TracksId {
		track, err := useCase.trackAgent.GetById(trackId)
		if err != nil {
			return nil, err
		}

		artist, err := useCase.artistAgent.GetById(track.ArtistId)
		if err != nil {
			return nil, err
		}

		trackDto, err := Gateway.CastTrackToDtoWithoutArtistName(track)
		if err != nil {
			return nil, err
		}
		trackDto.Artist = artist.Name

		tracks[idx] = trackDto
	}

	playlistDto := &playlistProto.PlaylistDataTransfer{
		Id:     playlist.Id,
		Title:  playlist.Title,
		Tracks: tracks,
	}

	return playlistDto, nil
}

func (useCase playlistUseCase) GetAll(userId int64) ([]*playlistProto.PlaylistDataTransfer, error) {
	playlists, err := useCase.playlistAgent.GetAll(userId)

	if err != nil {
		return nil, err
	}

	dto := make([]*playlistProto.PlaylistDataTransfer, len(playlists))

	for idx, obj := range playlists {
		result, err := useCase.CastToDTO(obj)
		if err != nil {
			return nil, err
		}
		dto[idx] = result
	}
	return dto, nil
}

func (useCase playlistUseCase) GetLastId(userId int64) (int64, error) {
	id, err := useCase.playlistAgent.GetLastId(userId)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (useCase playlistUseCase) Create(userId int64, playlist *playlistProto.Playlist) error {
	err := useCase.playlistAgent.Create(userId, playlist)
	return err
}

func (useCase playlistUseCase) Update(userId int64, playlist *playlistProto.Playlist) error {
	err := useCase.playlistAgent.Update(userId, playlist)
	return err
}

func (useCase playlistUseCase) Delete(userId int64, id int64) error {
	err := useCase.playlistAgent.Delete(userId, id)
	return err
}

func (useCase playlistUseCase) GetById(userId int64, id int64) (*playlistProto.PlaylistDataTransfer, error) {
	playlist, err := useCase.playlistAgent.GetById(userId, id)
	if err != nil {
		return nil, err
	}

	dto, err := useCase.CastToDTO(playlist)
	if err != nil {
		return nil, err
	}
	return dto, err
}

func (useCase playlistUseCase) GetSize(userId int64) (int64, error) {
	return useCase.playlistAgent.GetSize(userId)
}
