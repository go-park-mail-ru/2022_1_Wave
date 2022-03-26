package redis

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2022_1_Wave/config"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/domain"
	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
	"strconv"
	"strings"
	"time"
)

const (
	SessionsKey              = "sessions"
	UnauthorizedKey          = "unauthorized"
	UserIdSessionIdSeparator = "_"
)

func generateSessionId(isAuthorized bool, userId uint) string {
	var result string

	if isAuthorized {
		result = strconv.Itoa(int(userId)) + UserIdSessionIdSeparator + uuid.NewString()
	} else {
		result = UnauthorizedKey + UserIdSessionIdSeparator + uuid.NewString()
	}

	return result
}

func getUserIdFromSessionId(sessionId string) string {
	return strings.Split(sessionId, UserIdSessionIdSeparator)[0]
}

func getSessionHashTableName(userId string) string {
	return SessionsKey + ":" + userId
}

func setSession(client redis.Conn, sessionId string, isAuthorized bool, userId uint, expires time.Duration) error {
	/*
		sessions:<user_id>: {
			<session_id>: {...} - json,
			<session_id>: {...} - json,
			...
		} - по id пользователя хранятся его сессии
		при этом в session_id также зашит user_id, чтобы мы могли по session_id получить id пользователя
	*/
	var sessionHashTableName string
	if isAuthorized {
		sessionHashTableName = getSessionHashTableName(strconv.Itoa(int(userId)))
	} else {
		sessionHashTableName = getSessionHashTableName(UnauthorizedKey)
	}

	session := domain.Session{
		UserId:       userId,
		IsAuthorized: isAuthorized,
	}

	tableValue, _ := json.Marshal(session)

	_, err := client.Do("HSET", sessionHashTableName, sessionId, tableValue)
	if err != nil {
		return domain.ErrSetSession
	}

	_, err = client.Do("EXPIRE", sessionHashTableName, expires.Seconds())
	if err != nil {
		return domain.ErrSetSession
	}

	return nil
}

type redisSessionRepo struct {
	pool *redis.Pool
}

func NewRedisSessionRepo() domain.SessionRepo {
	return &redisSessionRepo{
		pool: &redis.Pool{
			Dial: func() (redis.Conn, error) {
				return redis.Dial("tcp", config.C.RedisAddress)
			},
		},
	}
}

func (a *redisSessionRepo) GetSession(sessionId string) (*domain.Session, error) {
	client := a.pool.Get()
	defer client.Close()

	userId := getUserIdFromSessionId(sessionId)
	// sessions:<user_id> -> <session_id> -> need session
	sessionHashTableName := getSessionHashTableName(userId)

	sessionJson, err := client.Do("HGET", sessionHashTableName, sessionId)
	if err != nil {
		return nil, domain.ErrGetSession
	}

	var session domain.Session

	_ = json.Unmarshal(sessionJson.([]byte), &session)

	return &session, nil
}

func (a *redisSessionRepo) SetNewUnauthorizedSession(expires time.Duration) (string, error) {
	client := a.pool.Get()
	defer client.Close()

	sessionId := generateSessionId(false, 0)

	err := setSession(client, sessionId, false, 0, expires)

	if err != nil {
		return "", domain.ErrSetSession
	}

	return sessionId, nil
}

func (a *redisSessionRepo) SetNewSession(expires time.Duration, userId uint) (string, error) {
	client := a.pool.Get()
	defer client.Close()

	sessionId := generateSessionId(true, userId)

	err := setSession(client, sessionId, true, userId, expires)

	if err != nil {
		return "", domain.ErrSetSession
	}

	return sessionId, nil
}

func (a *redisSessionRepo) DeleteSession(sessionId string) error {
	client := a.pool.Get()
	defer client.Close()

	userId := getUserIdFromSessionId(sessionId)
	// sessions:<user_id>
	sessionHashTableName := getSessionHashTableName(userId)

	_, err := client.Do("HDEL", sessionHashTableName, sessionId)
	if err != nil {
		return domain.ErrDeleteSession
	}

	// удалим таблицу сессий пользователя, если у него не осталось сессий
	result, err := client.Do("HLEN", sessionHashTableName)
	if err != nil {
		return domain.ErrDeleteSession
	}

	if result.(int) == 0 {
		_, err = client.Do("DEL", sessionHashTableName)
	}

	return err
}
