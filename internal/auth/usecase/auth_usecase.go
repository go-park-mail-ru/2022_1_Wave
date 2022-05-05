package auth_usecase

import (
	auth_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/auth"
	user_microservice_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/user"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/tools/utils"
	user_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/user"
	"time"
)

type authService struct {
	authAgent auth_domain.AuthAgent
	userAgent user_domain.UserAgent
}

var SessionExpires = time.Hour * 24

func NewAuthService(authAgent auth_domain.AuthAgent, userAgent user_domain.UserAgent) auth_domain.AuthUseCase {
	return &authService{authAgent: authAgent, userAgent: userAgent}
}

func (a *authService) Login(login string, password string) (string, error) {
	user, err := a.userAgent.GetByUsername(login)

	if err != nil {
		user, err = a.userAgent.GetByEmail(login)
		if err != nil {
			return "", err
		}
	}

	if !utils.CheckPassword(user.Password, password) {
		return "", auth_domain.ErrInvalidLoginOrPassword
	}

	sessionId, err := a.authAgent.SetNewAuthorizedSession(user.ID, SessionExpires)

	if err != nil {
		return "", auth_domain.ErrSetSession
	}

	return sessionId, nil
}

func (a *authService) Logout(sessionId string) error {
	err := a.authAgent.DeleteSession(sessionId)

	if err != nil {
		return auth_domain.ErrDeleteSession
	}

	return nil
}

func (a *authService) SignUp(user *user_microservice_domain.User) (string, error) {
	_, err := a.userAgent.GetByEmail(user.Email)
	if err == nil {
		return "", auth_domain.ErrUserAlreadyExist
	}

	_, err = a.userAgent.GetByUsername(user.Username)
	if err == nil {
		return "", auth_domain.ErrUserAlreadyExist
	}

	passwordHash, _ := utils.GetPasswordHash(user.Password)

	user.Password = string(passwordHash)

	err = a.userAgent.Create(user)
	if err != nil {
		return "", auth_domain.ErrDatabaseUnexpected
	}

	userToId, err := a.userAgent.GetByEmail(user.Email)
	if err != nil {
		return "", auth_domain.ErrDatabaseUnexpected
	}

	sessionId, err := a.authAgent.SetNewAuthorizedSession(userToId.ID, SessionExpires)

	if err != nil {
		return "", auth_domain.ErrSetSession
	}

	return sessionId, nil
}

func (a *authService) GetUnauthorizedSession() (string, error) {
	sessionId, err := a.authAgent.SetNewUnauthorizedSession(SessionExpires)
	if err != nil {
		return "", auth_domain.ErrSetSession
	}

	return sessionId, nil
}

func (a *authService) IsSession(sessionId string) bool {
	return a.authAgent.IsSession(sessionId)
}

func (a *authService) IsAuthSession(sessionId string) bool {
	return a.authAgent.IsAuthSession(sessionId)
}
