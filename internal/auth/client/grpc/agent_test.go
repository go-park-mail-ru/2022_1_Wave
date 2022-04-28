package auth_grpc_agent_test

import (
	"errors"
	auth_grpc_agent "github.com/go-park-mail-ru/2022_1_Wave/internal/auth/client/grpc"
	auth_microservice_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/auth"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/auth/mocks"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/auth/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestGetSession(t *testing.T) {
	sessionId := "some_sess_id"
	resultSession := auth_microservice_domain.Session{IsAuthorized: true, UserId: 1}
	protoResultSession := proto.Session{IsAuthorized: true, UserId: 1}
	authClientMock := mocks.AuthorizationClient{}

	t.Run("success", func(t *testing.T) {
		authClientMock.On("GetSession", mock.Anything, &proto.SessionId{SessionId: sessionId}).Return(&protoResultSession, nil)

		authAgent := auth_grpc_agent.NewAuthGRPCAgent(&authClientMock)
		sessResult, err := authAgent.GetSession(sessionId)
		assert.NoError(t, err)
		assert.Equal(t, *sessResult, resultSession)
	})

	t.Run("error", func(t *testing.T) {
		authClientMock.On("GetSession", mock.Anything, &proto.SessionId{SessionId: sessionId + "a"}).Return(nil, errors.New("error"))

		authAgent := auth_grpc_agent.NewAuthGRPCAgent(&authClientMock)
		sessResult, err := authAgent.GetSession(sessionId + "a")
		assert.Error(t, err)
		assert.Nil(t, sessResult)
	})
}

func TestSetNewUnauthorizedSession(t *testing.T) {
	sessionId := "some_sess_id"
	authClientMock := mocks.AuthorizationClient{}
	expires := time.Hour
	sessionIdProto := &proto.SessionId{SessionId: sessionId}

	t.Run("success", func(t *testing.T) {
		authClientMock.On("SetNewUnauthorizedSession", mock.Anything,
			&proto.SetUnauthSession{Expires: expires.String()}).Return(sessionIdProto, nil)

		authAgent := auth_grpc_agent.NewAuthGRPCAgent(&authClientMock)
		sessionIdResult, err := authAgent.SetNewUnauthorizedSession(expires)
		assert.NoError(t, err)
		assert.Equal(t, sessionIdResult, sessionId)
	})

	t.Run("error", func(t *testing.T) {
		authClientMock.On("SetNewUnauthorizedSession", mock.Anything,
			&proto.SetUnauthSession{Expires: (expires * 2).String()}).Return(nil, errors.New("error"))

		authAgent := auth_grpc_agent.NewAuthGRPCAgent(&authClientMock)
		sessionIdResult, err := authAgent.SetNewUnauthorizedSession(expires * 2)
		assert.Error(t, err)
		assert.Equal(t, sessionIdResult, "")
	})
}

func TestSetAuthorizedSession(t *testing.T) {
	sessionId := "some_sess_id"
	authClientMock := mocks.AuthorizationClient{}
	expires := time.Hour
	sessionIdProto := &proto.SessionId{SessionId: sessionId}

	t.Run("success", func(t *testing.T) {
		authClientMock.On("SetNewAuthorizedSession", mock.Anything,
			&proto.SetAuthSessionMsg{Expires: expires.String(), UserId: 1}).Return(sessionIdProto, nil)

		authAgent := auth_grpc_agent.NewAuthGRPCAgent(&authClientMock)
		sessionIdResult, err := authAgent.SetNewAuthorizedSession(1, expires)
		assert.NoError(t, err)
		assert.Equal(t, sessionIdResult, sessionId)
	})

	t.Run("error", func(t *testing.T) {
		authClientMock.On("SetNewAuthorizedSession", mock.Anything,
			&proto.SetAuthSessionMsg{Expires: expires.String(), UserId: 2}).Return(nil, errors.New("error"))

		authAgent := auth_grpc_agent.NewAuthGRPCAgent(&authClientMock)
		sessionIdResult, err := authAgent.SetNewAuthorizedSession(2, expires)
		assert.Error(t, err)
		assert.Equal(t, sessionIdResult, "")
	})
}

func TestMakeSessionAuthorized(t *testing.T) {
	sessionId := "some_sess_id"
	authClientMock := mocks.AuthorizationClient{}
	sessionIdProto := &proto.SessionId{SessionId: sessionId + "a"}

	t.Run("success", func(t *testing.T) {
		authClientMock.On("MakeSessionAuthorized", mock.Anything,
			&proto.UserSessionId{SessionId: sessionId, UserId: 1}).Return(sessionIdProto, nil)

		authAgent := auth_grpc_agent.NewAuthGRPCAgent(&authClientMock)
		sessionIdResult, err := authAgent.MakeSessionAuthorized(sessionId, 1)
		assert.NoError(t, err)
		assert.Equal(t, sessionIdResult, sessionId+"a")
	})
	t.Run("error", func(t *testing.T) {
		authClientMock.On("MakeSessionAuthorized", mock.Anything,
			&proto.UserSessionId{SessionId: sessionId + "a", UserId: 1}).Return(nil, errors.New("error"))

		authAgent := auth_grpc_agent.NewAuthGRPCAgent(&authClientMock)
		sessionIdResult, err := authAgent.MakeSessionAuthorized(sessionId+"a", 1)
		assert.Error(t, err)
		assert.Equal(t, sessionIdResult, "")
	})
}

func TestMakeSessionUnauthorized(t *testing.T) {
	sessionId := "some_sess_id"
	authClientMock := mocks.AuthorizationClient{}
	sessionIdProto := &proto.SessionId{SessionId: sessionId + "a"}

	t.Run("success", func(t *testing.T) {
		authClientMock.On("MakeSessionUnauthorized", mock.Anything,
			&proto.SessionId{SessionId: sessionId}).Return(sessionIdProto, nil)

		authAgent := auth_grpc_agent.NewAuthGRPCAgent(&authClientMock)
		sessionIdResult, err := authAgent.MakeSessionUnauthorized(sessionId)
		assert.NoError(t, err)
		assert.Equal(t, sessionIdResult, sessionId+"a")
	})
	t.Run("error", func(t *testing.T) {
		authClientMock.On("MakeSessionUnauthorized", mock.Anything,
			&proto.SessionId{SessionId: sessionId + "a"}).Return(nil, errors.New("error"))

		authAgent := auth_grpc_agent.NewAuthGRPCAgent(&authClientMock)
		sessionIdResult, err := authAgent.MakeSessionUnauthorized(sessionId + "a")
		assert.Error(t, err)
		assert.Equal(t, sessionIdResult, "")
	})
}

func TestIsSession(t *testing.T) {
	sessionId := "some_sess_id"
	authClientMock := mocks.AuthorizationClient{}
	boolResult := &proto.BoolResult{Result: true}

	t.Run("success", func(t *testing.T) {
		authClientMock.On("IsSession", mock.Anything,
			&proto.SessionId{SessionId: sessionId}).Return(boolResult, nil)

		authAgent := auth_grpc_agent.NewAuthGRPCAgent(&authClientMock)
		isSessionResult := authAgent.IsSession(sessionId)
		assert.True(t, isSessionResult)
	})
}

func TestIsAuthSession(t *testing.T) {
	sessionId := "some_sess_id"
	authClientMock := mocks.AuthorizationClient{}
	boolResult := &proto.BoolResult{Result: true}

	t.Run("success", func(t *testing.T) {
		authClientMock.On("IsAuthSession", mock.Anything,
			&proto.SessionId{SessionId: sessionId}).Return(boolResult, nil)

		authAgent := auth_grpc_agent.NewAuthGRPCAgent(&authClientMock)
		isSessionResult := authAgent.IsAuthSession(sessionId)
		assert.True(t, isSessionResult)
	})
}
