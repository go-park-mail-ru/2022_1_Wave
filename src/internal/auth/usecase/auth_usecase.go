package usecase

import (
	"github.com/go-park-mail-ru/2022_1_Wave/config"
	domain2 "github.com/go-park-mail-ru/2022_1_Wave/internal/app/domain"
	"time"
)

type authUseCase struct {
	sessionRepo domain2.SessionRepo
	userRepo    domain2.UserRepo
}

var sessionExpire, _ = time.ParseDuration(config.C.SessionExpires)

func NewAuthUseCase(sessionRepo domain2.SessionRepo, userRepo domain2.UserRepo) domain2.AuthUseCase {
	return &authUseCase{
		sessionRepo: sessionRepo,
		userRepo:    userRepo,
	}
}

func (a *authUseCase) Login(login string, password string) (string, error) {

	if !a.userRepo.CheckEmailAndPassword(login, password) && !a.userRepo.CheckUsernameAndPassword(login, password) {
		return "", domain2.ErrInvalidLoginOrPassword
	}

	user, err := a.userRepo.SelectByUsername(login)
	if err != nil {
		user, _ = a.userRepo.SelectByEmail(login)
	}

	sessionId, err := a.sessionRepo.SetNewSession(sessionExpire, user.ID)
	if err != nil {
		return "", domain2.ErrWhileSetNewSession
	}

	return sessionId, nil
}

func (a *authUseCase) Logout(sessionId string) error {
	_ = a.sessionRepo.DeleteSession(sessionId)

	return nil
}

func (a *authUseCase) SignUp(user *domain2.User) (string, error) {
	_, err := a.userRepo.SelectByEmail(user.Email)
	if err != nil {
		return "", domain2.ErrUserAlreadyExist
	}

	_, err = a.userRepo.SelectByUsername(user.Username)
	if err != nil {
		return "", domain2.ErrUserAlreadyExist
	}

	err = a.userRepo.Insert(user)
	if err != nil {
		return "", domain2.ErrInsert
	}

	userToId, err := a.userRepo.SelectByEmail(user.Email)
	if err != nil {
		return "", domain2.ErrDatabaseUnexpected
	}

	sessionId, err := a.sessionRepo.SetNewSession(sessionExpire, userToId.ID)
	if err != nil {
		return "", domain2.ErrSessionStorageUnexpected
	}

	return sessionId, nil
}

func (a *authUseCase) GetUnauthorizedSession() (string, error) {
	return a.sessionRepo.SetNewUnauthorizedSession(sessionExpire)
}
