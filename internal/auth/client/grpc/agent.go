package auth_grpc_agent

import (
	"context"
	auth_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/auth"
	auth_microservice_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/auth"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/auth/proto"
	"time"
)

type authGRPCAgent struct {
	authClient proto.AuthorizationClient
}

func NewAuthGRPCAgent(authClient proto.AuthorizationClient) auth_domain.AuthAgent {
	return &authGRPCAgent{authClient: authClient}
}

func (a *authGRPCAgent) GetSession(sessionId string) (*auth_microservice_domain.Session, error) {
	session, err := a.authClient.GetSession(context.Background(), &proto.SessionId{SessionId: sessionId})

	if err != nil {
		return nil, err
	}

	var sessionResult auth_microservice_domain.Session
	sessionResult.UserId = uint(session.GetUserId())
	sessionResult.IsAuthorized = session.GetIsAuthorized()

	return &sessionResult, err
}

func (a *authGRPCAgent) SetNewUnauthorizedSession(expires time.Duration) (string, error) {
	sessionId, err := a.authClient.SetNewUnauthorizedSession(context.Background(), &proto.SetUnauthSession{Expires: expires.String()})
	if err != nil {
		return "", err
	}

	return sessionId.GetSessionId(), nil
}

func (a *authGRPCAgent) SetNewAuthorizedSession(userId uint, expires time.Duration) (string, error) {
	sessionId, err := a.authClient.SetNewAuthorizedSession(context.Background(),
		&proto.SetAuthSessionMsg{
			UserId:  uint64(userId),
			Expires: expires.String(),
		})

	if err != nil {
		return "", err
	}

	return sessionId.GetSessionId(), nil
}

func (a *authGRPCAgent) MakeSessionAuthorized(sessionId string, userId uint) (string, error) {
	newSessionId, err := a.authClient.MakeSessionAuthorized(context.Background(),
		&proto.UserSessionId{
			SessionId: sessionId,
			UserId:    uint64(userId),
		})

	if err != nil {
		return "", err
	}

	return newSessionId.GetSessionId(), nil
}

func (a *authGRPCAgent) MakeSessionUnauthorized(sessionId string) (string, error) {
	newSessionId, err := a.authClient.MakeSessionUnauthorized(context.Background(),
		&proto.SessionId{SessionId: sessionId})

	if err != nil {
		return "", err
	}

	return newSessionId.GetSessionId(), nil
}

func (a *authGRPCAgent) DeleteSession(sessionId string) error {
	_, err := a.authClient.DeleteSession(context.Background(), &proto.SessionId{SessionId: sessionId})

	return err
}

func (a *authGRPCAgent) IsSession(sessionId string) bool {
	boolResult, _ := a.authClient.IsSession(context.Background(), &proto.SessionId{SessionId: sessionId})

	return boolResult.GetResult()
}

func (a *authGRPCAgent) IsAuthSession(sessionId string) bool {
	boolResult, _ := a.authClient.IsAuthSession(context.Background(), &proto.SessionId{SessionId: sessionId})

	return boolResult.GetResult()
}
