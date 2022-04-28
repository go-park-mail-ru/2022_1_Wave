package auth_service

import (
	"context"
	"github.com/alicebob/miniredis/v2"
	auth_microservice_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/auth"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/auth/proto"
	auth_redis "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/auth/repository/redis"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var sessExpTest = time.Hour

func getAuthService(t *testing.T) proto.AuthorizationServer {
	miniRedis := miniredis.RunT(t)
	redisAuthRepo := auth_redis.NewRedisAuthRepo(miniRedis.Addr())
	service := NewAuthService(redisAuthRepo)

	return service
}

func TestGetSession(t *testing.T) {
	authService := getAuthService(t)

	unauth_sess := &proto.SetUnauthSession{Expires: sessExpTest.String()}
	sessionId, err := authService.SetNewUnauthorizedSession(context.Background(), unauth_sess)
	assert.Nil(t, err)

	session, err := authService.GetSession(context.Background(), sessionId)
	assert.Nil(t, err)
	assert.False(t, session.GetIsAuthorized())

	sessionId.SessionId += "a"
	session, err = authService.GetSession(context.Background(), sessionId)
	assert.Equal(t, auth_microservice_domain.ErrGetSession, err)
	assert.Nil(t, session)
}

func TestSetAuthorizedSession(t *testing.T) {
	authService := getAuthService(t)

	authSess := &proto.SetAuthSessionMsg{Expires: sessExpTest.String(), UserId: 1}
	sessionId, err := authService.SetNewAuthorizedSession(context.Background(), authSess)
	assert.Nil(t, err)

	session, err := authService.GetSession(context.Background(), sessionId)
	assert.Nil(t, err)
	assert.True(t, session.GetIsAuthorized())
	assert.Equal(t, session.GetUserId(), uint64(1))
}

func TestSetNewUnauthorizedSession(t *testing.T) {
	authService := getAuthService(t)

	unauth_sess := &proto.SetUnauthSession{Expires: sessExpTest.String()}
	sessionId, err := authService.SetNewUnauthorizedSession(context.Background(), unauth_sess)
	assert.Nil(t, err)

	session, err := authService.GetSession(context.Background(), sessionId)
	assert.Nil(t, err)
	assert.False(t, session.GetIsAuthorized())
}

func TestMakeSessionAuthorized(t *testing.T) {
	authService := getAuthService(t)

	unauth_sess := &proto.SetUnauthSession{Expires: sessExpTest.String()}
	sessionId, err := authService.SetNewUnauthorizedSession(context.Background(), unauth_sess)
	assert.Nil(t, err)

	userSessId := &proto.UserSessionId{SessionId: sessionId.GetSessionId(), UserId: 1}
	newSessionId, err := authService.MakeSessionAuthorized(context.Background(), userSessId)
	assert.Nil(t, err)
	assert.NotEqual(t, newSessionId.GetSessionId(), sessionId.GetSessionId())

	session, err := authService.GetSession(context.Background(), newSessionId)
	assert.Nil(t, err)
	assert.True(t, session.GetIsAuthorized())

	session, err = authService.GetSession(context.Background(), sessionId) // проверка что старая сессия удалилась
	assert.Nil(t, session)
	assert.NotNil(t, err)
}

func TestMakeSessionUnauthorized(t *testing.T) {
	authService := getAuthService(t)

	unauth_sess := &proto.SetUnauthSession{Expires: sessExpTest.String()}
	sessionId, err := authService.SetNewUnauthorizedSession(context.Background(), unauth_sess)
	assert.Nil(t, err)

	userSessId := &proto.UserSessionId{SessionId: sessionId.GetSessionId(), UserId: 1}
	newSessionId, err := authService.MakeSessionAuthorized(context.Background(), userSessId)
	assert.Nil(t, err)
	assert.NotEqual(t, newSessionId.GetSessionId(), sessionId.GetSessionId())

	session, err := authService.GetSession(context.Background(), newSessionId)
	assert.Nil(t, err)
	assert.True(t, session.GetIsAuthorized())

	session, err = authService.GetSession(context.Background(), sessionId) // проверка что старая сессия удалилась
	assert.Nil(t, session)
	assert.NotNil(t, err)

	newSessionId2, err := authService.MakeSessionUnauthorized(context.Background(), newSessionId)
	assert.Nil(t, err)
	assert.NotEqual(t, newSessionId.GetSessionId(), newSessionId2.GetSessionId())

	session, err = authService.GetSession(context.Background(), newSessionId2)
	assert.Nil(t, err)
	assert.False(t, session.GetIsAuthorized())

	session, err = authService.GetSession(context.Background(), newSessionId)
	assert.Nil(t, session)
	assert.NotNil(t, err)
}

func TestDeleteSession(t *testing.T) {
	authService := getAuthService(t)

	unauth_sess := &proto.SetUnauthSession{Expires: sessExpTest.String()}
	sessionId, err := authService.SetNewUnauthorizedSession(context.Background(), unauth_sess)
	assert.Nil(t, err)

	_, err = authService.DeleteSession(context.Background(), sessionId)
	assert.Nil(t, err)

	session, err := authService.GetSession(context.Background(), sessionId)
	assert.Nil(t, session)
	assert.NotNil(t, err)
}
