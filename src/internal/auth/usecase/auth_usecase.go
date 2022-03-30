package usecase

import (
	"github.com/go-park-mail-ru/2022_1_Wave/config"
	"github.com/go-park-mail-ru/2022_1_Wave/internal/domain"
	"time"
)

type authUseCase struct {
	sessionRepo domain.SessionRepo
	userRepo    domain.UserRepo
}

var sessionExpire, _ = time.ParseDuration(config.C.SessionExpires)

func NewAuthUseCase(sessionRepo domain.SessionRepo, userRepo domain.UserRepo) domain.AuthUseCase {
	return &authUseCase{
		sessionRepo: sessionRepo,
		userRepo:    userRepo,
	}
}

func (a *authUseCase) Login(login string, password string, sessionId string) error {
	if !a.userRepo.CheckEmailAndPassword(login, password) && !a.userRepo.CheckUsernameAndPassword(login, password) {
		return domain.ErrInvalidLoginOrPassword
	}

	user, err := a.userRepo.SelectByUsername(login)
	if err != nil {
		user, _ = a.userRepo.SelectByEmail(login)
	}

	err = a.sessionRepo.MakeSessionAuthorized(sessionId, user.ID)

	if err != nil {
		return domain.ErrWhileChangeSession
	}

	return nil
}

func (a *authUseCase) Logout(sessionId string) error {
	return a.sessionRepo.MakeSessionUnauthorized(sessionId)
}

func (a *authUseCase) SignUp(user *domain.User, sessionId string) error {
	_, err := a.userRepo.SelectByEmail(user.Email)
	if err != nil {
		return domain.ErrUserAlreadyExist
	}

	_, err = a.userRepo.SelectByUsername(user.Username)
	if err != nil {
		return domain.ErrUserAlreadyExist
	}

	err = a.userRepo.Insert(user)
	if err != nil {
		return domain.ErrInsert
	}

	userToId, err := a.userRepo.SelectByEmail(user.Email)
	if err != nil {
		return domain.ErrDatabaseUnexpected
	}

	err = a.sessionRepo.MakeSessionAuthorized(sessionId, userToId.ID)

	if err != nil {
		return domain.ErrWhileChangeSession
	}

	return nil
}

func (a *authUseCase) GetUnauthorizedSession() (string, error) {
	return a.sessionRepo.SetNewUnauthorizedSession(sessionExpire)
}

func (a *authUseCase) IsSession(sessionId string) bool {
	_, err := a.sessionRepo.GetSession(sessionId)
	return err == nil
}

func (a *authUseCase) IsAuthSession(sessionId string) bool {
	session, err := a.sessionRepo.GetSession(sessionId)
	if err != nil || session == nil {
		return false
	}

	return session.IsAuthorized
}
