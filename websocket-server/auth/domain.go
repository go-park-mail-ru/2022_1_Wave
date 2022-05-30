package auth_domain

import "time"

type Session struct {
	UserId       uint `json:"user_id"`
	IsAuthorized bool `json:"is_authorized"`
}

type AuthAgent interface {
	GetSession(sessionId string) (*Session, error)
	SetNewUnauthorizedSession(expires time.Duration) (sessionId string, err error)
	SetNewAuthorizedSession(userId uint, expires time.Duration) (sessionId string, err error)
	MakeSessionAuthorized(sessionId string, userId uint) (newSessionId string, err error)
	MakeSessionUnauthorized(sessionId string) (newSessionId string, err error)
	DeleteSession(sessionId string) error
	IsSession(sessionId string) bool
	IsAuthSession(sessionId string) bool
}
