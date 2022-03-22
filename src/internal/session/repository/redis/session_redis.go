package redis

import (
	"github.com/go-park-mail-ru/2022_1_Wave/config"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/domain"
	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
	"time"
)

const (
	SessionsKey             = "sessions"
	UnauthorizedIntSign     = 0
	AuthorizedIntSign       = 1
	SessionsUserIdKey       = "user_id"
	SessionsIsAuthorizedKey = "is_authorized"
)

func generateSessionId() string {
	return uuid.NewString()
}

func getSessionHashTableName(sessionId string) string {
	return SessionsKey + ":" + sessionId
}

func setSession(client redis.Conn, sessionId string, authorizedSign uint, userId uint, expires time.Duration) error {
	// sessions:<session_id>
	sessionHashTableName := getSessionHashTableName(sessionId)

	_, err := client.Do("HSET", sessionHashTableName, SessionsIsAuthorizedKey, authorizedSign)
	if err != nil {
		return domain.ErrSetSession
	}

	_, err = client.Do("HSET", sessionHashTableName, SessionsUserIdKey, userId)
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

	// sessions:<session_id>
	sessionHashTableName := getSessionHashTableName(sessionId)

	userId, err := client.Do("HGET", sessionHashTableName, SessionsUserIdKey)
	if err != nil {
		return nil, domain.ErrGetSession
	}

	isAuthorized, err := client.Do("HGET", sessionHashTableName, SessionsIsAuthorizedKey)
	if err != nil {
		return nil, domain.ErrGetSession
	}

	var isAuthorizedBool bool
	if isAuthorized.(int) == UnauthorizedIntSign {
		isAuthorizedBool = false
	} else {
		isAuthorizedBool = true
	}

	return &domain.Session{
		UserId:       userId.(uint),
		IsAuthorized: isAuthorizedBool,
	}, nil
}

func (a *redisSessionRepo) SetNewUnauthorizedSession(expires time.Duration) (string, error) {
	client := a.pool.Get()
	defer client.Close()

	sessionId := generateSessionId()

	err := setSession(client, sessionId, UnauthorizedIntSign, 0, expires)

	if err != nil {
		return "", domain.ErrSetSession
	}

	return sessionId, nil
}

func (a *redisSessionRepo) SetNewSession(expires time.Duration, userId uint) (string, error) {
	client := a.pool.Get()
	defer client.Close()

	sessionId := generateSessionId()

	err := setSession(client, sessionId, AuthorizedIntSign, userId, expires)

	if err != nil {
		return "", domain.ErrSetSession
	}

	return sessionId, nil
}

func (a *redisSessionRepo) DeleteSession(sessionId string) error {
	client := a.pool.Get()
	defer client.Close()

	// sessions:<session_id>
	sessionHashTableName := getSessionHashTableName(sessionId)

	_, err := client.Do("DEL", sessionHashTableName)
	if err != nil {
		return domain.ErrDeleteSession
	}

	return nil
}
