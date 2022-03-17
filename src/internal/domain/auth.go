package domain

type AuthUseCase interface {
	Login(login string, password string) (sessionId string, err error)
	Logout(sessionId string) error
	SignUp(user *User) (sessionId string, err error)
}
