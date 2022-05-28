package domain

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/track/trackProto"
	"golang.org/x/sys/unix"
)

type FromIs string
type TypePushState string

const (
	PlaylistFromIs  FromIs = "playlist"
	AlbumFromIs            = "album"
	FavoritesFromIs        = "favorites"
)

const (
	PushTrackInQueue        TypePushState = "push_track"
	NewTracksQueue                        = "new_tracks_queue"
	NewTrackInQueue                       = "new_track"
	OnPause                               = "on_pause"
	OffPause                              = "off_pause"
	ChangePosition                        = "change_position"
	NoTrackState                          = "no_track_state"
	InvalidTrackStateFormat               = "invalid_format"
	GetPlayerState                        = "get_player_state"
)

type UserPlayerState struct {
	TracksQueue     []trackProto.Track `json:"tracks_queue"`
	QueuePosition   int                `json:"queue_position"`
	OnPause         bool               `json:"on_pause"`
	LastSecPosition float64            `json:"last_sec_position"`
	TimeStateUpdate unix.Time_t        `json:"time_state_update"`
}

// сообщения такого типа будут приходить от клиента
type UserPlayerUpdateStateMessage struct {
	TypePushState TypePushState   `json:"type_push_state"`
	Data          UserPlayerState `json:"data,omitempty"`
}

type UserSyncPlayerRepo interface {
	CreateUserPlayerState(userId uint, state *UserPlayerState) error
	UpdateUserPlayerState(userId uint, state *UserPlayerState) error
	GetUserPlayerState(userId uint) (*UserPlayerState, error)
	DeleteUserPlayerState(userId uint) error
}

type UserSyncPlayerUseCase interface {
	PushTrackUpdateState(userId uint, tracksToAdd []trackProto.Track) error
	NewTrackQueueUpdateState(userId uint, tracksQueue []trackProto.Track, queuePosition int, lastSecPosition float64, timeStateUpdate unix.Time_t) error
	NewTrackUpdateState(userId uint, queuePosition int, timeStateUpdate unix.Time_t) error
	OnPauseUpdateState(userId uint, timeStateUpdate unix.Time_t) error
	OffPauseUpdateState(userId uint, timeStateUpdate unix.Time_t) error
	ChangePositionUpdateState(userId uint, lastSecPosition float64, timeStateUpdate unix.Time_t) error
	GetTrackState(userId uint) (*UserPlayerState, error)
}
