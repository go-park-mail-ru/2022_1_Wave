package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	auth_domain "github.com/go-park-mail-ru/2022_1_Wave/websocket-server/auth"
	"github.com/go-park-mail-ru/2022_1_Wave/websocket-server/domain"
	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
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

func (a *Handler) initRedisStructs(userId uint) (redisCon redis.Conn, redisPubSub *redis.PubSubConn, uid string, err error) {
	redisCon, err = redis.Dial("tcp", a.redisAddr)
	if err != nil {
		return
	}

	redisPubSub = &redis.PubSubConn{Conn: redisCon}
	err = redisPubSub.Subscribe(getRedisChannelName(userId))
	uid = uuid.NewString()
	return
}

func (a *Handler) initRedisConn() (redis.Conn, error) {
	return redis.Dial("tcp", a.redisAddr)
}

func (a *Handler) pushToRedisChannel(redisCon redis.Conn, channelName string, message string, uuid string) {
	fmt.Println("pub message ", message, " to redis channel ", channelName, " from ", redisCon)
	msg, err := redisCon.Do("PUBLISH", channelName, uuid+message)
	fmt.Println("redis publish msg = ", msg) // TODO: для публикации нужно создать другое подключение
	fmt.Println("redis publish err = ", err)
}

func getUuidAndMessageFromMessage(msg string) (string, string) {
	uid := msg[:36]
	msg = msg[36:len(msg)]

	return msg, uid
}

func (a *Handler) readRedisChannelLoop(redisChannel *redis.PubSubConn, wsCon *websocket.Conn, ctx context.Context, uid string) {
	defer fmt.Println("end redis channel loop")
	for {
		fmt.Println("in redis channel loop")
		switch v := redisChannel.Receive().(type) {
		case redis.Message:
			fmt.Println("message from ", wsCon, ":")
			dataStr := string(v.Data)
			msg, uidRedisCon := getUuidAndMessageFromMessage(dataStr)
			if uidRedisCon == uid {
				continue
			}
			if wsCon.WriteMessage(websocket.TextMessage, []byte(msg)) != nil {
				return
			}
		case redis.Error:
			fmt.Println("redis error")
			return
		default:
			select {
			case <-ctx.Done():
				fmt.Println("ctx done")
				return
			default:
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
		return err
	}
	defer wsCon.Close()
	defer fmt.Println("end loop with websocket")

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

	redisCon, redisPubSub, uidRedisCon, err := a.initRedisStructs(userId)
	fmt.Println("uid redis con len = ", len(uidRedisCon))
	if err != nil {
		return err
	}
	defer redisCon.Close()
	defer redisPubSub.Close()

	// запускаем бесконечный цикл, в котором будут читаться сообщения из redis channel'а и отправляться клиенту
	readRedisLoopCtx := context.Background()
	go a.readRedisChannelLoop(redisPubSub, wsCon, readRedisLoopCtx, uidRedisCon)
	defer readRedisLoopCtx.Done()

	redisConForPublish, err := a.initRedisConn()
	if err != nil {
		return err
	}
	defer redisConForPublish.Close()

	var clientMessage domain.UserPlayerUpdateStateMessage
	redisChannelName := getRedisChannelName(userId)
	for {
		_, message, err := wsCon.ReadMessage()

		if err != nil {
			fmt.Println("break 0, err = ", err)
			break
		}

		err = json.Unmarshal(message, &clientMessage)
		if err != nil {
			messageState, _ = json.Marshal(getInvalidTrackStateFormatMessage())
			if wsCon.WriteMessage(websocket.TextMessage, messageState) != nil {
				fmt.Println("break 1")
				break
			}
		} else {
			// обновляем состояние плеера
			err = a.updateStateMessageProcessing(userId, &clientMessage)
			if err == domain.ErrGetUserPlayerState {
				messageState, _ = json.Marshal(getNoTrackStateMessage())
				if wsCon.WriteMessage(websocket.TextMessage, messageState) != nil {
					fmt.Println("break 2")
					break
				}
			} else if err != nil {
				messageState, _ = json.Marshal(getInvalidTrackStateFormatMessage())
				if wsCon.WriteMessage(websocket.TextMessage, messageState) != nil {
					fmt.Println("break 3")
					break
				}
			} else {
				// публикуем обновление состояния плеера в redis channel. его считают другие клиенты
				a.pushToRedisChannel(redisConForPublish, redisChannelName, string(messageState), uidRedisCon)
			}
		}
	}

	return nil
}
