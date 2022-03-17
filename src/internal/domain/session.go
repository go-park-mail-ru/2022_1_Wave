package domain

import "time"

type Session struct {
	UserId       uint
	IsAuthorized bool
	Expires      time.Time
}

type SessionRepo interface {
	GetSession(sessionId string) (*Session, error)
	SetNewUnauthorizedSession() (sessionId string, err error)
	SetNewSession(userId uint) (sessionId string, err error)
	DeleteSession(sessionId string) error
}
