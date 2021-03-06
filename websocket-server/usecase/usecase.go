package usecase

import (
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/track/trackProto"
	"github.com/go-park-mail-ru/2022_1_Wave/websocket-server/domain"
	"golang.org/x/sys/unix"
)

type userSyncPlayerUseCase struct {
	syncPlayerRepo domain.UserSyncPlayerRepo
}

func NewUserSyncPlayerUseCase(repo domain.UserSyncPlayerRepo) domain.UserSyncPlayerUseCase {
	return &userSyncPlayerUseCase{
		syncPlayerRepo: repo,
	}
}

func (a *userSyncPlayerUseCase) PushTrackUpdateState(userId uint, tracksToAdd []trackProto.TrackDataTransfer) error {
	oldTrack, err := a.syncPlayerRepo.GetUserPlayerState(userId)
	if err != nil {
		return err
	}

	oldTrack.TracksQueue = append(oldTrack.TracksQueue, tracksToAdd...)
	err = a.syncPlayerRepo.UpdateUserPlayerState(userId, oldTrack)

	return err
}

func (a *userSyncPlayerUseCase) NewTrackQueueUpdateState(userId uint, tracksQueue []trackProto.TrackDataTransfer, queuePosition int, lastSecPosition float64, timeStateUpdate unix.Time_t) error {
	var oldTrack *domain.UserPlayerState
	oldTrack, err := a.syncPlayerRepo.GetUserPlayerState(userId)
	if err != nil {
		oldTrack = new(domain.UserPlayerState)
	}

	oldTrack.TracksQueue = tracksQueue
	oldTrack.QueuePosition = queuePosition
	oldTrack.TimeStateUpdate = timeStateUpdate
	oldTrack.OnPause = false
	oldTrack.LastSecPosition = lastSecPosition

	err = a.syncPlayerRepo.UpdateUserPlayerState(userId, oldTrack)
	return err
}

func (a *userSyncPlayerUseCase) NewTrackUpdateState(userId uint, queuePosition int, timeStateUpdate unix.Time_t) error {
	oldTrack, err := a.syncPlayerRepo.GetUserPlayerState(userId)
	if err != nil {
		return err
	}

	oldTrack.QueuePosition = queuePosition
	oldTrack.TimeStateUpdate = timeStateUpdate
	oldTrack.OnPause = false
	oldTrack.LastSecPosition = 0

	err = a.syncPlayerRepo.UpdateUserPlayerState(userId, oldTrack)

	return err
}

func (a *userSyncPlayerUseCase) OnPauseUpdateState(userId uint, lastSecPosition float64, timeStateUpdate unix.Time_t) error {
	oldTrack, err := a.syncPlayerRepo.GetUserPlayerState(userId)
	if err != nil {
		return err
	}

	oldTrack.LastSecPosition = lastSecPosition
	oldTrack.TimeStateUpdate = timeStateUpdate
	oldTrack.OnPause = true

	err = a.syncPlayerRepo.UpdateUserPlayerState(userId, oldTrack)

	return err
}

func (a *userSyncPlayerUseCase) OffPauseUpdateState(userId uint, timeStateUpdate unix.Time_t) error {
	oldTrack, err := a.syncPlayerRepo.GetUserPlayerState(userId)
	if err != nil {
		return err
	}

	oldTrack.TimeStateUpdate = timeStateUpdate
	oldTrack.OnPause = false

	err = a.syncPlayerRepo.UpdateUserPlayerState(userId, oldTrack)

	return err
}

func (a *userSyncPlayerUseCase) ChangePositionUpdateState(userId uint, lastSecPosition float64, timeStateUpdate unix.Time_t) error {
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
