package auth_redis_test

import (
	"github.com/alicebob/miniredis/v2"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	auth_redis "github.com/go-park-mail-ru/2022_1_Wave/internal/app/microservices/auth/repository/redis"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetSession(t *testing.T) {
	miniRedis := miniredis.RunT(t)
	redisAuthRepo := auth_redis.NewRedisAuthRepo(miniRedis.Addr())

	sessionId, err := redisAuthRepo.SetNewUnauthorizedSession(time.Hour * 10)
	assert.Nil(t, err)

	session, err := redisAuthRepo.GetSession(sessionId)
	assert.Nil(t, err)
	assert.False(t, session.IsAuthorized)

	session, err = redisAuthRepo.GetSession(sessionId + "a")
	assert.Equal(t, domain.ErrGetSession, err)
	assert.Nil(t, session)
}

func TestSetNewUnauthorizedSession(t *testing.T) {
	miniRedis := miniredis.RunT(t)
	redisAuthRepo := auth_redis.NewRedisAuthRepo(miniRedis.Addr())

	sessionId, err := redisAuthRepo.SetNewUnauthorizedSession(time.Hour * 10)
	assert.Nil(t, err)

	session, err := redisAuthRepo.GetSession(sessionId)
	assert.Nil(t, err)
	assert.False(t, session.IsAuthorized)
}

func TestMakeSessionAuthorized(t *testing.T) {
	miniRedis := miniredis.RunT(t)
	redisAuthRepo := auth_redis.NewRedisAuthRepo(miniRedis.Addr())

	sessionId, err := redisAuthRepo.SetNewUnauthorizedSession(time.Hour * 10)
	assert.Nil(t, err)

	newSessionId, err := redisAuthRepo.MakeSessionAuthorized(sessionId, 1)
	assert.Nil(t, err)
	assert.NotEqual(t, newSessionId, sessionId)

	session, err := redisAuthRepo.GetSession(newSessionId)
	assert.Nil(t, err)
	assert.True(t, session.IsAuthorized)

	session, err = redisAuthRepo.GetSession(sessionId) // проверка что старая сессия удалилась
	assert.Nil(t, session)
	assert.NotNil(t, err)
}

func TestMakeSessionUnauthorized(t *testing.T) {
	miniRedis := miniredis.RunT(t)
	redisAuthRepo := auth_redis.NewRedisAuthRepo(miniRedis.Addr())

	sessionId, err := redisAuthRepo.SetNewUnauthorizedSession(time.Hour * 10)
	assert.Nil(t, err)

	newSessionId1, err := redisAuthRepo.MakeSessionAuthorized(sessionId, 1)
	assert.Nil(t, err)
	assert.NotEqual(t, newSessionId1, sessionId)

	session, err := redisAuthRepo.GetSession(newSessionId1)
	assert.Nil(t, err)
	assert.True(t, session.IsAuthorized)

	session, err = redisAuthRepo.GetSession(sessionId) // проверка что старая сессия удалилась
	assert.Nil(t, session)
	assert.NotNil(t, err)

	newSessionId2, err := redisAuthRepo.MakeSessionUnauthorized(newSessionId1)
	assert.Nil(t, err)
	assert.NotEqual(t, newSessionId2, newSessionId1)

	session, err = redisAuthRepo.GetSession(newSessionId2)
	assert.Nil(t, err)
	assert.False(t, session.IsAuthorized)

	session, err = redisAuthRepo.GetSession(newSessionId1) // проверка что старая сессия удалилась
	assert.Nil(t, session)
	assert.NotNil(t, err)
}

func TestDeleteSession(t *testing.T) {
	miniRedis := miniredis.RunT(t)
	redisAuthRepo := auth_redis.NewRedisAuthRepo(miniRedis.Addr())

	sessionId, err := redisAuthRepo.SetNewUnauthorizedSession(time.Hour * 10)
	assert.Nil(t, err)

	err = redisAuthRepo.DeleteSession(sessionId)
	assert.Nil(t, err)

	session, err := redisAuthRepo.GetSession(sessionId)
	assert.NotNil(t, err)
	assert.Nil(t, session)
}
