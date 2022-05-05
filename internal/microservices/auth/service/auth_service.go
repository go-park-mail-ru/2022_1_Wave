package auth_service

import (
	"context"
	auth_microservice_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/auth"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/auth/proto"
	"time"
)

type authService struct {
	authRepo auth_microservice_domain.AuthRepo
	proto.UnsafeAuthorizationServer
}

/*
type AuthorizationServer interface {
	GetSession(context.Context, *SessionId) (*Session, error)
	SetNewUnauthorizedSession(context.Context, *Empty) (*SessionId, error)
	SetNewAuthorizedSession(context.Context, *UserId) (*SessionId, error)
	MakeSessionAuthorized(context.Context, *UserSessionId) (*SessionId, error)
	MakeSessionUnauthorized(context.Context, *SessionId) (*SessionId, error)
	DeleteSession(context.Context, *SessionId) (*Empty, error)
	IsSession(context.Context, *SessionId) (*Empty, error)
	IsAuthSession(context.Context, *SessionId) (*BoolResult, error)
}
*/

func NewAuthService(authRepo auth_microservice_domain.AuthRepo) proto.AuthorizationServer {
	return &authService{authRepo: authRepo}
}

func (a *authService) GetSession(ctx context.Context, sessionId *proto.SessionId) (*proto.Session, error) {
	session, err := a.authRepo.GetSession(sessionId.GetSessionId())
	if err != nil {
		return nil, err
	}

	var sessionProto proto.Session
	sessionProto.UserId = uint64(session.UserId)
	sessionProto.IsAuthorized = session.IsAuthorized

	return &sessionProto, nil
}

func (a *authService) SetNewUnauthorizedSession(ctx context.Context, data *proto.SetUnauthSession) (*proto.SessionId, error) {
	expires, _ := time.ParseDuration(data.GetExpires())
	sessionId, err := a.authRepo.SetNewUnauthorizedSession(expires)
	if err != nil {
		return nil, err
	}

	var sessionIdProto proto.SessionId
	sessionIdProto.SessionId = sessionId

	return &sessionIdProto, nil
}

func (a *authService) SetNewAuthorizedSession(ctx context.Context, data *proto.SetAuthSessionMsg) (*proto.SessionId, error) {
	expires, _ := time.ParseDuration(data.GetExpires())
	sessionId, err := a.authRepo.SetNewAuthorizedSession(uint(data.GetUserId()), expires)
	if err != nil {
		return nil, err
	}

	var sessionIdProto proto.SessionId
	sessionIdProto.SessionId = sessionId

	return &sessionIdProto, nil
}

func (a *authService) MakeSessionAuthorized(ctx context.Context, userAndSessionIds *proto.UserSessionId) (*proto.SessionId, error) {
	newSessionId, err := a.authRepo.MakeSessionAuthorized(userAndSessionIds.GetSessionId(), uint(userAndSessionIds.GetUserId()))
	if err != nil {
		return nil, err
	}

	var sessionIdProto proto.SessionId
	sessionIdProto.SessionId = newSessionId

	return &sessionIdProto, nil
}

func (a *authService) MakeSessionUnauthorized(ctx context.Context, sessionId *proto.SessionId) (*proto.SessionId, error) {
	newSessionId, err := a.authRepo.MakeSessionUnauthorized(sessionId.GetSessionId())
	if err != nil {
		return nil, err
	}

	var sessionIdProto proto.SessionId
	sessionIdProto.SessionId = newSessionId

	return &sessionIdProto, nil
}

func (a *authService) DeleteSession(ctx context.Context, sessionId *proto.SessionId) (*proto.Empty, error) {
	err := a.authRepo.DeleteSession(sessionId.GetSessionId())

	return &proto.Empty{}, err
}

func (a *authService) IsSession(ctx context.Context, sessionId *proto.SessionId) (*proto.BoolResult, error) {
	_, err := a.authRepo.GetSession(sessionId.GetSessionId())

	var boolResult proto.BoolResult
	boolResult.Result = err == nil

	return &boolResult, err
}

func (a *authService) IsAuthSession(ctx context.Context, sessionId *proto.SessionId) (*proto.BoolResult, error) {
	session, err := a.authRepo.GetSession(sessionId.GetSessionId())

	var boolResult proto.BoolResult
	boolResult.Result = err == nil && session != nil && session.IsAuthorized

	return &boolResult, err
}
