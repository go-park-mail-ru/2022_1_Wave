package domain

import "time"

type Session struct {
	UserId       uint `json:"user_id"`
	IsAuthorized bool `json:"is_authorized"`
}

type SessionRepo interface {
	GetSession(sessionId string) (*Session, error)
	SetNewUnauthorizedSession(expires time.Duration) (sessionId string, err error)
	MakeSessionAuthorized(sessionId string, userId uint) error
	MakeSessionUnauthorized(sessionId string) error
	DeleteSession(sessionId string) error
	GetSize() (int, error)
}
