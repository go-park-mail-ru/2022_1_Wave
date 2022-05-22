package http

import (
	"encoding/json"
	"errors"
	"fmt"
	auth_domain "github.com/go-park-mail-ru/2022_1_Wave/websocket-server/auth"
	"github.com/go-park-mail-ru/2022_1_Wave/websocket-server/domain"
	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Handler struct {
	userSyncPlayerUseCase domain.UserSyncPlayerUseCase
	authAgent             auth_domain.AuthAgent
	redisAddr             string
}

func NewHandler(usecase domain.UserSyncPlayerUseCase, redisAddr string, authAgent auth_domain.AuthAgent) *Handler {
	return &Handler{
		userSyncPlayerUseCase: usecase,
		redisAddr:             redisAddr,
		authAgent:             authAgent,
	}
}

func (a *Handler) getUserId(c echo.Context) (uint, error) {
	cookie, err := c.Cookie("session_id")
	if err != nil {
		return 0, errors.New("no session_id cookie")
	}

	session, err := a.authAgent.GetSession(cookie.Value)
	if err != nil {
		return 0, errors.New("cannot get session")
	}

	return session.UserId, nil
}

func getRedisChannelName(userId uint) string {
	return strconv.Itoa(int(userId))
}

func (a *Handler) initRedisStructs(userId uint) (redisCon redis.Conn, redisPubSub *redis.PubSubConn, err error) {
	redisCon, err = redis.Dial("tcp", a.redisAddr)
	if err != nil {
		return
	}

	redisPubSub = &redis.PubSubConn{Conn: redisCon}
	err = redisPubSub.Subscribe(getRedisChannelName(userId))

	return
}

func (a *Handler) pushToRedisChannel(redisCon redis.Conn, channelName string, message string) {
	redisCon.Do("PUBLISH", channelName, message)
}

func (a *Handler) readRedisChannelLoop(redisChannel *redis.PubSubConn, wsCon *websocket.Conn) {
	defer fmt.Println("end redis channel loop")
	for {
		switch v := redisChannel.Receive().(type) {
		case redis.Message:
			if wsCon.WriteMessage(websocket.TextMessage, v.Data) != nil {
				return
			}
		}
	}
}

func (a *Handler) updateStateMessageProcessing(userId uint, message *domain.UserPlayerUpdateStateMessage) error {
	switch message.TypePushState {
	case domain.PushTrackInQueue:
		return a.userSyncPlayerUseCase.PushTrackUpdateState(userId, message.Data.TracksQueue)
	case domain.NewTracksQueue:
		return a.userSyncPlayerUseCase.NewTrackQueueUpdateState(userId, message.Data.TracksQueue, message.Data.QueuePosition, message.Data.TimeStateUpdate)
	case domain.NewTrackInQueue:
		return a.userSyncPlayerUseCase.NewTrackUpdateState(userId, message.Data.QueuePosition, message.Data.TimeStateUpdate)
	case domain.OnPause:
		return a.userSyncPlayerUseCase.OnPauseUpdateState(userId, message.Data.TimeStateUpdate)
	case domain.OffPause:
		return a.userSyncPlayerUseCase.OffPauseUpdateState(userId, message.Data.TimeStateUpdate)
	case domain.ChangePosition:
		return a.userSyncPlayerUseCase.ChangePositionUpdateState(userId, message.Data.LastSecPosition, message.Data.TimeStateUpdate)
	}

	return errors.New("no such push state type")
}

func (a *Handler) PlayerStateLoop(c echo.Context) error {
	var upgrader = websocket.Upgrader{}
	userId, err := a.getUserId(c)

	if err != nil {
		return err
	}

	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	wsCon, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		fmt.Println("err = ", err)
		return err
	}
	defer wsCon.Close()

	// первоначально мы либо получаем текущее состояние плеера и отсылаем пользователю, либо отсылаем сообщение
	// об отсутствии состояния. в таком случае мы ожидаем, что придет сообщение с обновлением состояния
	trackState, err := a.userSyncPlayerUseCase.GetTrackState(userId)
	var messageState []byte
	if err != nil {
		messageState, _ = json.Marshal(getNoTrackStateMessage())
	} else {
		messageState, _ = json.Marshal(getTrackStateMessage(trackState))
	}

	err = wsCon.WriteMessage(websocket.TextMessage, messageState)
	if err != nil {
		return err
	}

	redisCon, redisPubSub, err := a.initRedisStructs(userId)
	if err != nil {
		return err
	}
	defer redisCon.Close()
	defer redisPubSub.Close()

	// запускаем бесконечный цикл, в котором будут читаться сообщения из redis channel'а и отправляться клиенту
	go a.readRedisChannelLoop(redisPubSub, wsCon)

	var clientMessage domain.UserPlayerUpdateStateMessage
	redisChannelName := getRedisChannelName(userId)
	for {
		_, message, err := wsCon.ReadMessage()

		if err != nil {
			break
		}

		err = json.Unmarshal(message, &clientMessage)
		if err != nil {
			messageState, _ = json.Marshal(getInvalidTrackStateFormatMessage())
			if wsCon.WriteMessage(websocket.TextMessage, messageState) != nil {
				break
			}
		} else {
			// обновляем состояние плеера
			err = a.updateStateMessageProcessing(userId, &clientMessage)
			if err == domain.ErrGetUserPlayerState {
				messageState, _ = json.Marshal(getNoTrackStateMessage())
				if wsCon.WriteMessage(websocket.TextMessage, messageState) != nil {
					break
				}
			} else if err != nil {
				messageState, _ = json.Marshal(getInvalidTrackStateFormatMessage())
				if wsCon.WriteMessage(websocket.TextMessage, messageState) != nil {
					break
				}
			} else {
				// публикуем обновление состояния плеера в redis channel. его считают другие клиенты
				a.pushToRedisChannel(redisCon, redisChannelName, string(messageState))
			}
		}
	}

	return nil
}
