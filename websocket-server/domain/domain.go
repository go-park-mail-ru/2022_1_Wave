package domain

import "time"

type FromIs string
type TypePushState string

const (
	PlaylistFromIs  FromIs = "playlist"
	AlbumFromIs            = "album"
	FavoritesFromIs        = "favorites"
)

const (
	PlayNewTrack            TypePushState = "new_track"
	OnPause                               = "on_pause"
	OffPause                              = "off_pause"
	ChangePosition                        = "change_position"
	NoTrackState                          = "no_track_state"
	InvalidTrackStateFormat               = "invalid_format"
)

type UserPlayerState struct {
	TrackId         uint      `json:"track_id,omitempty"`
	FromIs          FromIs    `json:"from_is,omitempty"`
	FromIsId        uint      `json:"from_is_id,omitempty"`
	OnPause         bool      `json:"on_pause,omitempty"`
	LastSecPosition uint      `json:"last_sec_position,omitempty"`
	TimeStateUpdate time.Time `json:"time_state_update,omitempty"`
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
	NewTrackUpdateState(userId uint, trackId uint, fromIs FromIs, fromIsId uint, timeStateUpdate time.Time) error
	OnPauseUpdateState(userId uint, timeStateUpdate time.Time) error
	OffPauseUpdateState(userId uint, timeStateUpdate time.Time) error
	ChangePositionUpdateState(userId uint, lastSecPosition uint, timeStateUpdate time.Time) error
	GetTrackState(userId uint) (*UserPlayerState, error)
}
