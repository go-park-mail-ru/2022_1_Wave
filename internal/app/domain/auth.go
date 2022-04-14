package domain

type AuthUseCase interface {
	Login(login string, password string, sessionId string) error
	Logout(sessionId string) error
	SignUp(user *User, sessionId string) error
	GetUnauthorizedSession() (sessionId string, err error)
	IsSession(sessionId string) bool
	IsAuthSession(sessionId string) bool
}
