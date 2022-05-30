package domain

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/track/trackProto"
	"golang.org/x/sys/unix"
)

type FromIs string
type TypePushState string

const (
	PlaylistFromIs  FromIs = "playlist"
	AlbumFromIs     FromIs = "album"
	FavoritesFromIs FromIs = "favorites"
)

const (
	PushTrackInQueue        TypePushState = "push_track"
	NewTracksQueue          TypePushState = "new_tracks_queue"
	NewTrackInQueue         TypePushState = "new_track"
	OnPause                 TypePushState = "on_pause"
	OffPause                TypePushState = "off_pause"
	ChangePosition          TypePushState = "change_position"
	NoTrackState            TypePushState = "no_track_state"
	InvalidTrackStateFormat TypePushState = "invalid_format"
	GetPlayerState          TypePushState = "get_player_state"
)

type UserPlayerState struct {
	TracksQueue     []trackProto.TrackDataTransfer `json:"tracks_queue"`
	QueuePosition   int                            `json:"queue_position"`
	OnPause         bool                           `json:"on_pause"`
	LastSecPosition float64                        `json:"last_sec_position"`
	TimeStateUpdate unix.Time_t                    `json:"time_state_update"`
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
	PushTrackUpdateState(userId uint, tracksToAdd []trackProto.TrackDataTransfer) error
	NewTrackQueueUpdateState(userId uint, tracksQueue []trackProto.TrackDataTransfer, queuePosition int, lastSecPosition float64, timeStateUpdate unix.Time_t) error
	NewTrackUpdateState(userId uint, queuePosition int, timeStateUpdate unix.Time_t) error
	OnPauseUpdateState(userId uint, lastSecPosition float64, timeStateUpdate unix.Time_t) error
	OffPauseUpdateState(userId uint, timeStateUpdate unix.Time_t) error
	ChangePositionUpdateState(userId uint, lastSecPosition float64, timeStateUpdate unix.Time_t) error
	GetTrackState(userId uint) (*UserPlayerState, error)
}
