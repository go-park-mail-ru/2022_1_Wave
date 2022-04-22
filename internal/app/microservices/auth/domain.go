package auth_domain

import (
	"time"
)

type Session struct {
	UserId       uint `json:"user_id"`
	IsAuthorized bool `json:"is_authorized"`
}

type AuthRepo interface {
	GetSession(sessionId string) (*Session, error)
	SetNewUnauthorizedSession(expires time.Duration) (sessionId string, err error)
	MakeSessionAuthorized(sessionId string, userId uint) (newSessionId string, err error)
	MakeSessionUnauthorized(sessionId string) (newSessionId string, err error)
	DeleteSession(sessionId string) error
	GetSize() (int, error)
}

const (
	ErrGetSession     = "error while get session"
	ErrSetSession     = "error while set session"
	ErrDeleteSession  = "error while delete session"
	ErrNotAuthSession = "session is not authorized"

	ErrInvalidLoginOrPassword = "invalid login or password"
	ErrUserAlreadyExist       = "user already exist"
	ErrUserDoesNotExist       = "user does not exist"

	ErrWhileSetNewSession       = "error while set new session"
	ErrWhileChangeSession       = "error while change session"
	ErrSessionStorageUnexpected = "session storage unexpected error"
	ErrSessionDoesNotExist      = "session does not exist"

	ErrInsert             = "insertion error"
	ErrDatabaseUnexpected = "database unexpected error"
)
