package usecase

import (
	"github.com/go-park-mail-ru/2022_1_Wave/websocket-server/domain"
	"time"
)

type userSyncPlayerUseCase struct {
	syncPlayerRepo domain.UserSyncPlayerRepo
}

func NewUserSyncPlayerUseCase(repo domain.UserSyncPlayerRepo) domain.UserSyncPlayerUseCase {
	return &userSyncPlayerUseCase{
		syncPlayerRepo: repo,
	}
}

func (a *userSyncPlayerUseCase) NewTrackUpdateState(userId uint, trackId uint, fromIs domain.FromIs, fromIsId uint, timeStateUpdate time.Time) error {
	oldTrack, err := a.syncPlayerRepo.GetUserPlayerState(userId)
	if err != nil {
		return err
	}

	if fromIs != "" {
		oldTrack.FromIs = fromIs
		oldTrack.FromIsId = fromIsId
	}

	oldTrack.TrackId = trackId
	oldTrack.TimeStateUpdate = timeStateUpdate

	err = a.syncPlayerRepo.UpdateUserPlayerState(userId, oldTrack)

	return err
}

func (a *userSyncPlayerUseCase) OnPauseUpdateState(userId uint, timeStateUpdate time.Time) error {
	oldTrack, err := a.syncPlayerRepo.GetUserPlayerState(userId)
	if err != nil {
		return err
	}

	oldTrack.LastSecPosition = uint(timeStateUpdate.Sub(oldTrack.TimeStateUpdate).Seconds())
	oldTrack.TimeStateUpdate = timeStateUpdate
	oldTrack.OnPause = true

	err = a.syncPlayerRepo.UpdateUserPlayerState(userId, oldTrack)

	return err
}

func (a *userSyncPlayerUseCase) OffPauseUpdateState(userId uint, timeStateUpdate time.Time) error {
	oldTrack, err := a.syncPlayerRepo.GetUserPlayerState(userId)
	if err != nil {
		return err
	}

	oldTrack.TimeStateUpdate = timeStateUpdate
	oldTrack.OnPause = false

	err = a.syncPlayerRepo.UpdateUserPlayerState(userId, oldTrack)

	return err
}

func (a *userSyncPlayerUseCase) ChangePositionUpdateState(userId uint, lastSecPosition uint, timeStateUpdate time.Time) error {
	oldTrack, err := a.syncPlayerRepo.GetUserPlayerState(userId)
	if err != nil {
		return err
	}

	oldTrack.TimeStateUpdate = timeStateUpdate
	oldTrack.LastSecPosition = lastSecPosition

	err = a.syncPlayerRepo.UpdateUserPlayerState(userId, oldTrack)

	return err
}

func (a *userSyncPlayerUseCase) GetTrackState(userId uint) (*domain.UserPlayerState, error) {
	track, err := a.syncPlayerRepo.GetUserPlayerState(userId)
	if err != nil {
		return nil, err
	}

	return track, nil
}
