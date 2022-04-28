package auth_domain

import (
	"errors"
	auth_microservice_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/auth"
	user_microservice_domain "github.com/go-park-mail-ru/2022_1_Wave/internal/microservices/user"
	"time"
)

type AuthAgent interface {
	GetSession(sessionId string) (*auth_microservice_domain.Session, error)
	SetNewUnauthorizedSession(expires time.Duration) (sessionId string, err error)
	SetNewAuthorizedSession(userId uint, expires time.Duration) (sessionId string, err error)
	MakeSessionAuthorized(sessionId string, userId uint) (newSessionId string, err error)
	MakeSessionUnauthorized(sessionId string) (newSessionId string, err error)
	DeleteSession(sessionId string) error
	IsSession(sessionId string) bool
	IsAuthSession(sessionId string) bool
}

type AuthUseCase interface {
	Login(login string, password string) (sessionId string, err error)
	Logout(sessionId string) error
	SignUp(user *user_microservice_domain.User) (sessionId string, err error)
	GetUnauthorizedSession() (sessionId string, err error)
	IsSession(sessionId string) bool
	IsAuthSession(sessionId string) bool
}

var (
	ErrSetSession    = errors.New("error while set session")
	ErrDeleteSession = errors.New("error while delete session")

	ErrInvalidLoginOrPassword = errors.New("invalid login or password")
	ErrUserAlreadyExist       = errors.New("user already exist")

	ErrDatabaseUnexpected = errors.New("database unexpected error")
)
