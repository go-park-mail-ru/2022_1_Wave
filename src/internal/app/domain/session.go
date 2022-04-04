package domain

import "time"

type Session struct {
	UserId       uint
	IsAuthorized bool
}

type SessionRepo interface {
	GetSession(sessionId string) (*Session, error)
	SetNewUnauthorizedSession(expires time.Duration) (sessionId string, err error)
	SetNewSession(expires time.Duration, userId uint) (sessionId string, err error)
	DeleteSession(sessionId string) error
}
