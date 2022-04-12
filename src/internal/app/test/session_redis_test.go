package test

import (
	"github.com/alicebob/miniredis/v2"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/app/session/repository/redis"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetSession(t *testing.T) {
	miniRedis := miniredis.RunT(t)
	redisSessionRepo := redis.NewRedisSessionRepo(miniRedis.Addr())

	sessionId, err := redisSessionRepo.SetNewUnauthorizedSession(time.Hour * 10)
	assert.Nil(t, err)

	session, err := redisSessionRepo.GetSession(sessionId)
	assert.Nil(t, err)
	assert.False(t, session.IsAuthorized)

	session, err = redisSessionRepo.GetSession(sessionId + "a")
	assert.Equal(t, domain.ErrGetSession, err)
	assert.Nil(t, session)
}

func TestSetNewUnauthorizedSession(t *testing.T) {
	miniRedis := miniredis.RunT(t)
	redisSessionRepo := redis.NewRedisSessionRepo(miniRedis.Addr())

	sessionId, err := redisSessionRepo.SetNewUnauthorizedSession(time.Hour * 10)
	assert.Nil(t, err)

	session, err := redisSessionRepo.GetSession(sessionId)
	assert.Nil(t, err)
	assert.False(t, session.IsAuthorized)
}

func TestMakeSessionAuthorized(t *testing.T) {
	miniRedis := miniredis.RunT(t)
	redisSessionRepo := redis.NewRedisSessionRepo(miniRedis.Addr())

	sessionId, err := redisSessionRepo.SetNewUnauthorizedSession(time.Hour * 10)
	assert.Nil(t, err)

	err = redisSessionRepo.MakeSessionAuthorized(sessionId, 1)
	assert.Nil(t, err)

	session, err := redisSessionRepo.GetSession(sessionId)
	assert.Nil(t, err)
	assert.True(t, session.IsAuthorized)
}

func TestMakeSessionUnauthorized(t *testing.T) {
	miniRedis := miniredis.RunT(t)
	redisSessionRepo := redis.NewRedisSessionRepo(miniRedis.Addr())

	sessionId, err := redisSessionRepo.SetNewUnauthorizedSession(time.Hour * 10)
	assert.Nil(t, err)

	err = redisSessionRepo.MakeSessionAuthorized(sessionId, 1)
	assert.Nil(t, err)

	session, err := redisSessionRepo.GetSession(sessionId)
	assert.Nil(t, err)
	assert.True(t, session.IsAuthorized)

	err = redisSessionRepo.MakeSessionUnauthorized(sessionId)

	session, err = redisSessionRepo.GetSession(sessionId)
	assert.Nil(t, err)
	assert.False(t, session.IsAuthorized)
}

func TestDeleteSession(t *testing.T) {
	miniRedis := miniredis.RunT(t)
	redisSessionRepo := redis.NewRedisSessionRepo(miniRedis.Addr())

	sessionId, err := redisSessionRepo.SetNewUnauthorizedSession(time.Hour * 10)
	assert.Nil(t, err)

	err = redisSessionRepo.DeleteSession(sessionId)
	assert.Nil(t, err)

	session, err := redisSessionRepo.GetSession(sessionId)
	assert.NotNil(t, err)
	assert.Nil(t, session)
}
