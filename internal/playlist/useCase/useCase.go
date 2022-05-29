package PlaylistUseCase

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal/domain"
	Gateway "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/gateway"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/playlist/playlistProto"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/track/trackProto"
)

type PlaylistUseCase interface {
	GetAllOfCurrentUser(userId int64) ([]*playlistProto.PlaylistDataTransfer, error)
	GetAll(userId int64) ([]*playlistProto.PlaylistDataTransfer, error)
	GetLastIdOfCurrentUser(userId int64) (int64, error)
	GetLastId() (int64, error)
	Create(userId int64, playlist *playlistProto.Playlist) error
	Update(userId int64, playlist *playlistProto.Playlist) error
	Delete(userId int64, playlistId int64) error
	GetByIdOfCurrentUser(userId int64, playlistId int64) (*playlistProto.PlaylistDataTransfer, error)
	GetById(playlistId int64, userId int64) (*playlistProto.PlaylistDataTransfer, error)
	GetSizeOfCurrentUser(userId int64) (int64, error)
	GetSize() (int64, error)
	AddToPlaylist(userId int64, playlistId int64, trackId int64) error
	RemoveFromPlaylist(userId int64, playlistId int64, trackId int64) error
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

func (useCase playlistUseCase) CastToDTO(userId int64, playlist *playlistProto.Playlist) (*playlistProto.PlaylistDataTransfer, error) {
	tracksFromPlaylist, err := useCase.trackAgent.GetTracksFromPlaylist(playlist.Id)
	if err != nil {
		return nil, err
	}

	dto := make([]*trackProto.TrackDataTransfer, len(tracksFromPlaylist))
	for idx, track := range tracksFromPlaylist {
		artist, err := useCase.artistAgent.GetById(track.ArtistId)
		if err != nil {
			return nil, err
		}
		dto[idx], err = Gateway.CastTrackToDto(track, artist, useCase.trackAgent, userId)
		if err != nil {
			return nil, err
		}
	}

	playlistDto := &playlistProto.PlaylistDataTransfer{
		Id:     playlist.Id,
		Title:  playlist.Title,
		Tracks: dto,
	}

	return playlistDto, nil
}

func (useCase playlistUseCase) GetAllOfCurrentUser(userId int64) ([]*playlistProto.PlaylistDataTransfer, error) {
	playlists, err := useCase.playlistAgent.GetAllOfCurrentUser(userId)

	if err != nil {
		return nil, err
	}

	dto := make([]*playlistProto.PlaylistDataTransfer, len(playlists))

	for idx, obj := range playlists {
		result, err := useCase.CastToDTO(userId, obj)
		if err != nil {
			return nil, err
		}
		dto[idx] = result
	}
	return dto, nil
}

func (useCase playlistUseCase) GetAll(userId int64) ([]*playlistProto.PlaylistDataTransfer, error) {
	playlists, err := useCase.playlistAgent.GetAll()

	if err != nil {
		return nil, err
	}

	dto := make([]*playlistProto.PlaylistDataTransfer, len(playlists))

	for idx, obj := range playlists {
		result, err := useCase.CastToDTO(userId, obj)
		if err != nil {
			return nil, err
		}
		dto[idx] = result
	}
	return dto, nil
}

func (useCase playlistUseCase) GetLastIdOfCurrentUser(userId int64) (int64, error) {
	id, err := useCase.playlistAgent.GetLastIdOfCurrentUser(userId)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (useCase playlistUseCase) GetLastId() (int64, error) {
	id, err := useCase.playlistAgent.GetLastId()
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

func (useCase playlistUseCase) GetByIdOfCurrentUser(userId int64, id int64) (*playlistProto.PlaylistDataTransfer, error) {
	playlist, err := useCase.playlistAgent.GetByIdOfCurrentUser(userId, id)
	if err != nil {
		return nil, err
	}

	dto, err := useCase.CastToDTO(userId, playlist)
	if err != nil {
		return nil, err
	}
	return dto, err
}

func (useCase playlistUseCase) GetById(id int64, userId int64) (*playlistProto.PlaylistDataTransfer, error) {
	playlist, err := useCase.playlistAgent.GetById(id)
	if err != nil {
		return nil, err
	}

	dto, err := useCase.CastToDTO(userId, playlist)
	if err != nil {
		return nil, err
	}
	return dto, err
}

func (useCase playlistUseCase) GetSizeOfCurrentUser(userId int64) (int64, error) {
	return useCase.playlistAgent.GetSizeOfCurrentUser(userId)
}

func (useCase playlistUseCase) GetSize() (int64, error) {
	return useCase.playlistAgent.GetSize()
}

func (useCase playlistUseCase) AddToPlaylist(userId int64, playlistId int64, trackId int64) error {
	return useCase.playlistAgent.AddToPlaylist(userId, playlistId, trackId)
}

func (useCase playlistUseCase) RemoveFromPlaylist(userId int64, playlistId int64, trackId int64) error {
	return useCase.playlistAgent.RemoveFromPlaylist(userId, playlistId, trackId)
}
