package auth_redis

import (
	"encoding/json"
	"errors"
	auth_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/auth"
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

	session := auth_domain.Session{
		UserId:       userId,
		IsAuthorized: isAuthorized,
	}

	tableValue, _ := json.Marshal(session)

	_, err := client.Do("HSET", sessionHashTableName, sessionId, tableValue)
	if err != nil {
		return errors.New(auth_domain.ErrSetSession)
	}

	_, err = client.Do("EXPIRE", sessionHashTableName, expires.Seconds())
	if err != nil {
		return errors.New(auth_domain.ErrSetSession)
	}

	return nil
}

type redisAuthRepo struct {
	pool *redis.Pool
}

func NewRedisAuthRepo(redisAddr string) auth_domain.AuthRepo {
	return &redisAuthRepo{
		pool: &redis.Pool{
			Dial: func() (redis.Conn, error) {
				return redis.Dial("tcp", redisAddr)
			},
		},
	}
}

func (a *redisAuthRepo) GetSession(sessionId string) (*auth_domain.Session, error) {
	client := a.pool.Get()
	defer client.Close()

	userId := getUserIdFromSessionId(sessionId)
	// sessions:<user_id> -> <session_id> -> need session
	sessionHashTableName := getSessionHashTableName(userId)

	sessionJson, err := client.Do("HGET", sessionHashTableName, sessionId)
	if err != nil || sessionJson == nil {
		return nil, errors.New(auth_domain.ErrGetSession)
	}

	var session auth_domain.Session

	_ = json.Unmarshal(sessionJson.([]byte), &session)

	return &session, nil
}

func (a *redisAuthRepo) SetNewAuthorizedSession(userId uint, expires time.Duration) (string, error) {
	client := a.pool.Get()
	defer client.Close()

	sessionId := generateSessionId(true, userId)

	err := setSession(client, sessionId, true, userId, expires)

	if err != nil {
		return "", errors.New(auth_domain.ErrSetSession)
	}

	return sessionId, nil
}

func (a *redisAuthRepo) SetNewUnauthorizedSession(expires time.Duration) (string, error) {
	client := a.pool.Get()
	defer client.Close()

	sessionId := generateSessionId(false, 0)

	err := setSession(client, sessionId, false, 0, expires)

	if err != nil {
		return "", errors.New(auth_domain.ErrSetSession)
	}

	return sessionId, nil
}

func (a *redisAuthRepo) changeSession(sessionId string, isAuthorized bool, userId uint) (string, error) {
	session, err := a.GetSession(sessionId)
	if err != nil {
		return "", errors.New(auth_domain.ErrGetSession)
	}
	session.IsAuthorized = isAuthorized
	session.UserId = userId
	tableValue, _ := json.Marshal(session)

	client := a.pool.Get()
	defer client.Close()

	// sessions:<user_id> -> <session_id> -> need session
	var sessionHashTableName string
	newSessionId := generateSessionId(isAuthorized, userId)
	if isAuthorized {
		sessionHashTableName = getSessionHashTableName(strconv.Itoa(int(userId)))
	} else {
		sessionHashTableName = getSessionHashTableName(UnauthorizedKey)
	}

	_, err = client.Do("HSET", sessionHashTableName, newSessionId, tableValue)
	if err != nil {
		return "", errors.New(auth_domain.ErrSetSession)
	}

	_ = a.DeleteSession(sessionId) // удаляем старую версию сессии

	return newSessionId, nil
}

func (a *redisAuthRepo) MakeSessionAuthorized(sessionId string, userId uint) (string, error) {
	return a.changeSession(sessionId, true, userId)
}

func (a *redisAuthRepo) MakeSessionUnauthorized(sessionId string) (string, error) {
	return a.changeSession(sessionId, false, 0)
}

func (a *redisAuthRepo) DeleteSession(sessionId string) error {
	client := a.pool.Get()
	defer client.Close()

	userId := getUserIdFromSessionId(sessionId)
	// sessions:<user_id>
	sessionHashTableName := getSessionHashTableName(userId)

	_, err := client.Do("HDEL", sessionHashTableName, sessionId)
	if err != nil {
		return errors.New(auth_domain.ErrDeleteSession)
	}

	// удалим таблицу сессий пользователя, если у него не осталось сессий
	result, err := client.Do("HLEN", sessionHashTableName)
	if err != nil || result == nil {
		return errors.New(auth_domain.ErrDeleteSession)
	}

	if result == 0 {
		_, err = client.Do("DEL", sessionHashTableName)
	}

	return err
}

func (a *redisAuthRepo) GetSize() (int, error) {
	var size int
	return size, nil
}
