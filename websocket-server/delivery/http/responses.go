package http

import "github.com/go-park-mail-ru/2022_1_Wave/websocket-server/domain"

func getNoTrackStateMessage() *domain.UserPlayerUpdateStateMessage {
	return &domain.UserPlayerUpdateStateMessage{TypePushState: domain.NoTrackState}
}

func getTrackStateMessage(trackState *domain.UserPlayerState) *domain.UserPlayerUpdateStateMessage {
	return &domain.UserPlayerUpdateStateMessage{
		TypePushState: domain.NewTracksQueue,
		Data:          *trackState,
	}
}

func getInvalidTrackStateFormatMessage() *domain.UserPlayerUpdateStateMessage {
	return &domain.UserPlayerUpdateStateMessage{
		TypePushState: domain.InvalidTrackStateFormat,
	}
}
